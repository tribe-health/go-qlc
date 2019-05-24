package relation

import (
	"encoding/base64"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/qlcchain/go-qlc/common"
	"github.com/qlcchain/go-qlc/common/event"
	"github.com/qlcchain/go-qlc/common/types"
	"github.com/qlcchain/go-qlc/config"
	"github.com/qlcchain/go-qlc/ledger/relation/db"
	"github.com/qlcchain/go-qlc/log"
	"go.uber.org/zap"
)

type Relation struct {
	store         db.DbStore
	eb            event.EventBus
	logger        *zap.SugaredLogger
	addBlkChan    chan *types.StateBlock
	deleteBlkChan chan types.Hash
}

type blocksHash struct {
	Id        int64
	Hash      string
	Type      string
	Address   string
	Timestamp int64
}

type blocksMessage struct {
	Id        int64
	Hash      string
	Sender    string
	Receiver  string
	Message   string
	Timestamp int64
}

var (
	once     sync.Once
	relation *Relation
)

func NewRelation(cfg *config.Config) (*Relation, error) {
	var err error
	once.Do(func() {
		store := new(db.DBSQL)
		store, err = db.NewSQLDB(cfg)
		relation = &Relation{store: store,
			eb:            event.GetEventBus(cfg.LedgerDir()),
			addBlkChan:    make(chan *types.StateBlock, 65535),
			deleteBlkChan: make(chan types.Hash, 65535),
			logger:        log.NewLogger("relation")}
		go relation.processBlocks()
	})
	if err != nil {
		return nil, err
	}
	return relation, nil
}

func (r *Relation) Close() error {
	r.logger.Debugf("close sqlite, %p, ", r.store)
	return r.store.Close()
}

func (r *Relation) AccountBlocks(address types.Address, limit int, offset int) ([]types.Hash, error) {
	condition := make(map[db.Column]interface{})
	condition[db.ColumnAddress] = address.String()
	var h []blocksHash
	err := r.store.Read(db.TableBlockHash, condition, offset, limit, db.ColumnTimestamp, &h)
	if err != nil {
		return nil, err
	}
	return blockHash(h)
}

