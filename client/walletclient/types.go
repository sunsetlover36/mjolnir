package walletclient

import (
	"github.com/sunsetlover36/mjolnir/internal"
	"github.com/sunsetlover36/mjolnir/types"
)

type WalletClient struct {
	client  *internal.RpcClient
	account *types.Account
}
