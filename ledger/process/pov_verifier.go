package process

import (
	"errors"
	"fmt"
	"github.com/qlcchain/go-qlc/common"
	"github.com/qlcchain/go-qlc/common/merkle"
	"github.com/qlcchain/go-qlc/common/types"
	"github.com/qlcchain/go-qlc/ledger"
	"github.com/qlcchain/go-qlc/log"
	"github.com/qlcchain/go-qlc/trie"
	"go.uber.org/zap"
	"time"
)

type PovVerifier struct {
	store  ledger.Store
	chain  PovVerifierChainReader
	logger *zap.SugaredLogger
}

type PovVerifyStat struct {
	Result    ProcessResult
	ErrMsg    string
	TxResults map[types.Hash]ProcessResult

	PrevBlock     *types.PovBlock
	PrevStateTrie *trie.Trie
	StateTrie     *trie.Trie
	TxBlocks      map[types.Hash]*types.StateBlock
}

func NewPovVerifyStat() *PovVerifyStat {
	pvs := new(PovVerifyStat)
	pvs.TxResults = make(map[types.Hash]ProcessResult)
	pvs.TxBlocks = make(map[types.Hash]*types.StateBlock)
	return pvs
}

func (pvs *PovVerifyStat) setResult(result ProcessResult, err error) {
	pvs.Result = result
	if err != nil {
		pvs.ErrMsg = err.Error()
	}
}

func (pvs *PovVerifyStat) getPrevBlock(pv *PovVerifier, prevHash types.Hash) *types.PovBlock {
	if pvs.PrevBlock == nil {
		pvs.PrevBlock = pv.chain.GetBlockByHash(prevHash)
	}

	return pvs.PrevBlock
}

type PovVerifierChainReader interface {
	GetBlockByHash(hash types.Hash) *types.PovBlock
	CalcNextRequiredTarget(block *types.PovBlock) (types.Signature, error)
	GenStateTrie(prevStateHash types.Hash, txs []*types.PovTransaction) (*trie.Trie, error)
	GetStateTrie(stateHash *types.Hash) *trie.Trie
}

func NewPovVerifier(store ledger.Store, chain PovVerifierChainReader) *PovVerifier {
	return &PovVerifier{store: store, chain: chain, logger: log.NewLogger("pov_verifier")}
}

func (pv *PovVerifier) Process(block types.Block) (ProcessResult, error) {
	return Other, nil
}

func (pv *PovVerifier) BlockCheck(block types.Block) (ProcessResult, error) {
	return Other, nil
}

func (pv *PovVerifier) VerifyNet(block *types.PovBlock) *PovVerifyStat {
	stat := NewPovVerifyStat()

	result, err := pv.verifyHeader(block, stat)
	if err != nil {
		stat.Result = result
		stat.ErrMsg = err.Error()
		return stat
	}

	result, err = pv.verifyTimestamp(block, stat)
	if err != nil || result != Progress {
		stat.setResult(result, err)
		return stat
	}

	stat.Result = Progress
	return stat
}

func (pv *PovVerifier) VerifyFull(block *types.PovBlock) *PovVerifyStat {
	stat := NewPovVerifyStat()

	result, err := pv.verifyHeader(block, stat)
	if err != nil || result != Progress {
		stat.setResult(result, err)
		return stat
	}

	result, err = pv.verifyTimestamp(block, stat)
	if err != nil || result != Progress {
		stat.setResult(result, err)
		return stat
	}

	result, err = pv.verifyReferred(block, stat)
	if err != nil || result != Progress {
		stat.setResult(result, err)
		return stat
	}

	result, err = pv.verifyProducer(block, stat)
	if err != nil || result != Progress {
		stat.setResult(result, err)
		return stat
	}

	result, err = pv.verifyTarget(block, stat)
	if err != nil || result != Progress {
		stat.setResult(result, err)
		return stat
	}

	result, err = pv.verifyTransactions(block, stat)
	if err != nil || result != Progress {
		stat.setResult(result, err)
		return stat
	}

	result, err = pv.verifyState(block, stat)
	if err != nil || result != Progress {
		stat.setResult(result, err)
		return stat
	}

	stat.Result = Progress
	return stat
}

