// Copyright (c) 2017 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package maclaurin

// An Int64 is a Maclaurin polynomial where each coefficient is an int64.
type Int64 struct {
	c      map[uint64]int64
	Degree uint64
}

// NewInt64 returns a new zero-valued polynomial.
func NewInt64() *Int64 {
	return &Int64{c: make(map[uint64]int64)}
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
	p = new(Int64)
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

// Neg sets p equal to the negative of q, and returns p.
func (p *Int64) Neg(q *Int64) *Int64 {
	x := new(Int64)
	for n, a := range q.c {
		x.SetCoeff(n, -a)
	}
	return p.Set(x)
}

// Add sets p equal to q+r, and returns z.
func (p *Int64) Add(q, r *Int64) *Int64 {
	x, y := new(Int64), new(Int64)
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
