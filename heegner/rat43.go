// Copyright (c) 2017 Melvin Eloy Irizarry-Gelpí
// Licensed under the MIT License.

package heegner

import (
	"math/big"
	"math/rand"
	"reflect"
	"strings"

	"github.com/meirizarrygelpi/numbers/maclaurin"
	"github.com/meirizarrygelpi/numbers/pade"
)

// A Rat43 represents an arbitrary-precision element of ℚ(√−43).
type Rat43 struct {
	l, r big.Rat
}

// One sets z equal to 1, and then it returns z.
func (z *Rat43) One() *Rat43 {
	z.l.SetInt64(1)
	z.r.SetInt64(0)
	return z
}

// Real returns the real part of z.
func (z *Rat43) Real() *big.Rat {
	return &z.l
}

// Unreal returns the unreal part of z.
func (z *Rat43) Unreal() *big.Rat {
	return &z.r
}

// String returns the string version of a Rat value.
//
// If z corresponds to a + bC, then the string is "(a+bC)", similar to
// complex128 values.
func (z *Rat43) String() string {
	a := make([]string, 5)
	a[0] = leftBracket
	a[1] = z.l.RatString()
	if z.r.Sign() < 0 {
		a[2] = z.r.RatString()
	} else {
		a[2] = "+" + z.r.RatString()
	}
	a[3] = radical43
	a[4] = rightBracket
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Rat43) Equals(y *Rat43) bool {
	return z.l.Cmp(&y.l) == 0 && z.r.Cmp(&y.r) == 0
}

// Set sets z equal to y, and returns z.
func (z *Rat43) Set(y *Rat43) *Rat43 {
	z.l.Set(&y.l)
	z.r.Set(&y.r)
	return z
}

// SetPair sets z equal to an element made with a given pair, and then
// it returns z.
func (z *Rat43) SetPair(a, b *big.Rat) *Rat43 {
	z.l.Set(a)
	z.r.Set(b)
	return z
}

// SetReal sets the real part of z equal to a, and then it returns z.
func (z *Rat43) SetReal(a *big.Rat) *Rat43 {
	z.l.Set(a)
	return z
}

// SetUnreal sets the unreal part of z equal to b, and then it returns z.
func (z *Rat43) SetUnreal(b *big.Rat) *Rat43 {
	z.r.Set(b)
	return z
}

// NewRat43 returns a pointer to the Rat43 value a+bC.
func NewRat43(a, b *big.Rat) *Rat43 {
	z := new(Rat43)
	z.SetPair(a, b)
	return z
}

// Plus sets z equal to y+a, with a real, and returns z.
func (z *Rat43) Plus(y *Rat43, a *big.Rat) *Rat43 {
	z.l.Add(&y.l, a)
	z.r.Set(&y.r)
	return z
}

// Minus sets z equal to y-a, with a real, and returns z.
func (z *Rat43) Minus(y *Rat43, a *big.Rat) *Rat43 {
	z.l.Sub(&y.l, a)
	z.r.Set(&y.r)
	return z
}

