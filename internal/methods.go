package internal

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/sunsetlover36/mjolnir/types"
)

func (c *RpcClient) GetBlock(params types.GetBlockParams) (*types.Block, error) {
	var rpcMethod string
	var rpcParams []interface{}
	if params.BlockHash != nil {
		rpcMethod = "eth_getBlockByHash"
		rpcParams = append(rpcParams, params.BlockHash)
	} else if params.BlockNumber != nil {
		rpcMethod = "eth_getBlockByNumber"
		rpcParams = append(rpcParams, fmt.Sprintf("0x%x", params.BlockNumber))
	} else if params.BlockTag != nil {
		rpcMethod = "eth_getBlockByNumber"
		rpcParams = append(rpcParams, params.BlockTag)
	} else {
		rpcMethod = "eth_getBlockByNumber"
		rpcParams = append(rpcParams, "latest")
	}

	rpcParams = append(rpcParams, true)

	result, err := c.Call(rpcMethod, rpcParams)
	if err != nil {
		return nil, err
	}

	var rawBlock types.RawBlock
	if err := json.Unmarshal(result, &rawBlock); err != nil {
		return nil, fmt.Errorf("failed to unmarshal rawBlock: %v", err)
	}

	block := types.Block{
		Number:           HexToBigInt(rawBlock.Number),
		Difficulty:       HexToBigInt(rawBlock.Difficulty),
		TotalDifficulty:  HexToBigInt(rawBlock.TotalDifficulty),
		Size:             HexToBigInt(rawBlock.Size),
		GasLimit:         HexToBigInt(rawBlock.GasLimit),
		GasUsed:          HexToBigInt(rawBlock.GasUsed),
		Timestamp:        HexToBigInt(rawBlock.Timestamp),
		Hash:             rawBlock.Hash,
		ParentHash:       rawBlock.ParentHash,
		Nonce:            rawBlock.Nonce,
		Sha3Uncles:       rawBlock.Sha3Uncles,
		LogsBloom:        rawBlock.LogsBloom,
		TransactionsRoot: rawBlock.TransactionsRoot,
		StateRoot:        rawBlock.StateRoot,
		Miner:            rawBlock.Miner,
		ExtraData:        rawBlock.ExtraData,
		Uncles:           rawBlock.Uncles,
	}
	for _, rawTx := range rawBlock.Transactions {
		tx := ConvertRawTransaction(rawTx)
		block.Transactions = append(block.Transactions, tx)
	}

	return &block, nil
}

func (c *RpcClient) GetBlockNumber() (uint64, error) {
	result, err := c.Call("eth_blockNumber", []interface{}{})
	if err != nil {
		return 0, err
	}

	var blockNumberHex string
	if err := json.Unmarshal(result, &blockNumberHex); err != nil {
		return 0, fmt.Errorf("failed to unmarshal blockNumberHex: %v", err)
	}

	blockNumber, err := strconv.ParseUint(blockNumberHex, 0, 64)
	if err != nil {
		return 0, fmt.Errorf("error converting block number to int: %w", err)
	}

	return blockNumber, nil
}

func (c *RpcClient) GetBlockTransactionCount(params types.GetBlockTransactionCountParams) (uint64, error) {
	var rpcParams []interface{}
	if params.BlockHash != nil {
		rpcParams = append(rpcParams, params.BlockHash)
	} else if params.BlockNumber != nil {
		rpcParams = append(rpcParams, fmt.Sprintf("0x%x", params.BlockNumber))
	} else if params.BlockTag != nil {
		rpcParams = append(rpcParams, params.BlockTag)
	} else {
		rpcParams = append(rpcParams, "latest")
	}

	result, err := c.Call("eth_getBlockTransactionCountByHash", rpcParams)
	if err != nil {
		return 0, err
	}

	var txCountHex string
	if err := json.Unmarshal(result, &txCountHex); err != nil {
		return 0, fmt.Errorf("failed to unmarshal txCountHex: %v", err)
	}

	txCount, err := strconv.ParseUint(txCountHex, 0, 64)
	if err != nil {
		return 0, fmt.Errorf("error converting transaction count to int: %w", err)
	}

	return txCount, nil
}

func (c *RpcClient) GetBalance(address string) (*big.Int, error) {
	result, err := c.Call("eth_getBalance", []interface{}{address, "latest"})
	if err != nil {
		return nil, err
	}

	var balanceHex string
	if err := json.Unmarshal(result, &balanceHex); err != nil {
		return nil, fmt.Errorf("failed to unmarshal balanceHex: %v", err)
	}

	balance := new(big.Int)
	balance.SetString(balanceHex[2:], 16)

	return balance, nil
}

