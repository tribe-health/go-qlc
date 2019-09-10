package process

import (
	"errors"
	"fmt"

	"github.com/qlcchain/go-qlc/common"
	"github.com/qlcchain/go-qlc/common/types"
	"github.com/qlcchain/go-qlc/ledger"
	"github.com/qlcchain/go-qlc/ledger/db"
	"github.com/qlcchain/go-qlc/vm/contract"
)

func (lv *LedgerVerifier) BlockSyncCheck(block types.Block) (ProcessResult, error) {
	if b, ok := block.(*types.StateBlock); ok {
		if fn, ok := lv.checkSyncBlockFns[b.Type]; ok {
			r, err := fn(lv, b)
			if err != nil {
				lv.logger.Error(fmt.Sprintf("error:%s, sync block:%s", err.Error(), b.GetHash().String()))
			}
			if r != Progress {
				lv.logger.Debugf(fmt.Sprintf("process result:%s, sync block:%s", r.String(), b.GetHash().String()))
			}
			return r, err
		} else {
			return Other, fmt.Errorf("unsupport block type %s", b.Type.String())
		}
	} else if _, ok := block.(*types.SmartContractBlock); ok {
		return Other, errors.New("smart contract block")
	}
	return Other, errors.New("invalid block")
}

func checkSyncStateBlock(lv *LedgerVerifier, block *types.StateBlock) (ProcessResult, error) {
	hash := block.GetHash()
	address := block.GetAddress()

	lv.logger.Debug("check block ", hash)
	blockExist, err := lv.l.HasStateBlockConfirmed(hash)
	if err != nil {
		return Other, err
	}

	if blockExist {
		return Old, nil
	}

	if block.GetType() == types.ContractSend {
		if block.GetLink() == types.Hash(types.RewardsAddress) {
			return Progress, nil
		}
	}
	if block.GetType() == types.ContractReward {
		//linkBlk, err := lv.l.GetStateBlockConfirmed(block.GetLink())
		//if err != nil {
		//	return GapSource, nil
		//}
		//if linkBlk.GetLink() == types.Hash(types.RewardsAddress) {
		//	return Progress, nil
		//}
		return Progress, nil
	}

	if !block.IsValid() {
		return BadWork, errors.New("bad work")
	}

	signature := block.GetSignature()
	if !address.Verify(hash[:], signature[:]) {
		return BadSignature, errors.New("bad signature")
	}

	return Progress, nil
}

func checkSyncReceiveBlock(lv *LedgerVerifier, block *types.StateBlock) (ProcessResult, error) {
	result, err := checkSyncStateBlock(lv, block)
	if err != nil || result != Progress {
		return result, err
	}

	// check previous
	if previous, err := lv.l.GetStateBlockConfirmed(block.Previous); err != nil {
		return GapPrevious, nil
	} else {
		//check fork
		if tm, err := lv.l.GetTokenMetaConfirmed(block.Address, block.GetToken()); err == nil && previous.GetHash() != tm.Header {
			return Fork, nil
		}
	}
	return Progress, nil
}

func checkSyncOpenBlock(lv *LedgerVerifier, block *types.StateBlock) (ProcessResult, error) {
	result, err := checkSyncStateBlock(lv, block)
	if err != nil || result != Progress {
		return result, err
	}

	//check previous
	if !block.Previous.IsZero() {
		return Other, fmt.Errorf("open block previous is not zero")
	}

	//check fork
	if _, err := lv.l.GetTokenMetaConfirmed(block.Address, block.Token); err == nil {
		return Fork, nil
	}
	return Progress, nil
}

func checkSyncContractReceiveBlock(lv *LedgerVerifier, block *types.StateBlock) (ProcessResult, error) {
	//ignore chain genesis block
	if common.IsGenesisBlock(block) {
		return Progress, nil
	}

	result, err := checkSyncStateBlock(lv, block)
	if err != nil || result != Progress {
		return result, err
	}
	// check previous
	if !block.IsOpen() {
		// check previous
		if previous, err := lv.l.GetStateBlockConfirmed(block.Previous); err != nil {
			return GapPrevious, nil
		} else {
			//check fork
			if tm, err := lv.l.GetTokenMetaConfirmed(block.Address, block.GetToken()); err == nil && previous.GetHash() != tm.Header {
				return Fork, nil
			}
		}
	} else {
		//check fork
		if _, err := lv.l.GetTokenMetaConfirmed(block.Address, block.Token); err == nil {
			return Fork, nil
		}
	}
	return Progress, nil
}

