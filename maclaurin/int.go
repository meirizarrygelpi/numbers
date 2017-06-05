// Copyright (c) 2017 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package maclaurin

import "math/big"

// A Int is a Maclaurin polynomial where each coefficient is a *big.Int.
type Int struct {
	c      map[uint64]*big.Int
	Degree uint64
}

// NewInt returns a new zero-valued polynomial.
func NewInt() *Int {
	return &Int{c: make(map[uint64]*big.Int)}
}

// SetCoeff sets a term in p with degree n and coefficient a.
func (p *Int) SetCoeff(n uint64, a *big.Int) {
	if n > p.Degree {
		p.Degree = n
	}
	p.c[n] = a
}

// Set sets p equal to q, and returns p.
func (p *Int) Set(q *Int) *Int {
	p = new(Int)
	for n, a := range q.c {
		p.SetCoeff(n, a)
	}
	return p
}

// Coeff returns the coefficient of the term in p with degree n. If p does
// not have a term of degree n, ok is false.
func (p *Int) Coeff(n uint64) (a *big.Int, ok bool) {
	a, ok = p.c[n]
	return
}

// Len returns the number of terms in p.
func (p *Int) Len() int {
	return len(p.c)
}

// Degrees returns a reverse-sorted slice with the non-negative degrees of p.
func (p *Int) Degrees() Degrees {
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

// Add sets p equal to q+r, and returns z.
func (p *Int) Add(q, r *Int) *Int {
	x, y := new(Int), new(Int)
	x.Set(q)
	y.Set(r)
	for n, a := range x.c {
		if b, ok := y.Coeff(n); ok {
			p.SetCoeff(n, new(big.Int).Add(a, b))
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
