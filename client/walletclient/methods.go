package walletclient

import (
	"math/big"

	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/sunsetlover36/mjolnir/types"
)

func (c *WalletClient) GetBlock(params types.GetBlockParams) (*types.Block, error) {
	return c.client.GetBlock(params)
}

func (c *WalletClient) GetBlockNumber() (uint64, error) {
	return c.client.GetBlockNumber()
}

func (c *WalletClient) GetBlockTransactionCount(params types.GetBlockTransactionCountParams) (uint64, error) {
	return c.client.GetBlockTransactionCount(params)
}

func (c *WalletClient) GetBalance() (*big.Int, error) {
	return c.client.GetBalance(c.account.Address)
}

func (c *WalletClient) GetTransactionCount() (uint64, error) {
	return c.client.GetTransactionCount(c.account.Address)
}

func (c *WalletClient) GetGasPrice() (*big.Int, error) {
	return c.client.GetGasPrice()
}
func (c *WalletClient) GetMaxPriorityFeePerGas() (*big.Int, error) {
	return c.client.GetMaxPriorityFeePerGas()
}
func (c *WalletClient) EstimateGas(params types.CallParams) (*big.Int, error) {
	return c.client.EstimateGas(params)
}

func (c *WalletClient) PrepareTxRequest(params types.TxInteractionParams) (*ethTypes.Transaction, error) {
	return c.client.PrepareTxRequest(params)
}
func (c *WalletClient) SimulateTx(params types.TxInteractionParams) (*types.SimulateTxResult, error) {
	params.Account = c.account
	return c.client.SimulateTx(params)
}
func (c *WalletClient) SendTx(params *types.TxInteractionParams) (string, error) {
	params.Account = c.account
	return c.client.SendTx(*params)
}

func (c *WalletClient) ReadContract(params types.ReadContractParams) ([]byte, error) {
	return c.client.ReadContract(params)
}
func (c *WalletClient) WriteContract(params types.ContractInteractionParams) (string, error) {
	params.Account = c.account
	return c.client.WriteContract(params)
}
func (c *WalletClient) SimulateContract(params types.ContractInteractionParams) (*types.SimulateTxResult, error) {
	params.Account = c.account
	return c.client.SimulateContract(params)
}
