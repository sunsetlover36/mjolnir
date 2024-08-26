package internal

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sunsetlover36/mjolnir/types"
)

func HexToBigInt(hexStr string) *big.Int {
	bigInt := new(big.Int)
	bigInt.SetString(strings.TrimPrefix(hexStr, "0x"), 16)

	return bigInt
}

func HexToUint64(hexStr string) uint64 {
	value, err := strconv.ParseUint(hexStr, 0, 64)
	if err != nil {
		panic(err)
	}
	return value
}

func ConvertRawTransaction(rawTx types.RawTransaction) types.Transaction {
	return types.Transaction{
		Hash:     rawTx.Hash,
		From:     rawTx.From,
		To:       rawTx.To,
		Nonce:    rawTx.Nonce,
		Value:    HexToBigInt(rawTx.Value),
		GasPrice: HexToBigInt(rawTx.GasPrice),
		Gas:      HexToUint64(rawTx.Gas),
	}
}

func GeneratePrivateKeyEcdsa() (*ecdsa.PrivateKey, error) {
	privateKey, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}
func GeneratePrivateKey() (string, error) {
	privateKey, err := GeneratePrivateKeyEcdsa()
	if err != nil {
		return "", err
	}
	privateKeyBytes := privateKey.D.Bytes()
	return "0x" + hex.EncodeToString(privateKeyBytes), nil
}
func PrivateKeyToAccount(privateKeyHex string) (*types.Account, error) {
	privateKeyBytes, err := hex.DecodeString(strings.TrimPrefix(privateKeyHex, "0x"))
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	address := crypto.PubkeyToAddress(*publicKey).Hex()

	return &types.Account{
		Address:    address,
		PrivateKey: privateKey,
	}, nil
}

func convertArgs(args []interface{}) []interface{} {
	convertedArgs := make([]interface{}, len(args))

	for i, arg := range args {
		switch v := arg.(type) {
		case string:
			if len(v) == 42 && v[:2] == "0x" {
				convertedArgs[i] = common.HexToAddress(v)
			} else {
				convertedArgs[i] = arg
			}
		default:
			convertedArgs[i] = arg
		}
	}

	return convertedArgs
}

func ParseEther(etherStr string) (*big.Int, error) {
	parts := strings.Split(etherStr, ".")
	if len(parts) > 2 {
		return nil, fmt.Errorf("invalid ether string")
	}

	wholePart := parts[0]
	fractionalPart := ""
	if len(parts) == 2 {
		fractionalPart = parts[1]
	}

	fractionalPart = fractionalPart + strings.Repeat("0", 18-len(fractionalPart))
	etherCombined := wholePart + fractionalPart

	wei := new(big.Int)
	_, ok := wei.SetString(etherCombined, 10)
	if !ok {
		return nil, fmt.Errorf("failed to parse ether")
	}

	return wei, nil
}

func FormatEther(wei *big.Int) string {
	weiStr := wei.String()

	if len(weiStr) <= 18 {
		weiStr = strings.Repeat("0", 18-len(weiStr)) + weiStr
	}

	wholePart := weiStr[:len(weiStr)-18]
	fractionalPart := weiStr[len(weiStr)-18:]

	fractionalPart = strings.TrimRight(fractionalPart, "0")

	if fractionalPart == "" {
		return wholePart
	}

	return wholePart + "." + fractionalPart
}

func ParseGwei(gweiStr string) (*big.Int, error) {
	gwei := new(big.Float)
	gwei, ok := gwei.SetString(gweiStr)
	if !ok {
		return nil, fmt.Errorf("failed to parse Gwei value: %s", gweiStr)
	}

	wei := new(big.Float).Mul(gwei, big.NewFloat(1e9))
	weiInt := new(big.Int)
	wei.Int(weiInt)

	return weiInt, nil
}

func FormatGwei(wei *big.Int) string {
	weiFloat := new(big.Float).SetInt(wei)
	gwei := new(big.Float).Quo(weiFloat, big.NewFloat(1e9))

	return gwei.Text('f', -1)
}
