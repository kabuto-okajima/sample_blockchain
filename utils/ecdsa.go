package utils

import (
	"fmt"
	"math/big"
)

// R is made with X coordinate of public key.
// S is made with senderPrivateKey and Transaction.
type Signature struct {
	R *big.Int
	S *big.Int
}

func (s *Signature) String() string {
	return fmt.Sprintf("%x%x", s.R, s.S)
}
