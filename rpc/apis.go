package rpc

import (
	"github.com/qlcchain/go-qlc/rpc/api"
	rpc "github.com/qlcchain/jsonrpc2"
)

func (r *RPC) getApi(apiModule string) rpc.API {
	switch apiModule {
	case "qlcclassic":
		return rpc.API{
			Namespace: "qlcclassic",
			Version:   "1.0",
			Service:   api.NewQlcApi(r.ledger, r.eb),
			Public:    true,
		}
	case "account":
		return rpc.API{
			Namespace: "account",
			Version:   "1.0",
			Service:   api.NewAccountApi(),
			Public:    true,
		}
	case "ledger":
		return rpc.API{
			Namespace: "ledger",
			Version:   "1.0",
			Service:   api.NewLedgerApi(r.ledger, r.relation, r.eb),
			Public:    true,
		}
	case "net":
		return rpc.API{
			Namespace: "net",
			Version:   "1.0",
			Service:   api.NewNetApi(r.ledger, r.eb),
			Public:    true,
		}
	case "util":
		return rpc.API{
			Namespace: "util",
			Version:   "1.0",
			Service:   api.NewUtilApi(r.ledger),
			Public:    true,
		}
	case "wallet":
		return rpc.API{
			Namespace: "wallet",
			Version:   "1.0",
			Service:   api.NewWalletApi(r.ledger, r.wallet),
			Public:    true,
		}
	case "contract":
		return rpc.API{
			Namespace: "contract",
			Version:   "1.0",
			Service:   api.NewContractApi(r.ledger),
			Public:    true,
		}
	case "mintage":
		return rpc.API{
			Namespace: "mintage",
			Version:   "1.0",
			Service:   api.NewMintageApi(r.ledger),
			Public:    true,
		}
	case "pledge":
		return rpc.API{
			Namespace: "pledge",
			Version:   "1.0",
			Service:   api.NewNEP5PledgeApi(r.ledger),
			Public:    true,
		}
	case "sms":
		return rpc.API{
			Namespace: "sms",
			Version:   "1.0",
			Service:   api.NewSMSApi(r.ledger, r.relation),
			Public:    true,
		}
	case "rewards":
		return rpc.API{
			Namespace: "rewards",
			Version:   "1.0",
			Service:   api.NewRewardsApi(r.ledger),
			Public:    true,
		}
	case "pov":
		return rpc.API{
			Namespace: "pov",
			Version:   "1.0",
			Service:   api.NewPovApi(r.ledger, r.eb),
			Public:    true,
		}
	case "miner":
		return rpc.API{
			Namespace: "miner",
			Version:   "1.0",
			Service:   api.NewMinerApi(r.ledger),
			Public:    true,
		}
	default:
		return rpc.API{}
	}
}

func (r *RPC) GetApis(apiModule ...string) []rpc.API {
	var apis []rpc.API
	for _, m := range apiModule {
		apis = append(apis, r.getApi(m))
	}
	return apis
}

//In-proc apis
func (r *RPC) GetInProcessApis() []rpc.API {
	return r.GetPublicApis()
}

//Ipc apis
func (r *RPC) GetIpcApis() []rpc.API {
	return r.GetPublicApis()
}

//Http apis
func (r *RPC) GetHttpApis() []rpc.API {
	return r.GetPublicApis()
}

//WS apis
func (r *RPC) GetWSApis() []rpc.API {
	return r.GetPublicApis()
}

func (r *RPC) GetPublicApis() []rpc.API {
	apiModules := []string{"ledger", "account", "net", "util", "wallet", "mintage", "contract", "sms", "pledge", "rewards", "pov", "miner"}
	return r.GetApis(apiModules...)
}

func (r *RPC) GetAllApis() []rpc.API {
	//return GetApis("ledger", "wallet", "private_onroad", "net", "contract", "pledge", "register", "vote", "mintage", "consensusGroup", "testapi", "pow", "tx", "debug", "dashboard")
	return r.GetApis("qlc")
}
