// Copyright (c) 2017 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package maclaurin

import "fmt"

// An Int64 is a Maclaurin polynomial where each coefficient is an int64.
type Int64 struct {
	c      map[uint64]int64
	Degree uint64
}

// NewInt64 returns a new zero-valued polynomial.
func NewInt64() *Int64 {
	return &Int64{
		Degree: 0,
		c:      map[uint64]int64{0: 0},
	}
}

// SetCoeff sets a term in p with degree n and coefficient a.
func (p *Int64) SetCoeff(n uint64, a int64) {
	if n > p.Degree {
		p.Degree = n
	}
	p.c[n] = a
}

// Set sets p equal to q, and returns p.
func (p *Int64) Set(q *Int64) *Int64 {
	p = NewInt64()
	for n, a := range q.c {
		p.SetCoeff(n, a)
	}
	return p
}

// Coeff returns the coefficient of the term in p with degree n. If p does
// not have a term of degree n, ok is false.
func (p *Int64) Coeff(n uint64) (a int64, ok bool) {
	a, ok = p.c[n]
	return
}

// Len returns the number of terms in p.
func (p *Int64) Len() int {
	return len(p.c)
}

// Degrees returns a reverse-sorted slice with the non-negative degrees of p.
func (p *Int64) Degrees() Degrees {
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
func (p *Int64) String() string {
	l := p.Len()
	if l == 0 {
		return "0"
	}
	var s string
	degs := p.Degrees()
	s += fmt.Sprint(p.c[degs[0]])
	s += "x^"
	s += fmt.Sprint(degs[0])
	if l > 2 {
		for _, n := range degs[1:] {
			if p.c[n] < 0 {
				s += fmt.Sprint(p.c[n])
			} else {
				s += "+" + fmt.Sprint(p.c[n])
			}
			s += "x^"
			s += fmt.Sprint(n)
		}
	}
	return s
}

// Neg sets p equal to the negative of q, and returns p.
func (p *Int64) Neg(q *Int64) *Int64 {
	x := NewInt64()
	for n, a := range q.c {
		x.SetCoeff(n, -a)
	}
	return p.Set(x)
}

// Add sets p equal to q+r, and returns z.
func (p *Int64) Add(q, r *Int64) *Int64 {
	x := new(Int64).Set(q)
	y := new(Int64).Set(r)
	z := NewInt64()
	for n, a := range x.c {
		if b, ok := y.Coeff(n); ok {
			z.SetCoeff(n, a+b)
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
func (p *Int64) Sub(q, r *Int64) *Int64 {
	x := new(Int64).Set(q)
	y := new(Int64).Set(r)
	z := NewInt64()
	for n, a := range x.c {
		if b, ok := y.Coeff(n); ok {
			z.SetCoeff(n, a-b)
		} else {
			z.SetCoeff(n, a)
		}
	}
	for n, b := range y.c {
		if _, ok := x.Coeff(n); !ok {
			z.SetCoeff(n, -b)
		}
	}
	return p.Set(z)
}

// Mul sets p equal to q*r, and returns z.
func (p *Int64) Mul(q, r *Int64) *Int64 {
	x := new(Int64).Set(q)
	y := new(Int64).Set(r)
	z := NewInt64()
	var l uint64
	for n, a := range x.c {
		for m, b := range y.c {
			l = n + m
			if coeff, ok := z.Coeff(l); ok {
				z.SetCoeff(l, coeff+(a*b))
			} else {
				z.SetCoeff(l, a*b)
			}
		}
	}
	return p.Set(z)
}
