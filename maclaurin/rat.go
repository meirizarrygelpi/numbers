// Copyright (c) 2017 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package maclaurin

import "math/big"

// A Rat is a Maclaurin polynomial where each coefficient is a *big.Rat.
type Rat struct {
	c      map[uint64]*big.Rat
	Degree uint64
}

// NewRat returns a new zero-valued polynomial.
func NewRat() *Rat {
	return &Rat{c: make(map[uint64]*big.Rat)}
}

// SetTerm sets a term in p with degree n and coefficient a.
func (p *Rat) SetTerm(n uint64, a *big.Rat) {
	if n > p.Degree {
		p.Degree = n
	}
	p.c[n] = a
}

// Coeff returns the coefficient of the term in p with degree n. If p does
// not have a term of degree n, ok is false.
func (p *Rat) Coeff(n uint64) (a *big.Rat, ok bool) {
	a, ok = p.c[n]
	return
}

// Len returns the number of terms in p.
func (p *Rat) Len() int {
	return len(p.c)
}

// Degrees returns a reverse-sorted slice with the non-negative degrees of p.
func (p Rat) Degrees() Degrees {
	n := p.Len()
	deg := make(Degrees, n)
	i := 0
	for k := range p.c {
		deg[i] = k
		i++
	}
	deg.ReverseSort()
	return deg
}
