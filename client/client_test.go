package client

import (
	"context"
	"os"
	"testing"

	clientModel "github.com/6boris/web3-go/model/client"
)

var (
	testCtx              context.Context
	testPool             *Pool
	testEvmEthClient     *EvmClient
	testEvmPolygonClient *EvmClient
	testSolanaClient     *SolanaClient
)

func TestMain(m *testing.M) {
	var err error
	testEvmEthClient, err = NewEvmClient(&clientModel.ConfEvmChainClient{TransportURL: "https://1rpc.io/eth"})
	if err != nil {
		panic(err)
	}
	testEvmPolygonClient, err = NewEvmClient(&clientModel.ConfEvmChainClient{TransportURL: "https://1rpc.io/matic"})
	if err != nil {
		panic(err)
	}
	testSolanaClient, err = NewSolanaClient(&clientModel.ConfSolanaClient{TransportURL: "https://api.mainnet-beta.solana.com"})
	if err != nil {
		panic(err)
	}
	testPool = NewPool(clientModel.GetDefaultConfPool())
	testCtx = context.TODO()
	// Before Test
	code := m.Run()
	// After Test
	os.Exit(code)
}
