// Copyright (c) 2017 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package pade

import "github.com/meirizarrygelpi/numbers/maclaurin"

// An Int is a Padé approximant where each coefficient is a *big.Int.
type Int struct {
	P, Q maclaurin.Int
}

// NewInt returns a new zero-valued Padé approximant.
func NewInt() *Int {
	return &Int{P: *maclaurin.NewInt(), Q: *maclaurin.NewInt()}
}
