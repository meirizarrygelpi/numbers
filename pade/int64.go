// Copyright (c) 2017 Melvin Eloy Irizarry-Gelpí
// Licensed under the MIT License.

package pade

import "github.com/meirizarrygelpi/numbers/maclaurin"

// An Int64 is a Padé approximant where each coefficient is an int64.
type Int64 struct {
	P, Q maclaurin.Int64
}

// NewInt64 returns a new zero-valued Padé approximant.
func NewInt64() *Int64 {
	return &Int64{P: *maclaurin.NewInt64(), Q: *maclaurin.NewInt64()}
}