func (r *Relation) BlocksCount() (uint64, error) {
	var count uint64
	err := r.store.Count(db.TableBlockHash, &count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

type blocksType struct {
	Type  string
	Count uint64
}

func (r *Relation) BlocksCountByType() (map[string]uint64, error) {
	var t []blocksType
	err := r.store.Group(db.TableBlockHash, db.ColumnType, &t)
	if err != nil {
		return nil, err
	}
	return blockType(t), nil
}

func (r *Relation) Blocks(limit int, offset int) ([]types.Hash, error) {
	var h []blocksHash
	err := r.store.Read(db.TableBlockHash, nil, offset, limit, db.ColumnTimestamp, &h)
	if err != nil {
		return nil, err
	}
	return blockHash(h)
}

func (r *Relation) PhoneBlocks(phone []byte, sender bool, limit int, offset int) ([]types.Hash, error) {
	condition := make(map[db.Column]interface{})
	if sender == true {
		condition[db.ColumnSender] = phoneToString(phone)
	} else {
		condition[db.ColumnReceiver] = phoneToString(phone)
	}
	var m []blocksMessage
	err := r.store.Read(db.TableBlockMessage, condition, offset, limit, db.ColumnTimestamp, &m)
	if err != nil {
		return nil, err
	}
	return blockMessage(m)
}

func (r *Relation) MessageBlocks(hash types.Hash, limit int, offset int) ([]types.Hash, error) {
	condition := make(map[db.Column]interface{})
	condition[db.ColumnMessage] = hash.String()
	var m []blocksMessage
	err := r.store.Read(db.TableBlockMessage, condition, offset, limit, db.ColumnTimestamp, &m)
	if err != nil {
		return nil, err
	}
	return blockMessage(m)
}

func (r *Relation) AddBlock(block *types.StateBlock) error {
	r.logger.Debug("add relation, ", block.GetHash())
	conHash := make(map[db.Column]interface{})
	conHash[db.ColumnHash] = block.GetHash().String()
	conHash[db.ColumnTimestamp] = block.GetTimestamp()
	conHash[db.ColumnType] = block.GetType().String()
	conHash[db.ColumnAddress] = block.GetAddress().String()
	if err := r.store.Create(db.TableBlockHash, conHash); err != nil {
		return err
	}
	message := block.GetMessage()
	if block.GetSender() != nil || block.GetReceiver() != nil || !message.IsZero() {
		conMessage := make(map[db.Column]interface{})
		conMessage[db.ColumnHash] = block.GetHash().String()
		conMessage[db.ColumnMessage] = message.String()
		conMessage[db.ColumnSender] = phoneToString(block.GetSender())
		conMessage[db.ColumnReceiver] = phoneToString(block.GetReceiver())
		conMessage[db.ColumnTimestamp] = block.GetTimestamp()
		if err := r.store.Create(db.TableBlockMessage, conMessage); err != nil {
			return err
		}
	}
	return nil
}

func (r *Relation) DeleteBlock(hash types.Hash) error {
	r.logger.Info("delete relation, ", hash.String())
	condition := make(map[db.Column]interface{})
	condition[db.ColumnHash] = hash.String()
	err := r.store.Delete(db.TableBlockHash, condition)
	if err != nil {
		return err
	}
	return r.store.Delete(db.TableBlockMessage, condition)
}

func (r *Relation) BatchUpdate(fn func(txn *sqlx.Tx) error) error {
	tx := r.store.NewTransaction()
	if err := fn(tx); err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (r *Relation) AddBlocks(txn *sqlx.Tx, blocks []*types.StateBlock) error {
	blksHashes := make([]*blocksHash, 0)
	blksMessage := make([]*blocksMessage, 0)
	//r.logger.Info("batch block count: ", len(blocks))
	for _, block := range blocks {
		blksHashes = append(blksHashes, &blocksHash{
			Hash:      block.GetHash().String(),
			Type:      block.GetType().String(),
			Address:   block.GetAddress().String(),
			Timestamp: block.GetTimestamp(),
		})

		message := block.GetMessage()
		if block.GetSender() != nil || block.GetReceiver() != nil || !message.IsZero() {
			blksMessage = append(blksMessage, &blocksMessage{
				Hash:      block.GetHash().String(),
				Sender:    phoneToString(block.GetSender()),
				Receiver:  phoneToString(block.GetReceiver()),
				Message:   message.String(),
				Timestamp: block.GetTimestamp(),
			})
		}
	}
	if len(blksHashes) > 0 {
		if _, err := txn.NamedExec("INSERT INTO BLOCKHASH(hash, type,address,timestamp) VALUES (:hash,:type,:address,:timestamp) ", blksHashes); err != nil {
			r.logger.Errorf("insert block hash, ", err)
			return err
		}
	}
	if len(blksMessage) > 0 {
		if _, err := txn.NamedExec("INSERT INTO BLOCKMESSAGE(hash, sender,receiver,message,timestamp) VALUES (:hash,:sender,:receiver,:message,:timestamp)", blksMessage); err != nil {
			r.logger.Error("insert message, ", err)
			return err
		}
	}
	return nil
}

func (r *Relation) EmptyStore() error {
	r.logger.Info("empty store")
	err := r.store.Delete(db.TableBlockHash, nil)
	if err != nil {
		return err
	}
	return r.store.Delete(db.TableBlockMessage, nil)
}

func phoneToString(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func blockHash(bs []blocksHash) ([]types.Hash, error) {
	hs := make([]types.Hash, 0)
	for _, b := range bs {
		var h types.Hash
		if err := h.Of(b.Hash); err != nil {
			return nil, err
		}
		hs = append(hs, h)
	}
	return hs, nil
}

func blockMessage(bs []blocksMessage) ([]types.Hash, error) {
	hs := make([]types.Hash, 0)
	for _, b := range bs {
		var h types.Hash
		if err := h.Of(b.Hash); err != nil {
			return nil, err
		}
		hs = append(hs, h)
	}
	return hs, nil
}

func blockType(bs []blocksType) map[string]uint64 {
	t := make(map[string]uint64)
	for _, b := range bs {
		t[b.Type] = b.Count
	}
	return t
}

func (r *Relation) processBlocks() {
	for {
		select {
		case blk := <-r.addBlkChan:
			if err := r.AddBlock(blk); err != nil {
				r.logger.Error(err)
			}
		case blk := <-r.deleteBlkChan:
			if err := r.DeleteBlock(blk); err != nil {
				r.logger.Error(err)
			}
		default:
			time.Sleep(5 * time.Millisecond)
		}
	}
}

func (r *Relation) waitAddBlocks(block *types.StateBlock) {
	r.addBlkChan <- block
}

func (r *Relation) waitDeleteBlocks(hash types.Hash) {
	r.deleteBlkChan <- hash
}

func (r *Relation) SetEvent() error {
	err := r.eb.Subscribe(string(common.EventAddRelation), r.waitAddBlocks)
	if err != nil {
		r.logger.Error(err)
		return err
	}
	err = r.eb.Subscribe(string(common.EventDeleteRelation), r.waitDeleteBlocks)
	if err != nil {
		r.logger.Error(err)
		return err
	}
	return nil
}

func (r *Relation) UnsubscribeEvent() error {
	err := r.eb.Unsubscribe(string(common.EventAddRelation), r.waitAddBlocks)
	if err != nil {
		r.logger.Error(err)
		return err
	}
	err = r.eb.Unsubscribe(string(common.EventDeleteRelation), r.waitDeleteBlocks)
	if err != nil {
		r.logger.Error(err)
		return err
	}
	return nil
}
