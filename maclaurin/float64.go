// Copyright (c) 2017 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package maclaurin

import (
	"fmt"
	"math"
)

// A Float64 is a Maclaurin polynomial where each coefficient is a float64.
type Float64 struct {
	c      map[uint64]float64
	Degree uint64
}

// NewFloat64 returns a new zero-valued polynomial.
func NewFloat64() *Float64 {
	return &Float64{
		Degree: 0,
		c:      map[uint64]float64{0: 0.0},
	}
}

// Equals returns true if p is equal to q.
func (p *Float64) Equals(q *Float64) bool {
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
		if a != b {
			return false
		}
	}
	return true
}

// SetCoeff sets a term in p with degree n and coefficient a.
func (p *Float64) SetCoeff(n uint64, a float64) {
	if n > p.Degree {
		p.Degree = n
	}
	p.c[n] = a
}

// Set sets p equal to q, and returns p.
func (p *Float64) Set(q *Float64) *Float64 {
	p = NewFloat64()
	for n, a := range q.c {
		p.SetCoeff(n, a)
	}
	return p
}

// Coeff returns the coefficient of the term in p with degree n. If p does
// not have a term of degree n, ok is false.
func (p *Float64) Coeff(n uint64) (a float64, ok bool) {
	a, ok = p.c[n]
	return
}

// Len returns the number of terms in p.
func (p *Float64) Len() int {
	return len(p.c)
}

// Degrees returns a reverse-sorted slice with the non-negative degrees of p.
func (p *Float64) Degrees() Degrees {
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

func sprintFloat64(a float64) string {
	if math.Signbit(a) {
		return fmt.Sprintf("%g", a)
	}
	if math.IsInf(a, +1) {
		return "+Inf"
	}
	return fmt.Sprintf("+%g", a)
}

// String returns the string version of a polynomial.
func (p *Float64) String() string {
	l := p.Len()
	if l == 0 {
		return "0.0"
	}
	var s string
	degs := p.Degrees()
	s += fmt.Sprintf("%g", p.c[degs[0]])
	s += "x^"
	s += fmt.Sprint(degs[0])
	if l > 2 {
		for _, n := range degs[1:] {
			s += sprintFloat64(p.c[n])
			s += "x^"
			s += fmt.Sprint(n)
		}
	}
	return s
}

// Neg sets p equal to the negative of q, and returns p.
func (p *Float64) Neg(q *Float64) *Float64 {
	x := NewFloat64()
	for n, a := range q.c {
		x.SetCoeff(n, -a)
	}
	return p.Set(x)
}

// Add sets p equal to q+r, and returns z.
func (p *Float64) Add(q, r *Float64) *Float64 {
	x := new(Float64).Set(q)
	y := new(Float64).Set(r)
	z := NewFloat64()
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
func (p *Float64) Sub(q, r *Float64) *Float64 {
	x := new(Float64).Set(q)
	y := new(Float64).Set(r)
	z := NewFloat64()
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
func (p *Float64) Mul(q, r *Float64) *Float64 {
	x := new(Float64).Set(q)
	y := new(Float64).Set(r)
	z := NewFloat64()
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
