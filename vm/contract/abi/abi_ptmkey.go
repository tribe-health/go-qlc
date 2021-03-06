package abi

import (
	"fmt"
	"strings"

	"github.com/qlcchain/go-qlc/common"
	_ "github.com/qlcchain/go-qlc/common/statedb"
	"github.com/qlcchain/go-qlc/common/types"
	"github.com/qlcchain/go-qlc/common/util"
	"github.com/qlcchain/go-qlc/common/vmcontract/contractaddress"
	_ "github.com/qlcchain/go-qlc/config"
	"github.com/qlcchain/go-qlc/vm/abi"
	"github.com/qlcchain/go-qlc/vm/vmstore"
)

const (
	JsonPtmKey = `[
		{"type":"function","name":"PtmKeyDelete","inputs":[
			{"name":"btype","type":"uint16"}
		]},
		{"type":"function","name":"PtmKeyUpdate","inputs":[
			{"name":"btype","type":"uint16"},
			{"name":"pubkey","type":"string"}
		]},
		{"type":"variable","name":"PtmKeyStorageVar","inputs":[
			{"name":"pubkey","type":"string"},
			{"name":"valid","type":"bool"}
		]}
	]`

	MethodNamePtmKeyDelete       = "PtmKeyDelete"
	MethodNamePtmKeyUpdate       = "PtmKeyUpdate"
	VariableNamePtmKeyStorageVar = "PtmKeyStorageVar"
)

var (
	PtmKeyABI, _ = abi.JSONToABIContract(strings.NewReader(JsonPtmKey))
)

const (
	// contractAddr(32) + account(32) + btype(2) => pubkey + valid
	PtmKeyAccIndexS   = types.AddressSize + 1
	PtmKeyAccIndexE   = PtmKeyAccIndexS + 32
	PtmKeyBtypeIndexS = PtmKeyAccIndexE + 1
	PtmKeyBtypeIndexE = PtmKeyBtypeIndexS + 2
)

type PtmKeyInfo struct {
	Account types.Address `msg:"-" json:"account"`
	Btype   uint16        `msg:"b" json:"btype"`
	Pubkey  string        `msg:"k" json:"pubkey"`
}

type PtmKeyDeleteInfo struct {
	Account types.Address `msg:"-" json:"account"`
	Btype   uint16        `msg:"b" json:"btype"`
}

type PtmKeyStorage struct {
	Pubkey string `msg:"k" json:"pubkey"`
	Valid  bool   `msg:"v" json:"valid"`
}

func PtmKeyInfoCheck(ctx *vmstore.VMContext, pt uint16, pk []byte) error {
	switch pt {
	case common.PtmKeyVBtypeDefault, common.PtmKeyVBtypeA2p, common.PtmKeyVBtypeDod, common.PtmKeyVBtypeCloud:
		if len(pk) != 44 {
			return fmt.Errorf("pk len err")
		}
	default:
		return fmt.Errorf("PtmKeyInfoCheck type(%s) err", common.PtmKeyBtypeToString(pt))
	}

	return nil
}

func GetPtmKeyByAccountAndBtype(ctx *vmstore.VMContext, account types.Address, vBtype uint16) ([]*PtmKeyInfo, error) {
	pks := make([]*PtmKeyInfo, 0)

	var key []byte
	key = append(key, account[:]...)
	key = append(key, util.BE_Uint16ToBytes(vBtype)...)
	//fmt.Printf("GetPtmKeyByAccountAndBtype:get account(%s) key(%s)\n",account,string(key[:]))
	val, err := ctx.GetStorage(contractaddress.PtmKeyKVAddress[:], key)
	if err != nil {
		return nil, err
	}

	pkstorage := new(PtmKeyStorage)
	err = PtmKeyABI.UnpackVariable(pkstorage, VariableNamePtmKeyStorageVar, val)
	if err != nil || !pkstorage.Valid {
		return nil, err
	}
	pk := new(PtmKeyInfo)
	pk.Account = account
	pk.Btype = vBtype
	pk.Pubkey = pkstorage.Pubkey
	pks = append(pks, pk)
	return pks, nil
}

func GetPtmKeyByAccount(ctx *vmstore.VMContext, account types.Address) ([]*PtmKeyInfo, error) {
	pks := make([]*PtmKeyInfo, 0)

	itKey := append(contractaddress.PtmKeyKVAddress[:], account[:]...)
	btype := common.PtmKeyVBtypeDefault
	err := ctx.Iterator(itKey, func(key []byte, value []byte) error {
		var info PtmKeyStorage
		err := PtmKeyABI.UnpackVariable(&info, VariableNamePtmKeyStorageVar, value)
		if err != nil {
			return err
		}
		if info.Valid {
			pk := &PtmKeyInfo{
				Account: account,
				Btype:   btype,
				Pubkey:  info.Pubkey,
			}
			pks = append(pks, pk)
		}
		btype++
		return nil
	})
	if err != nil {
		return nil, err
	}
	if len(pks) == 0 {
		return nil, fmt.Errorf("no result")
	}
	return pks, nil
}
