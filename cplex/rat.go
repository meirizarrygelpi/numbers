// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licensed under the MIT License.

package cplex

import (
	"math/big"
	"math/rand"
	"reflect"
	"strings"

	"github.com/meirizarrygelpi/numbers/maclaurin"
	"github.com/meirizarrygelpi/numbers/pade"
)

// A Rat represents an arbitrary-precision complex rational.
type Rat struct {
	l, r big.Rat
}

// One sets z equal to 1, and then it returns z.
func (z *Rat) One() *Rat {
	z.l.SetInt64(1)
	z.r.SetInt64(0)
	return z
}

// Real returns the real part of z.
func (z *Rat) Real() *big.Rat {
	return &z.l
}

// Unreal returns the unreal part of z.
func (z *Rat) Unreal() *big.Rat {
	return &z.r
}

// String returns the string version of a Rat value.
//
// If z corresponds to a + bi, then the string is "(a+bi)", similar to
// complex128 values.
func (z *Rat) String() string {
	a := make([]string, 5)
	a[0] = leftBracket
	a[1] = z.l.RatString()
	if z.r.Sign() < 0 {
		a[2] = z.r.RatString()
	} else {
		a[2] = "+" + z.r.RatString()
	}
	a[3] = unitName
	a[4] = rightBracket
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Rat) Equals(y *Rat) bool {
	return z.l.Cmp(&y.l) == 0 && z.r.Cmp(&y.r) == 0
}

// Set sets z equal to y, and returns z.
func (z *Rat) Set(y *Rat) *Rat {
	z.l.Set(&y.l)
	z.r.Set(&y.r)
	return z
}

// SetPair sets z equal to a complex number made with a given pair, and then
// it returns z.
func (z *Rat) SetPair(a, b *big.Rat) *Rat {
	z.l.Set(a)
	z.r.Set(b)
	return z
}

// SetReal sets the real part of z equal to a, and then it returns z.
func (z *Rat) SetReal(a *big.Rat) *Rat {
	z.l.Set(a)
	return z
}

// SetUnreal sets the unreal part of z equal to b, and then it returns z.
func (z *Rat) SetUnreal(b *big.Rat) *Rat {
	z.r.Set(b)
	return z
}

// NewRat returns a pointer to the Rat value a+bi.
func NewRat(a, b *big.Rat) *Rat {
	z := new(Rat)
	z.SetPair(a, b)
	return z
}

// Plus sets z equal to y+a, with a real, and returns z.
func (z *Rat) Plus(y *Rat, a *big.Rat) *Rat {
	z.l.Add(&y.l, a)
	z.r.Set(&y.r)
	return z
}

// Minus sets z equal to y-a, with a real, and returns z.
func (z *Rat) Minus(y *Rat, a *big.Rat) *Rat {
	z.l.Sub(&y.l, a)
	z.r.Set(&y.r)
	return z
}

