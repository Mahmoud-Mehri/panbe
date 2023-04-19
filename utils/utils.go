package utils

import (
	"fmt"
	"math/big"
)

type Signature struct {
	R *big.Int
	S *big.Int
}

// String method will get called automatically by fmt
func (s *Signature) String() string { 
	return fmt.Sprintf("%x%x", s.R, s.S)
}