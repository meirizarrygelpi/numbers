// Copyright (c) 2017 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package maclaurin

import "math/big"

// A Float is a Maclaurin polynomial where each coefficient is a *big.Float.
type Float struct {
	c      map[uint64]*big.Float
	Degree uint64
}

// NewFloat returns a new zero-valued polynomial.
func NewFloat() *Float {
	return &Float{c: make(map[uint64]*big.Float)}
}

// SetCoeff sets a term in p with degree n and coefficient a.
func (p *Float) SetCoeff(n uint64, a *big.Float) {
	if n > p.Degree {
		p.Degree = n
	}
	p.c[n] = a
}

// Set sets p equal to q, and returns p.
func (p *Float) Set(q *Float) *Float {
	p = new(Float)
	for n, a := range q.c {
		p.SetCoeff(n, a)
	}
	return p
}

// Coeff returns the coefficient of the term in p with degree n. If p does
// not have a term of degree n, ok is false.
func (p *Float) Coeff(n uint64) (a *big.Float, ok bool) {
	a, ok = p.c[n]
	return
}

// Len returns the number of terms in p.
func (p *Float) Len() int {
	return len(p.c)
}

// Degrees returns a reverse-sorted slice with the non-negative degrees of p.
func (p *Float) Degrees() Degrees {
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
func (p *Float) Neg(q *Float) *Float {
	x := new(Float)
	for n, a := range q.c {
		x.SetCoeff(n, new(big.Float).Neg(a))
	}
	return p.Set(x)
}

// Add sets p equal to q+r, and returns z.
func (p *Float) Add(q, r *Float) *Float {
	x, y := new(Float), new(Float)
	x.Set(q)
	y.Set(r)
	for n, a := range x.c {
		if b, ok := y.Coeff(n); ok {
			p.SetCoeff(n, new(big.Float).Add(a, b))
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
