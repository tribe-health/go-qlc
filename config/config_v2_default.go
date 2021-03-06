// +build !testnet

/*
 * Copyright (c) 2019 QLC Chain Team
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package config

func DefaultConfigV2(dir string) (*ConfigV2, error) {
	pk, id, err := identityConfig()
	if err != nil {
		return nil, err
	}
	var cfg ConfigV2
	modules := []string{"qlcclassic", "ledger", "account", "net", "util", "wallet", "mintage", "contract", "sms"}
	cfg = ConfigV2{
		Version:             2,
		DataDir:             dir,
		StorageMax:          "10GB",
		AutoGenerateReceive: false,
		LogLevel:            "error",
		PerformanceEnabled:  false,
		RPC: &RPCConfigV2{
			Enable:           false,
			HTTPEnabled:      true,
			HTTPEndpoint:     "tcp4://0.0.0.0:9735",
			HTTPCors:         []string{"*"},
			HttpVirtualHosts: []string{},
			WSEnabled:        true,
			WSEndpoint:       "tcp4://0.0.0.0:9736",
			IPCEnabled:       true,
			IPCEndpoint:      defaultIPCEndpoint(dir),
			PublicModules:    modules,
		},
		P2P: &P2PConfigV2{
			BootNodes:          bootNodes,
			IsBootNode:         false,
			BootNodeHttpServer: bootNodeHttpServer,
			Listen:             "/ip4/0.0.0.0/tcp/9734",
			ListeningIp:        "127.0.0.1",
			SyncInterval:       120,
			Discovery: &DiscoveryConfigV2{
				DiscoveryInterval: 60,
				Limit:             2000,
				MDNSEnabled:       true,
				MDNSInterval:      30,
			},
			ID: &IdentityConfigV2{id, pk},
		},
	}

	return &cfg, nil
}

func defaultGRPCConfig() *GRPCConfig {
	return &GRPCConfig{
		Enable:                 false,
		ListenAddress:          "tcp://0.0.0.0:9746",
		HTTPEnable:             false,
		HTTPListenAddress:      "tcp://0.0.0.0:9745",
		MaxSubscriptionClients: 100,
	}
}

func defaultModules() []string {
	modules := []string{"ledger", "account", "net", "util", "mintage", "contract", "pledge",
		"rewards", "pov", "miner", "config", "debug", "destroy", "metrics", "rep", "chain", "dpki",
		"permission", "privacy", "ptmkey"}
	return modules
}
