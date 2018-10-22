// Copyright (c) 2017 Melvin Eloy Irizarry-Gelpí
// Licensed under the MIT License.

package maclaurin

import (
	"fmt"
	"math/big"
)

// A Float is a Maclaurin polynomial where each coefficient is a *big.Float.
type Float struct {
	c      map[uint64]*big.Float
	Degree uint64
}

// NewFloat returns a new zero-valued polynomial.
func NewFloat() *Float {
	return &Float{
		Degree: 0,
		c:      map[uint64]*big.Float{0: big.NewFloat(0.0)},
	}
}

// Equals returns true if p is equal to q. This is a very naive method, as it
// does not account for zero coefficients.
func (p *Float) Equals(q *Float) bool {
	if p.Degree != q.Degree {
		return false
	}
	pdegs := p.Degrees()
	qdegs := q.Degrees()
	if pdegs.Len() != qdegs.Len() {
		return false
	}
	for i, deg := range pdegs {
		if deg != qdegs[i] {
			return false
		}
		a, _ := p.Coeff(deg)
		b, _ := q.Coeff(deg)
		if a.Cmp(b) != 0 {
			return false
		}
	}
	return true
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
	p = NewFloat()
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

func sprintFloat(a *big.Float) string {
	if a.Signbit() {
		return a.String()
	}
	if a.IsInf() {
		return "+Inf"
	}
	return "+" + a.String()
}

// String returns the string version of a polynomial.
func (p *Float) String() string {
	l := p.Len()
	if l == 0 {
		return "0.0"
	}
	var s string
	degs := p.Degrees()
	s += p.c[degs[0]].String()
	s += "x^"
	s += fmt.Sprint(degs[0])
	if l > 2 {
		for _, n := range degs[1:] {
			s += sprintFloat(p.c[n])
			s += "x^"
			s += fmt.Sprint(n)
		}
	}
	return s
}

// Neg sets p equal to the negative of q, and returns p.
func (p *Float) Neg(q *Float) *Float {
	x := NewFloat()
	for n, a := range q.c {
		x.SetCoeff(n, new(big.Float).Neg(a))
	}
	return p.Set(x)
}

// Add sets p equal to q+r, and returns z.
func (p *Float) Add(q, r *Float) *Float {
	x := new(Float).Set(q)
	y := new(Float).Set(r)
	z := NewFloat()
	for n, a := range x.c {
		if b, ok := y.Coeff(n); ok {
			z.SetCoeff(n, new(big.Float).Add(a, b))
		} else {
			z.SetCoeff(n, a)
		}
	}
	for n, b := range y.c {
		if _, ok := x.Coeff(n); !ok {
			z.SetCoeff(n, b)
		}
	}
	return p.Set(z)
}

// Sub sets p equal to q-r, and returns z.
func (p *Float) Sub(q, r *Float) *Float {
	x := new(Float).Set(q)
	y := new(Float).Set(r)
	z := NewFloat()
	for n, a := range x.c {
		if b, ok := y.Coeff(n); ok {
			z.SetCoeff(n, new(big.Float).Sub(a, b))
		} else {
			z.SetCoeff(n, a)
		}
	}
	for n, b := range y.c {
		if _, ok := x.Coeff(n); !ok {
			z.SetCoeff(n, new(big.Float).Neg(b))
		}
	}
	return p.Set(z)
}

// Mul sets p equal to q*r, and returns z.
func (p *Float) Mul(q, r *Float) *Float {
	x := new(Float).Set(q)
	y := new(Float).Set(r)
	z := NewFloat()
	var l uint64
	for n, a := range x.c {
		for m, b := range y.c {
			l = n + m
			if coeff, ok := z.Coeff(l); ok {
				z.SetCoeff(l, new(big.Float).Add(coeff, new(big.Float).Mul(a, b)))
			} else {
				z.SetCoeff(l, new(big.Float).Mul(a, b))
			}
		}
	}
	return p.Set(z)
}

// QuoRem sets p equal to the quotient of q/r, and s equal to the reminder of
// q/r. Then it returns p and s. WIP
func (p *Float) QuoRem(q, r, s *Float) (*Float, *Float) {
	if zero := NewFloat(); r.Equals(zero) {
		panic(zeroDenominator)
	}
	return p, s
}

// Quo sets p equal to the quotient of q/r, and returns z. WIP
func (p *Float) Quo(q, r *Float) *Float {
	if zero := NewFloat(); r.Equals(zero) {
		panic(zeroDenominator)
	}
	return p
}

// Rem sets p equal to the reminder of q/r, and returns z. WIP
func (p *Float) Rem(q, r *Float) *Float {
	if zero := NewFloat(); r.Equals(zero) {
		panic(zeroDenominator)
	}
	return p
}