func (c *RpcClient) GetTransactionCount(address string) (uint64, error) {
	result, err := c.Call("eth_getTransactionCount", []interface{}{address, "latest"})
	if err != nil {
		return 0, err
	}

	var transactionCountHex string
	if err := json.Unmarshal(result, &transactionCountHex); err != nil {
		return 0, fmt.Errorf("failed to unmarshal transactionCountHex: %v", err)
	}

	transactionCount, err := strconv.ParseUint(transactionCountHex, 0, 64)
	if err != nil {
		return 0, fmt.Errorf("error converting transaction count to int: %w", err)
	}

	return transactionCount, nil
}

func (c *RpcClient) GetGasPrice() (*big.Int, error) {
	result, err := c.Call("eth_gasPrice", []interface{}{})
	if err != nil {
		return nil, fmt.Errorf("failed to get gas price: %w", err)
	}

	var gasPriceHex string
	if err := json.Unmarshal(result, &gasPriceHex); err != nil {
		return nil, fmt.Errorf("failed to unmarshal gasPriceHex: %v", err)
	}

	gasPrice := new(big.Int)
	gasPrice.SetString(gasPriceHex[2:], 16)

	return gasPrice, nil
}
func (c *RpcClient) GetMaxPriorityFeePerGas() (*big.Int, error) {
	params := []interface{}{
		"0xA",
		"latest",
		[]float64{90.0},
	}

	result, err := c.Call("eth_feeHistory", params)
	if err != nil {
		return nil, fmt.Errorf("failed to get fee history: %w", err)
	}

	var feeHistory types.FeeHistoryResult
	if err := json.Unmarshal(result, &feeHistory); err != nil {
		return nil, fmt.Errorf("failed to unmarshal feeHistory: %v", err)
	}

	lastReward := feeHistory.Reward[len(feeHistory.Reward)-1]
	suggestedPriorityFeeStr := lastReward[0]

	suggestedPriorityFee := new(big.Int)
	suggestedPriorityFee.SetString(suggestedPriorityFeeStr[2:], 16)

	return suggestedPriorityFee, nil
}
func (c *RpcClient) EstimateGas(params types.CallParams) (*big.Int, error) {
	result, err := c.Call("eth_estimateGas", []types.CallParams{params})
	if err != nil {
		return nil, fmt.Errorf("failed to estimate gas: %v", err)
	}

	var estimatedGasHex string
	if err := json.Unmarshal(result, &estimatedGasHex); err != nil {
		return nil, fmt.Errorf("failed to unmarshal estimatedGasHex: %v", err)
	}

	estimatedGas := new(big.Int)
	estimatedGas.SetString(estimatedGasHex[2:], 16)

	return estimatedGas, nil
}

func (c *RpcClient) PrepareTxRequest(params types.TxInteractionParams) (*ethTypes.Transaction, error) {
	toAddress := common.HexToAddress(params.TxData.To)

	nonce := params.TxData.Nonce
	if nonce == 0 {
		fetchedNonce, err := c.GetTransactionCount(params.Account.Address)
		if err != nil {
			return nil, err
		}
		nonce = fetchedNonce
	}

	gasTipCap := params.TxData.MaxPriorityFeePerGas
	if gasTipCap == nil {
		fetchedGasTipCap, err := c.GetMaxPriorityFeePerGas()
		if err != nil {
			return nil, err
		}
		gasTipCap = fetchedGasTipCap
	}

	gasFeeCap := params.TxData.MaxFeePerGas
	if gasFeeCap == nil {
		fetchedGasFeeCap, err := c.GetGasPrice()
		if err != nil {
			return nil, err
		}
		gasFeeCap = fetchedGasFeeCap.Add(fetchedGasFeeCap, gasTipCap)
	}

	gasLimit := params.TxData.Gas
	if gasLimit == 0 {
		estimatedGas, err := c.EstimateGas(types.CallParams{
			From:     params.Account.Address,
			To:       params.TxData.To,
			Gas:      params.TxData.Gas,
			GasPrice: gasFeeCap,
			Value:    params.TxData.Value,
			Data:     params.TxData.Data,
		})
		if err != nil {
			return nil, err
		}
		gasLimit = estimatedGas.Uint64()
	}

	dynamicFeeTx := ethTypes.DynamicFeeTx{
		ChainID:   params.TxData.ChainId,
		Nonce:     nonce,
		Gas:       gasLimit,
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		To:        &toAddress,
		Value:     params.TxData.Value,
		Data:      params.TxData.Data,
	}
	tx := ethTypes.NewTx(&dynamicFeeTx)

	if params.Account != nil {
		chainId := big.NewInt(c.chain.Id)
		signedTx, err := ethTypes.SignTx(tx, ethTypes.LatestSignerForChainID(chainId), params.Account.PrivateKey)
		if err != nil {
			return nil, fmt.Errorf("failed to sign transaction: %v", err)
		}
		tx = signedTx
	}

	return tx, nil
}