func (lv *LedgerVerifier) BlockSyncProcess(block types.Block) error {
	return lv.l.BatchUpdate(func(txn db.StoreTxn) error {
		if state, ok := block.(*types.StateBlock); ok {
			err := lv.processSyncBlock(state, txn)
			if err != nil {
				lv.logger.Error(fmt.Sprintf("%s, sync block:%s", err.Error(), state.GetHash().String()))
				return err
			}
			return nil
		} else if _, ok := block.(*types.SmartContractBlock); ok {
			return errors.New("smart contract block")
		}
		return errors.New("invalid block")
	})
}

func (lv *LedgerVerifier) processSyncBlock(block *types.StateBlock, txn db.StoreTxn) error {
	lv.logger.Debug("process sync block, ", block.GetHash())
	if err := lv.l.AddStateBlock(block, txn); err != nil {
		return err
	}
	if block.IsSendBlock() || block.GetType() == types.ContractReward {
		if err := lv.l.AddSyncBlock(block); err != nil {
			return fmt.Errorf("add sync block error: %s", err)
		}
	}
	am, err := lv.l.GetAccountMetaConfirmed(block.GetAddress(), txn)
	if err != nil && err != ledger.ErrAccountNotFound {
		return fmt.Errorf("get account meta error: %s", err)
	}
	tm, err := lv.l.GetTokenMetaConfirmed(block.GetAddress(), block.GetToken(), txn)
	if err != nil && err != ledger.ErrAccountNotFound && err != ledger.ErrTokenNotFound {
		return fmt.Errorf("get token meta error: %s", err)
	}
	if err := lv.updateRepresentative(block, am, tm, txn); err != nil {
		return fmt.Errorf("update representative error: %s", err)
	}
	if err := lv.updateFrontier(block, tm, txn); err != nil {
		return fmt.Errorf("update frontier error: %s", err)
	}
	if err := lv.updateAccountMeta(block, am, txn); err != nil {
		return fmt.Errorf("update account meta error: %s", err)
	}
	return nil
}

func (lv *LedgerVerifier) BlockSyncDown() error {
	txn := lv.l.Store.NewTransaction(true)
	count := 0
	err := lv.l.GetSyncBlocks(func(block *types.StateBlock) error {
		if block.IsSendBlock() {
			if _, err := lv.l.GetLinkBlock(block.GetHash()); err != nil {
				hash := block.GetHash()
				switch block.Type {
				case types.Send:
					preBlk, err := lv.l.GetStateBlockConfirmed(block.Previous)
					if err != nil {
						return errors.New("previous block not found")
					}
					pending := types.PendingInfo{
						Source: block.GetAddress(),
						Type:   block.GetToken(),
						Amount: preBlk.Balance.Sub(block.GetBalance()),
					}
					pendingKey := types.PendingKey{
						Address: types.Address(block.GetLink()),
						Hash:    hash,
					}
					lv.logger.Debug("add pending, ", pendingKey)
					if err := lv.l.AddPending(&pendingKey, &pending, txn); err != nil {
						return err
					}
				case types.ContractSend:
					if c, ok, err := contract.GetChainContract(types.Address(block.Link), block.Data); ok && err == nil {
						if pendingKey, pendingInfo, err := c.DoPending(block); err == nil && pendingKey != nil {
							lv.logger.Debug("contractSend add sync pending , ", pendingKey)
							if err := lv.l.AddPending(pendingKey, pendingInfo, txn); err != nil {
								return err
							}
						}
					}
				}
			}
		}
		if block.GetType() == types.ContractReward {
			if err := lv.updateContractData(block, txn); err != nil {
				return fmt.Errorf("update contract data error: %s", err)
			}
		}
		count++
		if count > 2000 {
			count = 0
			if err := txn.Commit(nil); err != nil {
				return fmt.Errorf("commit error, %s ", err)
			}
			txn = lv.l.Store.NewTransaction(true)
		}
		return nil
	})
	if err != nil {
		return err
	}
	if err := lv.l.DropSyncBlocks(); err != nil {
		return fmt.Errorf("sync done error, %s", err)
	}
	return txn.Commit(nil)
}