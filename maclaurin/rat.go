// Copyright (c) 2017 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package maclaurin

import "math/big"

// A Rat is a Maclaurin polynomial where each coefficient is a *big.Rat.
type Rat map[uint64]*big.Rat

// Degrees returns a reverse-sorted slice with the non-negative degrees of p.
func (p Rat) Degrees() []uint64 {
	m := len(p)
	deg := make([]uint64, m)
	i := 0
	for n := range p {
		deg[i] = n
		i++
	}
	return reverse(sort(deg))
}