func (pv *PovVerifier) verifyHeader(block *types.PovBlock, stat *PovVerifyStat) (ProcessResult, error) {
	computedHash := block.ComputeHash()
	if block.Hash.IsZero() || computedHash != block.Hash {
		return BadHash, fmt.Errorf("bad hash, %s != %s", computedHash, block.Hash)
	}

	if block.Coinbase.IsZero() {
		return BadSignature, errors.New("coinbase is zero")
	}

	if block.Signature.IsZero() {
		return BadSignature, errors.New("signature is zero")
	}

	isVerified := block.Coinbase.Verify(block.GetHash().Bytes(), block.GetSignature().Bytes())
	if !isVerified {
		return BadSignature, errors.New("bad signature")
	}

	return Progress, nil
}

func (pv *PovVerifier) verifyTimestamp(block *types.PovBlock, stat *PovVerifyStat) (ProcessResult, error) {
	if block.Timestamp <= 0 {
		return InvalidTime, errors.New("timestamp is zero")
	}

	ts := time.Unix(block.Timestamp, 0)
	if ts.After(time.Now().Add(time.Hour)) {
		return InvalidTime, errors.New("timestamp too far from future")
	}

	return Progress, nil
}

func (pv *PovVerifier) verifyReferred(block *types.PovBlock, stat *PovVerifyStat) (ProcessResult, error) {
	prevBlock := pv.chain.GetBlockByHash(block.GetPrevious())
	if prevBlock == nil {
		return GapPrevious, nil
	}

	if block.GetTimestamp() <= prevBlock.GetTimestamp() {
		return InvalidTime, fmt.Errorf("timestamp %d before than previous %d", block.GetTimestamp(), prevBlock.GetTimestamp())
	}

	/*
		if block.GetTimestamp() > (prevBlock.GetTimestamp() + int64(common.PovMaxTimeOffsetPrevSec)) {
			return InvalidTime, fmt.Errorf("timestamp %d too far from previous %d", block.GetTimestamp(), prevBlock.GetTimestamp())
		}
	*/

	if block.GetHeight() != prevBlock.GetHeight()+1 {
		return InvalidHeight, fmt.Errorf("height %d not continue with previous %d", block.GetHeight(), prevBlock.GetHeight())
	}

	return Progress, nil
}

func (pv *PovVerifier) verifyTransactions(block *types.PovBlock, stat *PovVerifyStat) (ProcessResult, error) {
	if block.TxNum != uint32(len(block.Transactions)) {
		return InvalidTxNum, nil
	}

	if len(block.Transactions) <= 0 {
		if !block.MerkleRoot.IsZero() {
			return BadMerkleRoot, fmt.Errorf("bad merkle root not zero when txs empty")
		}
		return Progress, nil
	}

	var txHashList []*types.Hash
	for _, tx := range block.Transactions {
		txHash := tx.Hash
		txHashList = append(txHashList, &txHash)
	}
	merkleRoot := merkle.CalcMerkleTreeRootHash(txHashList)
	if merkleRoot.IsZero() {
		return BadMerkleRoot, fmt.Errorf("bad merkle root is zero when txs exist")
	}
	if merkleRoot != block.MerkleRoot {
		return BadMerkleRoot, fmt.Errorf("bad merkle root not equals %s != %s", merkleRoot, block.MerkleRoot)
	}
	for _, tx := range block.Transactions {
		txBlock, _ := pv.store.GetStateBlock(tx.Hash)
		if txBlock == nil {
			stat.TxResults[tx.Hash] = GapTransaction
		} else {
			tx.Block = txBlock
			stat.TxBlocks[tx.Hash] = txBlock
		}
	}

	if len(stat.TxResults) > 0 {
		return GapTransaction, fmt.Errorf("total %d txs in pending", len(stat.TxResults))
	}

	return Progress, nil
}

