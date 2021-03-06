package apis

import (
	"context"
	"math/big"

	"go.uber.org/zap"

	chainctx "github.com/qlcchain/go-qlc/chain/context"
	"github.com/qlcchain/go-qlc/ledger"
	"github.com/qlcchain/go-qlc/log"
	"github.com/qlcchain/go-qlc/rpc/api"
	pb "github.com/qlcchain/go-qlc/rpc/grpc/proto"
	pbtypes "github.com/qlcchain/go-qlc/rpc/grpc/proto/types"
	"github.com/qlcchain/go-qlc/vm/contract/abi"
)

type BlackHoleAPI struct {
	blackHole *api.BlackHoleAPI
	logger    *zap.SugaredLogger
}

func NewBlackHoleAPI(l ledger.Store, cc *chainctx.ChainContext) *BlackHoleAPI {
	return &BlackHoleAPI{
		blackHole: api.NewBlackHoleApi(l, cc),
		logger:    log.NewLogger("grpc_blackHole"),
	}
}

func (b *BlackHoleAPI) GetSendBlock(ctx context.Context, param *pb.DestroyParam) (*pbtypes.StateBlock, error) {
	owner, err := toOriginAddressByValue(param.GetOwner())
	if err != nil {
		return nil, err
	}
	previous, err := toOriginHashByValue(param.GetPrevious())
	if err != nil {
		return nil, err
	}
	token, err := toOriginHashByValue(param.GetToken())
	if err != nil {
		return nil, err
	}
	amount := big.NewInt(param.GetAmount())
	sign, err := toOriginSignatureByValue(param.GetSign())
	if err != nil {
		return nil, err
	}
	destroyPara := &abi.DestroyParam{
		Owner:    owner,
		Previous: previous,
		Token:    token,
		Amount:   amount,
		Sign:     sign,
	}
	r, err := b.blackHole.GetSendBlock(destroyPara)
	if err != nil {
		return nil, err
	}
	return toStateBlock(r), nil
}

func (b *BlackHoleAPI) GetRewardsBlock(ctx context.Context, param *pbtypes.Hash) (*pbtypes.StateBlock, error) {
	hash, err := toOriginHash(param)
	if err != nil {
		return nil, err
	}
	r, err := b.blackHole.GetRewardsBlock(&hash)
	if err != nil {
		return nil, err
	}
	return toStateBlock(r), nil
}

func (b *BlackHoleAPI) GetTotalDestroyInfo(ctx context.Context, param *pbtypes.Address) (*pbtypes.Balance, error) {
	addr, err := toOriginAddress(param)
	if err != nil {
		return nil, err
	}
	r, err := b.blackHole.GetTotalDestroyInfo(&addr)
	if err != nil {
		return nil, err
	}
	return toBalance(r), nil
}

func (b *BlackHoleAPI) GetDestroyInfoDetail(ctx context.Context, param *pbtypes.Address) (*pbtypes.DestroyInfos, error) {
	addr, err := toOriginAddress(param)
	if err != nil {
		return nil, err
	}
	r, err := b.blackHole.GetDestroyInfoDetail(&addr)
	if err != nil {
		return nil, err
	}
	infos := make([]*pbtypes.DestroyInfo, 0)
	for _, info := range r {
		t := new(pbtypes.DestroyInfo)
		t.Owner = toAddressValue(info.Owner)
		t.Token = toHashValue(info.Token)
		t.Previous = toHashValue(info.Previous)
		t.Amount = toBalanceValueByBigInt(info.Amount)
		t.TimeStamp = info.TimeStamp
		infos = append(infos, t)
	}
	result := new(pbtypes.DestroyInfos)
	result.Infos = infos
	return result, nil
}
