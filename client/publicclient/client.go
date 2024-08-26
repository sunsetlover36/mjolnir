package publicclient

import (
	"github.com/sunsetlover36/mjolnir/internal"
	"github.com/sunsetlover36/mjolnir/types"
)

func NewPublicClient(params types.NewPublicClientParams) *PublicClient {
	return &PublicClient{
		client: internal.NewRpcClient(types.NewRpcClientParams{
			RpcUrl: params.RpcUrl,
		}),
	}
}
