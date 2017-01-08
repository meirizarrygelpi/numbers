// Copyright (c) 2017 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package maclaurin

// A Float64 is a Maclaurin polynomial where each coefficient is a float64.
type Float64 map[uint64]float64

// Degrees returns a reverse-sorted slice with the non-negative degrees of p.
func (p Float64) Degrees() []uint64 {
	m := len(p)
	deg := make([]uint64, m)
	i := 0
	for n := range p {
		deg[i] = n
		i++
	}
	return reverse(sort(deg))
}
