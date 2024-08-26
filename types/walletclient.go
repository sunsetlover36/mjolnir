package types

type NewWalletClientParams struct {
	RpcUrl  string
	Chain   Chain
	Account *Account
}