func (c *RpcClient) SimulateTx(params types.TxInteractionParams) (*types.SimulateTxResult, error) {
	tx, err := c.PrepareTxRequest(types.TxInteractionParams{
		TxData:  params.TxData,
		Account: params.Account,
	})
	if err != nil {
		return nil, err
	}

	callParams := types.CallParams{
		To:       params.TxData.To,
		Value:    params.TxData.Value,
		Data:     params.TxData.Data,
		Gas:      params.TxData.Gas,
		GasPrice: params.TxData.MaxFeePerGas,
	}
	callParamsJson, err := json.Marshal(callParams)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal callParams: %v", err)
	}

	result, err := c.Call("eth_call", []interface{}{json.RawMessage(callParamsJson), "latest"})
	if err != nil {
		return nil, fmt.Errorf("simulation failed: %w", err)
	}

	var simulationResult string
	if err := json.Unmarshal(result, &simulationResult); err != nil {
		return nil, fmt.Errorf("failed to unmarshal simulation result: %v", err)
	}

	return &types.SimulateTxResult{
		Tx:     tx,
		Result: simulationResult,
	}, nil
}
func (c *RpcClient) SendTx(params types.TxInteractionParams) (string, error) {
	tx, err := c.PrepareTxRequest(types.TxInteractionParams{
		TxData:  params.TxData,
		Account: params.Account,
	})
	if err != nil {
		return "", err
	}

	data, err := tx.MarshalBinary()
	if err != nil {
		return "", fmt.Errorf("failed to marshal signed transaction: %w", err)
	}

	result, err := c.Call("eth_sendRawTransaction", []interface{}{hexutil.Encode(data)})
	if err != nil {
		return "", fmt.Errorf("failed to send transaction: %w", err)
	}

	var txHash string
	if err := json.Unmarshal(result, &txHash); err != nil {
		return "", fmt.Errorf("failed to unmarshal txHash: %v", err)
	}

	return txHash, nil
}

func (c *RpcClient) ReadContract(params types.ReadContractParams) ([]byte, error) {
	parsedABI, err := abi.JSON(strings.NewReader(params.Abi))
	if err != nil {
		return nil, fmt.Errorf("failed to parse ABI: %v", err)
	}

	convertedArgs := convertArgs(params.Args)
	data, err := parsedABI.Pack(params.FunctionName, convertedArgs...)
	if err != nil {
		return nil, fmt.Errorf("failed to pack arguments: %v", err)
	}

	payload := map[string]interface{}{
		"to":   params.Address,
		"data": "0x" + hex.EncodeToString(data),
	}

	result, err := c.Call("eth_call", []interface{}{payload, "latest"})
	if err != nil {
		return nil, fmt.Errorf("failed to call contract: %v", err)
	}

	var calldata string
	if err := json.Unmarshal(result, &calldata); err != nil {
		return nil, fmt.Errorf("failed to unmarshal calldata: %v", err)
	}

	var output interface{}
	err = parsedABI.UnpackIntoInterface(&output, params.FunctionName, common.FromHex(calldata))
	if err != nil {
		return nil, fmt.Errorf("failed to unpack result: %v", err)
	}

	outputBytes, err := json.Marshal(output)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal result: %v", err)
	}

	return outputBytes, nil
}
func (c *RpcClient) WriteContract(params types.ContractInteractionParams) (string, error) {
	if params.Account == nil {
		return "", fmt.Errorf("account with private key is required to sign the transaction")
	}

	parsedABI, err := abi.JSON(strings.NewReader(params.Abi))
	if err != nil {
		return "", fmt.Errorf("failed to parse ABI: %v", err)
	}

	data, err := parsedABI.Pack(params.FunctionName, params.Args...)
	if err != nil {
		return "", fmt.Errorf("failed to pack arguments: %v", err)
	}

	toAddress := common.HexToAddress(params.Address)

	txData := &types.TxData{
		To:    toAddress.Hex(),
		Value: params.Value,
		Nonce: params.Nonce,
		Data:  data,
	}

	txHash, err := c.SendTx(types.TxInteractionParams{
		TxData:  txData,
		Account: params.Account,
	})
	if err != nil {
		return "", fmt.Errorf("failed to send transaction: %v", err)
	}

	return txHash, nil
}
func (c *RpcClient) SimulateContract(params types.ContractInteractionParams) (*types.SimulateTxResult, error) {
	parsedABI, err := abi.JSON(strings.NewReader(params.Abi))
	if err != nil {
		return nil, fmt.Errorf("failed to parse ABI: %v", err)
	}

	data, err := parsedABI.Pack(params.FunctionName, params.Args...)
	if err != nil {
		return nil, fmt.Errorf("failed to pack arguments: %v", err)
	}

	toAddress := common.HexToAddress(params.Address)

	txData := &types.TxData{
		To:    toAddress.Hex(),
		Value: params.Value,
		Data:  data,
	}

	simulationResult, err := c.SimulateTx(types.TxInteractionParams{
		TxData:  txData,
		Account: params.Account,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to send transaction: %v", err)
	}

	return simulationResult, nil
}
