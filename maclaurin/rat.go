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

// SetCoeff sets a term in p with degree n and coefficient a.
func (p *Rat) SetCoeff(n uint64, a *big.Rat) {
	if n > p.Degree {
		p.Degree = n
	}
	p.c[n] = a
}

// Set sets p equal to q, and returns p.
func (p *Rat) Set(q *Rat) *Rat {
	p = new(Rat)
	for n, a := range q.c {
		p.SetCoeff(n, a)
	}
	return p
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
func (p *Rat) Degrees() Degrees {
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

// Neg sets p equal to the negative of q, and returns p.
func (p *Rat) Neg(q *Rat) *Rat {
	x := new(Rat)
	for n, a := range q.c {
		x.SetCoeff(n, new(big.Rat).Neg(a))
	}
	return p.Set(x)
}

// Add sets p equal to q+r, and returns z.
func (p *Rat) Add(q, r *Rat) *Rat {
	x, y := new(Rat), new(Rat)
	x.Set(q)
	y.Set(r)
	for n, a := range x.c {
		if b, ok := y.Coeff(n); ok {
			p.SetCoeff(n, new(big.Rat).Add(a, b))
		} else {
			p.SetCoeff(n, a)
		}
	}
	for n, b := range y.c {
		if _, ok := y.Coeff(n); !ok {
			p.SetCoeff(n, b)
		}
	}
	return p
}