func (pv *PovVerifier) verifyState(block *types.PovBlock, stat *PovVerifyStat) (ProcessResult, error) {
	prevBlock := stat.getPrevBlock(pv, block.GetPrevious())
	if prevBlock == nil {
		return GapPrevious, fmt.Errorf("prev block %s pending", block.GetPrevious())
	}

	stateTrie, err := pv.chain.GenStateTrie(prevBlock.StateHash, block.Transactions)
	if err != nil {
		return BadStateHash, err
	}
	stateHash := types.Hash{}
	if stateTrie != nil {
		stateHash = *stateTrie.Hash()
	}
	if stateHash != block.StateHash {
		return BadStateHash, fmt.Errorf("state hash is not equals %s != %s", stateHash, block.StateHash)
	}
	stat.StateTrie = stateTrie

	return Progress, nil
}

func (pv *PovVerifier) verifyTarget(block *types.PovBlock, stat *PovVerifyStat) (ProcessResult, error) {
	prevBlock := stat.getPrevBlock(pv, block.GetPrevious())
	if prevBlock == nil {
		return GapPrevious, nil
	}

	expectedTarget, err := pv.chain.CalcNextRequiredTarget(prevBlock)
	if err != nil {
		return BadTarget, err
	}
	if expectedTarget != block.Target {
		return BadTarget, errors.New("target not equal next required target")
	}

	voteHash := block.ComputeVoteHash()
	voteSig := block.GetVoteSignature()

	isVerified := block.GetCoinbase().Verify(voteHash.Bytes(), voteSig.Bytes())
	if !isVerified {
		return BadSignature, errors.New("bad vote signature")
	}

	voteSigInt := voteSig.ToBigInt()

	targetSig := block.GetTarget()
	targetInt := targetSig.ToBigInt()

	if targetInt.Cmp(common.PovMinimumTargetInt) < 0 {
		return BadTarget, errors.New("target lesser than minimum")
	}
	if targetInt.Cmp(common.PovMaximumTargetInt) > 0 {
		return BadTarget, errors.New("target greater than maximum")
	}

	if voteSigInt.Cmp(targetInt) > 0 {
		return BadTarget, errors.New("target greater than vote signature")
	}

	return Progress, nil
}

func (pv *PovVerifier) verifyProducer(block *types.PovBlock, stat *PovVerifyStat) (ProcessResult, error) {
	if block.GetHeight() < common.PovMinerVerifyHeightStart {
		return Progress, nil
	}

	prevBlock := stat.getPrevBlock(pv, block.GetPrevious())
	if prevBlock == nil {
		return GapPrevious, nil
	}
	stateHash := prevBlock.GetStateHash()

	prevTrie := pv.chain.GetStateTrie(&stateHash)
	if prevTrie == nil {
		return BadStateHash, errors.New("failed to get state tire by prev state hash")
	}
	stat.PrevStateTrie = prevTrie

	asBytes := prevTrie.GetValue(block.GetCoinbase().Bytes())
	if len(asBytes) <= 0 {
		return BadCoinbase, errors.New("failed to get account state value")
	}

	as := new(types.PovAccountState)
	err := as.Deserialize(asBytes)
	if err != nil {
		return BadCoinbase, errors.New("failed to deserialize account state value")
	}

	if as.RepState == nil {
		return BadCoinbase, errors.New("account rep state is nil")
	}
	rs := as.RepState

	if rs.Vote.Compare(common.PovMinerPledgeAmountMin) == types.BalanceCompSmaller {
		return BadCoinbase, errors.New("coinbase pledge amount not enough")
	}

	return Progress, nil
}
