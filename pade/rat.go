// Copyright (c) 2017 Melvin Eloy Irizarry-Gelpí
// Licensed under the MIT License.

package pade

import "github.com/meirizarrygelpi/numbers/maclaurin"

// A Rat is a Padé approximant where each coefficient is a *big.Rat.
type Rat struct {
	P, Q maclaurin.Rat
}

// NewRat returns a new zero-valued Padé approximant.
func NewRat() *Rat {
	return &Rat{P: *maclaurin.NewRat(), Q: *maclaurin.NewRat()}
}
