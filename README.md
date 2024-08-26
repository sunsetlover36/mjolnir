# mjolnir
Go package to interact with EVM networks.

## ⚠️ Warning
The current implementation may not be suitable for production use!

## 📚 Example Usage

Here’s a basic example of how to use this package:

```go
package main

import (
	"fmt"
	"log"

	"github.com/sunsetlover36/mjolnir"
	"github.com/sunsetlover36/mjolnir/internal"
	"github.com/sunsetlover36/mjolnir/types"
)

func main() {
	// Convert your private key to an account
	account, err := mjolnir.PrivateKeyToAccount("YOUR_PRIVATE_KEY")
	if err != nil {
		log.Fatalf("Failed to convert private key to account: %v", err)
	}

	// Initialize a new wallet client with specified chain and RPC URL
	wc := mjolnir.NewWalletClient(types.NewWalletClientParams{
		Chain:   internal.Chains["Base"], // Set the correct chain
		RpcUrl:  "YOUR_RPC_URL",           // Replace with your actual RPC URL
		Account: account,
	})

	// Parse the ether value (in this case, 0.1 ETH)
	parsedEther, err := internal.ParseEther("0.1")
	if err != nil {
		log.Fatalf("Failed to parse ether value: %v", err)
	}

	// Send a transaction
	txHash, err := wc.SendTx(&types.TxInteractionParams{
		TxData: &types.TxData{
			To:    "TO_ADDRESS", // Replace with the recipient address
			Value: parsedEther,
		},
	})
	if err != nil {
		log.Fatalf("Failed to send transaction: %v", err)
	}
	fmt.Println("Transaction Hash:", txHash)

	// Read the balance from a token contract
	balance, err := wc.ReadContract(types.ReadContractParams{
		Address:      "TOKEN_ADDRESS", // Replace with the token contract address
		Abi:          TOKEN_ABI,       // Replace with the ABI of the token contract
		FunctionName: "balanceOf",     // Function to call on the contract
		Args:         []interface{}{account.Address},
	})
	if err != nil {
		log.Fatalf("Failed to read contract: %v", err)
	}
	fmt.Printf("Token Balance: %v\n", balance)
}
```

## 📖 Documentation
Documentation is coming soon! Stay tuned for detailed guides and examples.
