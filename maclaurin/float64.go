// Copyright (c) 2017 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package maclaurin

// A Float64 is a Maclaurin polynomial where each coefficient is a float64.
type Float64 struct {
	c      map[uint64]float64
	Degree uint64
}

// NewFloat64 returns a new zero-valued polynomial.
func NewFloat64() *Float64 {
	return &Float64{c: make(map[uint64]float64)}
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
	p = new(Float64)
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

// Neg sets p equal to the negative of q, and returns p.
func (p *Float64) Neg(q *Float64) *Float64 {
	x := new(Float64)
	for n, a := range q.c {
		x.SetCoeff(n, -a)
	}
	return p.Set(x)
}

// Add sets p equal to q+r, and returns z.
func (p *Float64) Add(q, r *Float64) *Float64 {
	x, y := new(Float64), new(Float64)
	x.Set(q)
	y.Set(r)
	for n, a := range x.c {
		if b, ok := y.Coeff(n); ok {
			p.SetCoeff(n, a+b)
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
