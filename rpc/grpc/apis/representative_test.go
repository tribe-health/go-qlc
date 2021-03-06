package apis

import (
	"context"
	"math/big"
	"testing"

	chainctx "github.com/qlcchain/go-qlc/chain/context"
	"github.com/qlcchain/go-qlc/common"
	"github.com/qlcchain/go-qlc/common/types"
	"github.com/qlcchain/go-qlc/common/vmcontract/contractaddress"
	"github.com/qlcchain/go-qlc/config"
	"github.com/qlcchain/go-qlc/mock"
	"github.com/qlcchain/go-qlc/rpc/api"
	cabi "github.com/qlcchain/go-qlc/vm/contract/abi"
	"github.com/qlcchain/go-qlc/vm/vmstore"
)

func TestRepApi_GetAvailRewardInfo(t *testing.T) {
	clear, l, cfgFile := getTestLedger()
	if l == nil {
		t.Fatal()
	}
	defer clear()

	cc := chainctx.NewChainContext(cfgFile)
	cfg, _ := cc.Config()
	r := NewRepAPI(cfg, l)
	account := mock.Address()

	ri, _ := r.GetAvailRewardInfo(context.Background(), toAddress(account))
	if ri != nil {
		t.Fatal()
	}

	pb := addPovBlock(l, nil, 100)
	ri, _ = r.GetAvailRewardInfo(context.Background(), toAddress(account))
	if ri == nil || ri.NeedCallReward {
		t.Fatal()
	}

	pb = addPovBlock(l, pb, common.PovMinerRewardHeightStart+uint64(common.POVChainBlocksPerDay)+common.PovMinerRewardHeightGapToLatest)
	ri, _ = r.GetAvailRewardInfo(context.Background(), toAddress(account))
	if ri == nil || ri.NeedCallReward {
		t.Fatal()
	}

	ds := types.NewPovMinerDayStat()
	it := types.NewPovMinerStatItem()
	it.FirstHeight = 1440
	it.LastHeight = 2880
	it.RepBlockNum = 120
	it.RepReward = types.NewBalance(10)
	ds.DayIndex = uint32(common.PovMinerRewardHeightStart / uint64(common.POVChainBlocksPerDay))
	ds.MinerStats[account.String()] = it
	l.AddPovMinerStat(ds)
	ri, _ = r.GetAvailRewardInfo(context.Background(), toAddress(account))
	if ri == nil || !ri.NeedCallReward {
		t.Fatal()
	}
}

func TestRepApi_UnpackRewardData(t *testing.T) {
	clear, l, cfgFile := getTestLedger()
	if l == nil {
		t.Fatal()
	}
	defer clear()

	cc := chainctx.NewChainContext(cfgFile)
	cfg, _ := cc.Config()
	r := NewRepAPI(cfg, l)
	param := new(api.RepRewardParam)
	param.Account = mock.Address()
	param.Beneficial = mock.Address()
	param.RewardBlocks = 10
	param.StartHeight = 0
	param.EndHeight = 1439
	param.RewardAmount = big.NewInt(100)
	data, _ := r.GetRewardData(context.Background(), toRepRewardParam(param))

	rr, err := r.UnpackRewardData(context.Background(), data)
	if err != nil {
		t.Fatal(err)
	}
	if err != nil || rr.Account != param.Account.String() || rr.Beneficial != param.Beneficial.String() || rr.EndHeight != param.EndHeight ||
		rr.StartHeight != param.StartHeight || rr.RewardBlocks != param.RewardBlocks ||
		toOriginBalanceByValue(rr.RewardAmount).Cmp(param.RewardAmount) != 0 {
		t.Fatal()
	}
}

func TestRepApi_GetRewardSendBlock(t *testing.T) {
	clear, l, cfgFile := getTestLedger()
	if l == nil {
		t.Fatal()
	}
	defer clear()

	cc := chainctx.NewChainContext(cfgFile)
	cfg, _ := cc.Config()
	r := NewRepAPI(cfg, l)
	param := new(api.RepRewardParam)
	param.RewardBlocks = 10
	param.StartHeight = common.PovMinerRewardHeightStart
	param.EndHeight = common.PovMinerRewardHeightStart + uint64(common.POVChainBlocksPerDay) - 1
	param.RewardAmount = big.NewInt(100)

	blk, _ := r.GetRewardSendBlock(context.Background(), toRepRewardParam(param))
	if blk != nil {
		t.Fatal()
	}

	param.Account = mock.Address()
	blk, _ = r.GetRewardSendBlock(context.Background(), toRepRewardParam(param))
	if blk != nil {
		t.Fatal()
	}

	param.Beneficial = mock.Address()
	blk, _ = r.GetRewardSendBlock(context.Background(), toRepRewardParam(param))
	if blk != nil {
		t.Fatal()
	}

	am := mock.AccountMeta(param.Account)
	l.AddAccountMeta(am, l.Cache().GetCache())
	blk, _ = r.GetRewardSendBlock(context.Background(), toRepRewardParam(param))
	if blk != nil {
		t.Fatal()
	}

	am.Tokens[0].Type = config.ChainToken()
	l.UpdateAccountMeta(am, l.Cache().GetCache())
	blk, _ = r.GetRewardSendBlock(context.Background(), toRepRewardParam(param))
	if blk != nil {
		t.Fatal()
	}

	pb, td := mock.GeneratePovBlock(nil, 0)
	pb.Header.BasHdr.Height = uint64(common.POVChainBlocksPerDay) * 4
	l.AddPovBlock(pb, td)
	l.SetPovLatestHeight(pb.Header.BasHdr.Height)
	l.AddPovBestHash(pb.Header.BasHdr.Height, pb.GetHash())

	ds := types.NewPovMinerDayStat()
	it := types.NewPovMinerStatItem()
	it.FirstHeight = param.StartHeight
	it.LastHeight = param.EndHeight
	it.RepBlockNum = uint32(param.RewardBlocks)
	it.RepReward = types.Balance{Int: param.RewardAmount}
	ds.DayIndex = uint32(common.PovMinerRewardHeightStart / uint64(common.POVChainBlocksPerDay))
	ds.MinerStats[param.Account.String()] = it
	l.AddPovMinerStat(ds)

	blk, err := r.GetRewardSendBlock(context.Background(), toRepRewardParam(param))
	if blk == nil {
		t.Fatal(err)
	}
}

