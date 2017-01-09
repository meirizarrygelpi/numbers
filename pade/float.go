// Copyright (c) 2017 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package pade

import "github.com/meirizarrygelpi/numbers/maclaurin"

// A Float is a Padé approximant where each coefficient is a *big.Float.
type Float struct {
	P, Q maclaurin.Float
}

// NewFloat returns a new zero-valued Padé approximant.
func NewFloat() *Float {
	return &Float{P: *maclaurin.NewFloat(), Q: *maclaurin.NewFloat()}
}
