package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sunsetlover36/mjolnir/types"
)

type RpcClient struct {
	rpcUrl string
	chain  types.Chain
}

func NewRpcClient(params types.NewRpcClientParams) *RpcClient {
	return &RpcClient{chain: params.Chain, rpcUrl: params.RpcUrl}
}

func (c *RpcClient) Call(method string, params interface{}) (json.RawMessage, error) {
	requestBody := types.RpcRequest{
		Jsonrpc: "2.0",
		Method:  method,
		Params:  params,
		Id:      1,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request: %w", err)
	}

	resp, err := http.Post(c.rpcUrl, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	var response types.RpcResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	if response.Error != nil {
		return nil, fmt.Errorf("rpc error: %s", response.Error.Message)
	}

	return response.Result, nil
}
