// Copyright (c) 2017 Melvin Eloy Irizarry-Gelpí
// Licensed under the MIT License.

package pade

import "github.com/meirizarrygelpi/numbers/maclaurin"

// A Float64 is a Padé approximant where each coefficient is a float64.
type Float64 struct {
	P, Q maclaurin.Float64
}

// NewFloat64 returns a new zero-valued Padé approximant.
func NewFloat64() *Float64 {
	return &Float64{P: *maclaurin.NewFloat64(), Q: *maclaurin.NewFloat64()}
}