// Scale sets z equal to y scaled by a, and returns z.
func (z *Rat) Scale(y *Rat, a *big.Rat) *Rat {
	z.l.Mul(&y.l, a)
	z.r.Mul(&y.r, a)
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Rat) Neg(y *Rat) *Rat {
	z.l.Neg(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Rat) Conj(y *Rat) *Rat {
	z.l.Set(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Add sets z equal to x+y, and returns z.
func (z *Rat) Add(x, y *Rat) *Rat {
	z.l.Add(&x.l, &y.l)
	z.r.Add(&x.r, &y.r)
	return z
}

// Sub sets z equal to x-y, and returns z.
func (z *Rat) Sub(x, y *Rat) *Rat {
	z.l.Sub(&x.l, &y.l)
	z.r.Sub(&x.r, &y.r)
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
func (z *Rat) Mul(x, y *Rat) *Rat {
	a, b, temp := new(big.Rat), new(big.Rat), new(big.Rat)
	a.Sub(
		a.Mul(&x.l, &y.l),
		temp.Mul(&y.r, &x.r),
	)
	b.Add(
		temp.Mul(&x.r, &y.l),
		b.Mul(&y.r, &x.l),
	)
	z.SetPair(a, b)
	return z
}

// Dot returns the dot product of y and z. If z = a+bi and y = c+di, then the
// dot product is
// 		ac + bd
// This can be positive, negative, or zero. The dot product is equivalent to
// 		½(Mul(Conj(z), y) + Mu(Conj(y), z))
// In this form it is clear that Dot is symmetric.
func (z *Rat) Dot(y *Rat) *big.Rat {
	dot := new(big.Rat)
	return dot.Add(
		dot.Mul(&z.l, &y.l),
		new(big.Rat).Mul(&z.r, &y.r),
	)
}

// Quad returns the quadrance of z. If z = a+bi, then the quadrance is
// 		a² + b²
// This is always non-negative.
func (z *Rat) Quad() *big.Rat {
	return z.Dot(z)
}

// Cross returns the cross product of y and z. If z = a+bi and y = c+di, then
// the cross product is
// 		ad - bc
// This can be positive, negative, or zero. The cross product is equivalent to
// the unreal part of
// 		½(Mul(Conj(z), y) - Mu(Conj(y), z))
// In this form it is clear that Cross is anti-symmetric.
func (z *Rat) Cross(y *Rat) *big.Rat {
	cross := new(big.Rat)
	return cross.Sub(
		cross.Mul(&z.l, &y.r),
		new(big.Rat).Mul(&z.r, &y.l),
	)
}

// Lozenge sets z equal to the lozenge product of v, w, x, and y:
// 		Mul(v, Conj(w)) - Mul(x, Conj(y))
// Then it returns z.
func (z *Rat) Lozenge(v, w, x, y *Rat) *Rat {
	a, b := new(Rat), new(Rat)
	a.Mul(v, a.Conj(w))
	b.Mul(x, b.Conj(y))
	return z.Sub(a, b)
}

// Inv sets z equal to the inverse of y, and returns z. If y is zero, then Inv
// panics.
func (z *Rat) Inv(y *Rat) *Rat {
	if zero := new(Rat); y.Equals(zero) {
		panic(zeroInverse)
	}
	a := y.Quad()
	a.Inv(a)
	return z.Scale(z.Conj(y), a)
}

// Quo sets z equal to the quotient of x and y, and returns z. If y is zero,
// then Quo panics.
func (z *Rat) Quo(x, y *Rat) *Rat {
	if zero := new(Rat); y.Equals(zero) {
		panic(zeroDenominator)
	}
	q := y.Quad()
	a, b, temp := new(big.Rat), new(big.Rat), new(big.Rat)
	a.Add(
		a.Mul(&x.l, &y.l),
		temp.Mul(&y.r, &x.r),
	)
	b.Sub(
		temp.Mul(&x.r, &y.l),
		b.Mul(&y.r, &x.l),
	)
	z.SetPair(a, b)
	return z.Scale(z, q.Inv(q))
}

// Gauss sets z equal to the Gaussian integer a+bi, and returns z.
func (z *Rat) Gauss(a, b *big.Int) *Rat {
	z.l.SetInt(a)
	z.r.SetInt(b)
	return z
}

// CrossRatio sets z equal to the cross-ratio of v, w, x, and y:
// 		Inv(w - x) * (v - x) * Inv(v - y) * (w - y)
// Then it returns z.
func (z *Rat) CrossRatio(v, w, x, y *Rat) *Rat {
	temp := new(Rat)
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
func (z *Rat) Mobius(y, a, b, c, d *Rat) *Rat {
	z.Mul(a, y)
	z.Add(z, b)
	temp := new(Rat)
	temp.Mul(c, y)
	temp.Add(temp, d)
	temp.Inv(temp)
	return z.Mul(z, temp)
}

// Maclaurin sets z equal to the value of the Maclaurin polynomial p evaluated
// at y, and returns z. Horner's method is used.
func (z *Rat) Maclaurin(y *Rat, p *maclaurin.Rat) *Rat {
	if p.Len() == 0 {
		z = new(Rat)
		return z
	}
	n := p.Degree
	var a *big.Rat
	if n == 0 {
		z = new(Rat)
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
func (z *Rat) Pade(y *Rat, r *pade.Rat) *Rat {
	p, q := new(Rat), new(Rat)
	p.Maclaurin(y, &r.P)
	q.Maclaurin(y, &r.Q)
	return z.Quo(p, q)
}

// Generate returns a random Rat value for quick.Check testing.
func (z *Rat) Generate(rand *rand.Rand, size int) reflect.Value {
	randomRat := &Rat{
		*big.NewRat(rand.Int63(), rand.Int63()),
		*big.NewRat(rand.Int63(), rand.Int63()),
	}
	return reflect.ValueOf(randomRat)
}