func TestRepApi_GetRewardRecvBlock(t *testing.T) {
	clear, l, cfgFile := getTestLedger()
	if l == nil {
		t.Fatal()
	}
	defer clear()

	cc := chainctx.NewChainContext(cfgFile)
	cfg, _ := cc.Config()
	r := NewRepAPI(cfg, l)

	sendBlock := mock.StateBlockWithoutWork()
	param := new(api.RepRewardParam)
	param.Account = mock.Address()
	param.Beneficial = mock.Address()
	param.RewardBlocks = 10
	param.StartHeight = 0
	param.EndHeight = 1439
	param.RewardAmount = big.NewInt(100)
	re, err := r.GetRewardData(context.Background(), toRepRewardParam(param))
	if err != nil {
		t.Fatal(err)
	}
	sendBlock.Data = re.GetValue()

	blk, _ := r.GetRewardRecvBlock(context.Background(), toStateBlock(sendBlock))
	if blk != nil {
		t.Fatal()
	}

	sendBlock.Type = types.ContractSend
	blk, _ = r.GetRewardRecvBlock(context.Background(), toStateBlock(sendBlock))
	if blk != nil {
		t.Fatal()
	}

	sendBlock.Link = contractaddress.RepAddress.ToHash()
	blk, _ = r.GetRewardRecvBlock(context.Background(), toStateBlock(sendBlock))
	if blk != nil {
		t.Fatal()
	}
}

func TestRepApi_GetRewardRecvBlockBySendHash(t *testing.T) {
	clear, l, cfgFile := getTestLedger()
	if l == nil {
		t.Fatal()
	}
	defer clear()

	cc := chainctx.NewChainContext(cfgFile)
	cfg, _ := cc.Config()
	r := NewRepAPI(cfg, l)

	sendBlock := mock.StateBlockWithoutWork()
	param := new(api.RepRewardParam)
	param.Account = mock.Address()
	param.Beneficial = mock.Address()
	param.RewardBlocks = 10
	param.StartHeight = common.PovMinerRewardHeightStart
	param.EndHeight = param.StartHeight + uint64(common.POVChainBlocksPerDay)
	param.RewardAmount = big.NewInt(100)
	re, err := r.GetRewardData(context.Background(), toRepRewardParam(param))
	if err != nil {
		t.Fatal(err)
	}
	sendBlock.Data = re.GetValue()
	sendBlock.Address = param.Account
	sendBlock.Type = types.ContractSend
	sendBlock.Link = contractaddress.RepAddress.ToHash()
	l.AddStateBlock(sendBlock)

	blk, _ := r.GetRewardRecvBlockBySendHash(context.Background(), toHash(sendBlock.GetHash()))
	if blk == nil {
		t.Fatal()
	}
}

func TestRepApi_GetRewardHistory(t *testing.T) {
	clear, l, cfgFile := getTestLedger()
	if l == nil {
		t.Fatal()
	}
	defer clear()

	cc := chainctx.NewChainContext(cfgFile)
	cfg, _ := cc.Config()
	r := NewRepAPI(cfg, l)
	ctx := vmstore.NewVMContext(l, &contractaddress.RepAddress)
	account := mock.Address()

	timeStamp := common.TimeNow().Unix()
	data, _ := cabi.RepABI.PackVariable(cabi.VariableNameRepReward, uint64(1440), uint64(100), timeStamp, big.NewInt(200))
	err := ctx.SetStorage(contractaddress.RepAddress.Bytes(), account[:], data)
	if err != nil {
		t.Fatal(err)
	}

	err = l.SaveStorage(vmstore.ToCache(ctx))
	if err != nil {
		t.Fatal(err)
	}

	h, _ := r.GetRewardHistory(context.Background(), toAddress(account))
	if h == nil || h.RewardBlocks != 100 || toOriginBalanceByValue(h.RewardAmount).Compare(types.NewBalance(200)) != types.BalanceCompEqual ||
		h.LastEndHeight != 1440 || h.LastRewardTime != timeStamp {
		t.Fatal()
	}
}
