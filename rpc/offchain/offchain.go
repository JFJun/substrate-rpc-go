package offchain

import "github.com/JFJun/substrate-rpc-go/client"

// Offchain exposes methods for retrieval of off-chain data
type Offchain struct {
	client client.Client
}

// NewOffchain creates a new Offchain struct
func NewOffchain(c client.Client) *Offchain {
	return &Offchain{client: c}
}
