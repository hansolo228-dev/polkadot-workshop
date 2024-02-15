package blockchain

import (
	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
)

func BlockNumber(rpc string) int64 {
	api, err := gsrpc.NewSubstrateAPI(rpc)
	if err != nil {
		panic(err)
	}
	header, err := api.RPC.Chain.GetHeaderLatest()
	if err != nil {
		panic(err)
	}
	return int64(header.Number)
}

func NodeName(rpc string) string {
	api, err := gsrpc.NewSubstrateAPI(rpc)
	if err != nil {
		panic(err)
	}
	nodeName, err := api.RPC.System.Name()
	if err != nil {
		panic(err)
	}
	return string(nodeName)
}

func NodeVersion(rpc string) string {
	api, err := gsrpc.NewSubstrateAPI(rpc)
	if err != nil {
		panic(err)
	}
	nodeVersion, err := api.RPC.System.Version()
	if err != nil {
		panic(err)
	}
	return string(nodeVersion)
}
