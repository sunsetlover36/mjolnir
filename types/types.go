package types

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"

	ethTypes "github.com/ethereum/go-ethereum/core/types"
)

type Chain struct {
	Id     int64
	Name   string
	RpcUrl string
}

type NewRpcClientParams struct {
	RpcUrl string
	Chain  Chain
}

type RpcRequest struct {
	Jsonrpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Id      int         `json:"id"`
}

type RpcResponse struct {
	Jsonrpc string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result"`
	Error   *RpcError       `json:"error"`
	Id      int             `json:"id"`
}

type RpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type GetBlockTransactionCountParams struct {
	BlockHash   *string  `json:"blockHash"`
	BlockNumber *big.Int `json:"blockNumber"`
	BlockTag    *string  `json:"blockTag"`
}
type GetBlockParams struct {
	BlockHash   *string  `json:"blockHash"`
	BlockNumber *big.Int `json:"blockNumber"`
	BlockTag    *string  `json:"blockTag"`
}

// eth_call params
type CallParams struct {
	From     string   `json:"from"`
	To       string   `json:"to"`
	Gas      uint64   `json:"gas,omitempty"`
	GasPrice *big.Int `json:"gasPrice,omitempty"`
	Value    *big.Int `json:"value"`
	Data     []byte   `json:"data"`
}

func (c CallParams) MarshalJSON() ([]byte, error) {
	type Alias CallParams
	return json.Marshal(&struct {
		Gas      string `json:"gas,omitempty"`
		GasPrice string `json:"gasPrice,omitempty"`
		Value    string `json:"value,omitempty"`
		Data     string `json:"data,omitempty"`
		*Alias
	}{
		Gas: func() string {
			if c.Gas > 0 {
				return fmt.Sprintf("0x%x", c.Gas)
			}
			return ""
		}(),
		GasPrice: func() string {
			if c.GasPrice != nil {
				return fmt.Sprintf("0x%x", c.GasPrice)
			}
			return ""
		}(),
		Value: func() string {
			if c.Value != nil {
				return fmt.Sprintf("0x%x", c.Value)
			}
			return ""
		}(),
		Data:  fmt.Sprintf("0x%x", c.Data),
		Alias: (*Alias)(&c),
	})
}

// --------

type ReadContractParams struct {
	Address      string
	Abi          string
	FunctionName string
	Args         []interface{}
}
type ContractInteractionParams struct {
	Address              string
	Abi                  string
	FunctionName         string
	Args                 []interface{}
	MaxFeePerGas         *big.Int
	MaxPriorityFeePerGas *big.Int
	GasLimit             uint64
	Value                *big.Int
	Nonce                uint64
	Account              *Account
}

type TxInteractionParams struct {
	TxData  *TxData
	Account *Account
}
type TxData struct {
	ChainId              *big.Int
	Nonce                uint64
	MaxFeePerGas         *big.Int
	MaxPriorityFeePerGas *big.Int
	Gas                  uint64
	To                   string
	Value                *big.Int
	Data                 []byte
}
type SendTxOptions struct {
	Simulate bool
}
type SimulateTxResult struct {
	Tx     *ethTypes.Transaction
	Result string
}

type FeeHistoryResult struct {
	OldestBlock   string
	Reward        [][]string
	BaseFeePerGas []string
	GasUsedRatio  []float64
}

type RawTransaction struct {
	Hash             string `json:"hash"`
	From             string `json:"from"`
	To               string `json:"to,omitempty"`
	Value            string `json:"value"`
	GasPrice         string `json:"gasPrice"`
	Gas              string `json:"gas"`
	Nonce            string `json:"nonce"`
	TransactionIndex string `json:"transactionIndex"`
}
type Transaction struct {
	Hash             string   `json:"hash"`
	From             string   `json:"from"`
	To               string   `json:"to,omitempty"`
	Value            *big.Int `json:"value"`
	GasPrice         *big.Int `json:"gasPrice"`
	Gas              uint64   `json:"gas"`
	Nonce            string   `json:"nonce"`
	TransactionIndex uint64   `json:"transactionIndex"`
}

type RawBlock struct {
	Number           string           `json:"number"`
	Hash             string           `json:"hash"`
	ParentHash       string           `json:"parentHash"`
	Nonce            string           `json:"nonce"`
	Sha3Uncles       string           `json:"sha3Uncles"`
	LogsBloom        string           `json:"logsBloom"`
	TransactionsRoot string           `json:"transactionsRoot"`
	StateRoot        string           `json:"stateRoot"`
	Miner            string           `json:"miner"`
	Difficulty       string           `json:"difficulty"`
	TotalDifficulty  string           `json:"totalDifficulty"`
	ExtraData        string           `json:"extraData"`
	Size             string           `json:"size"`
	GasLimit         string           `json:"gasLimit"`
	GasUsed          string           `json:"gasUsed"`
	Timestamp        string           `json:"timestamp"`
	Transactions     []RawTransaction `json:"transactions"`
	Uncles           []string         `json:"uncles"`
}
type Block struct {
	Number           *big.Int      `json:"number"`
	Hash             string        `json:"hash"`
	ParentHash       string        `json:"parentHash"`
	Nonce            string        `json:"nonce"`
	Sha3Uncles       string        `json:"sha3Uncles"`
	LogsBloom        string        `json:"logsBloom"`
	TransactionsRoot string        `json:"transactionsRoot"`
	StateRoot        string        `json:"stateRoot"`
	Miner            string        `json:"miner"`
	Difficulty       *big.Int      `json:"difficulty"`
	TotalDifficulty  *big.Int      `json:"totalDifficulty"`
	ExtraData        string        `json:"extraData"`
	Size             *big.Int      `json:"size"`
	GasLimit         *big.Int      `json:"gasLimit"`
	GasUsed          *big.Int      `json:"gasUsed"`
	Timestamp        *big.Int      `json:"timestamp"`
	Transactions     []Transaction `json:"transactions"`
	Uncles           []string      `json:"uncles"`
}

type Account struct {
	Address    string
	PrivateKey *ecdsa.PrivateKey
}
