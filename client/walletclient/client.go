package walletclient

import (
	"github.com/sunsetlover36/mjolnir/internal"
	"github.com/sunsetlover36/mjolnir/types"
)

func NewWalletClient(params types.NewWalletClientParams) *WalletClient {
	return &WalletClient{
		client: internal.NewRpcClient(types.NewRpcClientParams{
			Chain:  params.Chain,
			RpcUrl: params.RpcUrl,
		}),
		account: params.Account,
	}
}
