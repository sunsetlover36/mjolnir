package mjolnir

import (
	"math/big"

	"github.com/sunsetlover36/mjolnir/internal"
	"github.com/sunsetlover36/mjolnir/types"
)

func GeneratePrivateKey() (string, error) {
	return internal.GeneratePrivateKey()
}
func PrivateKeyToAccount(privateKeyHex string) (*types.Account, error) {
	return internal.PrivateKeyToAccount(privateKeyHex)
}
func ParseEther(etherStr string) (*big.Int, error) {
	return internal.ParseEther(etherStr)
}
func FormatEther(wei *big.Int) string {
	return internal.FormatEther(wei)
}
