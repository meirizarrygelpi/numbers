// Copyright (c) 2017 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package maclaurin

// An Int64 is a Maclaurin polynomial where each coefficient is an int64.
type Int64 struct {
	c      map[uint64]int64
	Degree uint64
}

// SetTerm sets a term in p with degree n and coefficient a.
func (p *Int64) SetTerm(n uint64, a int64) {
	if n > p.Degree {
		p.Degree = n
	}
	p.c[n] = a
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
func (p Int64) Degrees() Degrees {
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
