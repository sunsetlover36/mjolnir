package publicclient

import (
	"math/big"

	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/sunsetlover36/mjolnir/types"
)

func (c *PublicClient) GetBlock(params types.GetBlockParams) (*types.Block, error) {
	return c.client.GetBlock(params)
}

func (c *PublicClient) GetBlockNumber() (uint64, error) {
	return c.client.GetBlockNumber()
}

func (c *PublicClient) GetBlockTransactionCount(params types.GetBlockTransactionCountParams) (uint64, error) {
	return c.client.GetBlockTransactionCount(params)
}

func (c *PublicClient) GetBalance(address string) (*big.Int, error) {
	return c.client.GetBalance(address)
}

func (c *PublicClient) GetTransactionCount(address string) (uint64, error) {
	return c.client.GetTransactionCount(address)
}

func (c *PublicClient) GetGasPrice() (*big.Int, error) {
	return c.client.GetGasPrice()
}
func (c *PublicClient) GetMaxPriorityFeePerGas() (*big.Int, error) {
	return c.client.GetMaxPriorityFeePerGas()
}
func (c *PublicClient) EstimateGas(params types.CallParams) (*big.Int, error) {
	return c.client.EstimateGas(params)
}

func (c *PublicClient) PrepareTxRequest(params types.TxInteractionParams) (*ethTypes.Transaction, error) {
	return c.client.PrepareTxRequest(params)
}
func (c *PublicClient) SimulateTx(params types.TxInteractionParams) (*types.SimulateTxResult, error) {
	return c.client.SimulateTx(params)
}

func (c *PublicClient) ReadContract(params types.ReadContractParams) ([]byte, error) {
	return c.client.ReadContract(params)
}
func (c *PublicClient) SimulateContract(params types.ContractInteractionParams) (*types.SimulateTxResult, error) {
	return c.client.SimulateContract(params)
}
