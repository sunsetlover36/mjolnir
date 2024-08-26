package mjolnir

import (
	"github.com/sunsetlover36/mjolnir/client/publicclient"
	"github.com/sunsetlover36/mjolnir/client/walletclient"
	"github.com/sunsetlover36/mjolnir/types"
)

var Chains = map[string]types.Chain{
	"Base": {
		Id:     8453,
		Name:   "Base Mainnet",
		RpcUrl: "https://base.llamarpc.com",
	},
	"Ethereum": {
		Id:     1,
		Name:   "Ethereum Mainnet",
		RpcUrl: "https://eth.llamarpc.com",
	},
	"Polygon": {
		Id:     137,
		Name:   "Polygon Mainnet",
		RpcUrl: "https://polygon.llamarpc.com",
	},
}

func NewPublicClient(params types.NewPublicClientParams) *publicclient.PublicClient {
	return publicclient.NewPublicClient(params)
}

func NewWalletClient(params types.NewWalletClientParams) *walletclient.WalletClient {
	return walletclient.NewWalletClient(params)
}
