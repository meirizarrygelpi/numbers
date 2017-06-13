// Copyright (c) 2017 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package maclaurin

import (
	"fmt"
	"math/big"
)

// A Rat is a Maclaurin polynomial where each coefficient is a *big.Rat.
type Rat struct {
	c      map[uint64]*big.Rat
	Degree uint64
}

// NewRat returns a new zero-valued polynomial.
func NewRat() *Rat {
	return &Rat{
		Degree: 0,
		c:      map[uint64]*big.Rat{0: big.NewRat(0, 1)},
	}
}

// Equals returns true if p is equal to q.
func (p *Rat) Equals(q *Rat) bool {
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
func (p *Rat) SetCoeff(n uint64, a *big.Rat) {
	if n > p.Degree {
		p.Degree = n
	}
	p.c[n] = a
}

// Set sets p equal to q, and returns p.
func (p *Rat) Set(q *Rat) *Rat {
	p = NewRat()
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

// String returns the string version of a polynomial.
func (p *Rat) String() string {
	l := p.Len()
	if l == 0 {
		return "0"
	}
	var s string
	degs := p.Degrees()
	s += p.c[degs[0]].RatString()
	s += "x^"
	s += fmt.Sprint(degs[0])
	if l > 2 {
		for _, n := range degs[1:] {
			if p.c[n].Sign() < 0 {
				s += p.c[n].RatString()
			} else {
				s += "+" + p.c[n].RatString()
			}
			s += "x^"
			s += fmt.Sprint(n)
		}
	}
	return s
}

// Neg sets p equal to the negative of q, and returns p.
func (p *Rat) Neg(q *Rat) *Rat {
	x := NewRat()
	for n, a := range q.c {
		x.SetCoeff(n, new(big.Rat).Neg(a))
	}
	return p.Set(x)
}

// Add sets p equal to q+r, and returns z.
func (p *Rat) Add(q, r *Rat) *Rat {
	x := new(Rat).Set(q)
	y := new(Rat).Set(r)
	z := NewRat()
	for n, a := range x.c {
		if b, ok := y.Coeff(n); ok {
			z.SetCoeff(n, new(big.Rat).Add(a, b))
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
func (p *Rat) Sub(q, r *Rat) *Rat {
	x := new(Rat).Set(q)
	y := new(Rat).Set(r)
	z := NewRat()
	for n, a := range x.c {
		if b, ok := y.Coeff(n); ok {
			z.SetCoeff(n, new(big.Rat).Sub(a, b))
		} else {
			z.SetCoeff(n, a)
		}
	}
	for n, b := range y.c {
		if _, ok := x.Coeff(n); !ok {
			z.SetCoeff(n, new(big.Rat).Neg(b))
		}
	}
	return p.Set(z)
}

// Mul sets p equal to q*r, and returns z.
func (p *Rat) Mul(q, r *Rat) *Rat {
	x := new(Rat).Set(q)
	y := new(Rat).Set(r)
	z := NewRat()
	var l uint64
	for n, a := range x.c {
		for m, b := range y.c {
			l = n + m
			if coeff, ok := z.Coeff(l); ok {
				z.SetCoeff(l, new(big.Rat).Add(coeff, new(big.Rat).Mul(a, b)))
			} else {
				z.SetCoeff(l, new(big.Rat).Mul(a, b))
			}
		}
	}
	return p.Set(z)
}