// Scale sets z equal to y scaled by a, and returns z.
func (z *Rat43) Scale(y *Rat43, a *big.Rat) *Rat43 {
	z.l.Mul(&y.l, a)
	z.r.Mul(&y.r, a)
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Rat43) Neg(y *Rat43) *Rat43 {
	z.l.Neg(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Rat43) Conj(y *Rat43) *Rat43 {
	z.l.Set(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Add sets z equal to x+y, and returns z.
func (z *Rat43) Add(x, y *Rat43) *Rat43 {
	z.l.Add(&x.l, &y.l)
	z.r.Add(&x.r, &y.r)
	return z
}

// Sub sets z equal to x-y, and returns z.
func (z *Rat43) Sub(x, y *Rat43) *Rat43 {
	z.l.Sub(&x.l, &y.l)
	z.r.Sub(&x.r, &y.r)
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
func (z *Rat43) Mul(x, y *Rat43) *Rat43 {
	a, b, temp := new(big.Rat), new(big.Rat), new(big.Rat)
	a.Sub(
		a.Mul(&x.l, &y.l),
		temp.Mul(&y.r, temp.Mul(&x.r, h43)),
	)
	b.Add(
		temp.Mul(&x.r, &y.l),
		b.Mul(&y.r, &x.l),
	)
	z.SetPair(a, b)
	return z
}

// Quad returns the quadrance of z. If z = a+bC, then the quadrance is
// 		a² + 43b²
// This is always non-negative.
func (z *Rat43) Quad() *big.Rat {
	quad := new(big.Rat)
	return quad.Add(
		new(big.Rat).Mul(&z.l, &z.l),
		quad.Mul(&z.r, quad.Mul(&z.r, h43)),
	)
}

// Inv sets z equal to the inverse of y, and returns z. If y is zero, then Inv
// panics.
func (z *Rat43) Inv(y *Rat43) *Rat43 {
	if zero := new(Rat43); y.Equals(zero) {
		panic(zeroInverse)
	}
	a := y.Quad()
	a.Inv(a)
	return z.Scale(z.Conj(y), a)
}

// Quo sets z equal to the quotient of x and y, and returns z. If y is zero,
// then Quo panics.
func (z *Rat43) Quo(x, y *Rat43) *Rat43 {
	if zero := new(Rat43); y.Equals(zero) {
		panic(zeroDenominator)
	}
	q := y.Quad()
	a, b, temp := new(big.Rat), new(big.Rat), new(big.Rat)
	a.Add(
		a.Mul(&x.l, &y.l),
		temp.Mul(&y.r, temp.Mul(h43, &x.r)),
	)
	b.Sub(
		temp.Mul(&x.r, &y.l),
		b.Mul(&y.r, &x.l),
	)
	z.SetPair(a, b)
	return z.Scale(z, q.Inv(q))
}

// CrossRatio sets z equal to the cross-ratio of v, w, x, and y:
// 		Inv(w - x) * (v - x) * Inv(v - y) * (w - y)
// Then it returns z.
func (z *Rat43) CrossRatio(v, w, x, y *Rat43) *Rat43 {
	temp := new(Rat43)
	z.Sub(w, x)
	z.Inv(z)
	temp.Sub(v, x)
	z.Mul(z, temp)
	temp.Sub(v, y)
	temp.Inv(temp)
	z.Mul(z, temp)
	temp.Sub(w, y)
	return z.Mul(z, temp)
}

// Mobius sets z equal to the Möbius (fractional linear) transform of y:
// 		(a*y + b) * Inv(c*y + d)
// Then it returns z.
func (z *Rat43) Mobius(y, a, b, c, d *Rat43) *Rat43 {
	z.Mul(a, y)
	z.Add(z, b)
	temp := new(Rat43)
	temp.Mul(c, y)
	temp.Add(temp, d)
	temp.Inv(temp)
	return z.Mul(z, temp)
}

// Maclaurin sets z equal to the value of the Maclaurin polynomial p evaluated
// at y, and returns z. Horner's method is used.
func (z *Rat43) Maclaurin(y *Rat43, p *maclaurin.Rat) *Rat43 {
	if p.Len() == 0 {
		z = new(Rat43)
		return z
	}
	n := p.Degree
	var a *big.Rat
	if n == 0 {
		z = new(Rat43)
		a, _ = p.Coeff(n)
		z.SetReal(a)
		return z
	}
	a, _ = p.Coeff(n)
	z.Scale(y, a)
	for n > 1 {
		n--
		if a, ok := p.Coeff(n); ok {
			z.Plus(z, a)
		}
		z.Mul(z, y)
	}
	if a, ok := p.Coeff(0); ok {
		z.Plus(z, a)
	}
	return z
}

// Pade sets z equal to the value of the Padé approximant r evaluated at y,
// and returns z.
func (z *Rat43) Pade(y *Rat43, r *pade.Rat) *Rat43 {
	p, q := new(Rat43), new(Rat43)
	p.Maclaurin(y, &r.P)
	q.Maclaurin(y, &r.Q)
	return z.Quo(p, q)
}

// Generate returns a random Rat value for quick.Check testing.
func (z *Rat43) Generate(rand *rand.Rand, size int) reflect.Value {
	randomRat43 := &Rat43{
		*big.NewRat(rand.Int63(), rand.Int63()),
		*big.NewRat(rand.Int63(), rand.Int63()),
	}
	return reflect.ValueOf(randomRat43)
}
