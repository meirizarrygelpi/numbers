// Copyright (c) 2017 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package heegner

import (
	"math/big"
	"math/rand"
	"reflect"
	"strings"

	"github.com/meirizarrygelpi/numbers/maclaurin"
	"github.com/meirizarrygelpi/numbers/pade"
)

// A Rat19 represents an arbitrary-precision element of ℚ(√−19).
type Rat19 struct {
	l, r big.Rat
}

// One sets z equal to 1, and then it returns z.
func (z *Rat19) One() *Rat19 {
	z.l.SetInt64(1)
	z.r.SetInt64(0)
	return z
}

// Real returns the real part of z.
func (z *Rat19) Real() *big.Rat {
	return &z.l
}

// Unreal returns the unreal part of z.
func (z *Rat19) Unreal() *big.Rat {
	return &z.r
}

// String returns the string version of a Rat value.
//
// If z corresponds to a + bD, then the string is "(a+bD)", similar to
// complex128 values.
func (z *Rat19) String() string {
	a := make([]string, 5)
	a[0] = leftBracket
	a[1] = z.l.RatString()
	if z.r.Sign() < 0 {
		a[2] = z.r.RatString()
	} else {
		a[2] = "+" + z.r.RatString()
	}
	a[3] = radical19
	a[4] = rightBracket
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Rat19) Equals(y *Rat19) bool {
	return z.l.Cmp(&y.l) == 0 && z.r.Cmp(&y.r) == 0
}

// Set sets z equal to y, and returns z.
func (z *Rat19) Set(y *Rat19) *Rat19 {
	z.l.Set(&y.l)
	z.r.Set(&y.r)
	return z
}

// SetPair sets z equal to an element made with a given pair, and then
// it returns z.
func (z *Rat19) SetPair(a, b *big.Rat) *Rat19 {
	z.l.Set(a)
	z.r.Set(b)
	return z
}

// SetReal sets the real part of z equal to a, and then it returns z.
func (z *Rat19) SetReal(a *big.Rat) *Rat19 {
	z.l.Set(a)
	return z
}

// SetUnreal sets the unreal part of z equal to b, and then it returns z.
func (z *Rat19) SetUnreal(b *big.Rat) *Rat19 {
	z.r.Set(b)
	return z
}

// NewRat19 returns a pointer to the Rat19 value a+bD.
func NewRat19(a, b *big.Rat) *Rat19 {
	z := new(Rat19)
	z.SetPair(a, b)
	return z
}

// Plus sets z equal to y+a, with a real, and returns z.
func (z *Rat19) Plus(y *Rat19, a *big.Rat) *Rat19 {
	z.l.Add(&y.l, a)
	z.r.Set(&y.r)
	return z
}

// Minus sets z equal to y-a, with a real, and returns z.
func (z *Rat19) Minus(y *Rat19, a *big.Rat) *Rat19 {
	z.l.Sub(&y.l, a)
	z.r.Set(&y.r)
	return z
}

// Scale sets z equal to y scaled by a, and returns z.
func (z *Rat19) Scale(y *Rat19, a *big.Rat) *Rat19 {
	z.l.Mul(&y.l, a)
	z.r.Mul(&y.r, a)
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Rat19) Neg(y *Rat19) *Rat19 {
	z.l.Neg(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Rat19) Conj(y *Rat19) *Rat19 {
	z.l.Set(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Add sets z equal to x+y, and returns z.
func (z *Rat19) Add(x, y *Rat19) *Rat19 {
	z.l.Add(&x.l, &y.l)
	z.r.Add(&x.r, &y.r)
	return z
}

// Sub sets z equal to x-y, and returns z.
func (z *Rat19) Sub(x, y *Rat19) *Rat19 {
	z.l.Sub(&x.l, &y.l)
	z.r.Sub(&x.r, &y.r)
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
func (z *Rat19) Mul(x, y *Rat19) *Rat19 {
	a, b, temp := new(big.Rat), new(big.Rat), new(big.Rat)
	a.Sub(
		a.Mul(&x.l, &y.l),
		temp.Mul(&y.r, temp.Mul(&x.r, h19)),
	)
	b.Add(
		temp.Mul(&x.r, &y.l),
		b.Mul(&y.r, &x.l),
	)
	z.SetPair(a, b)
	return z
}

// Quad returns the quadrance of z. If z = a+bD, then the quadrance is
// 		a² + 19b²
// This is always non-negative.
func (z *Rat19) Quad() *big.Rat {
	quad := new(big.Rat)
	return quad.Add(
		new(big.Rat).Mul(&z.l, &z.l),
		quad.Mul(&z.r, quad.Mul(&z.r, h19)),
	)
}

// Inv sets z equal to the inverse of y, and returns z. If y is zero, then Inv
// panics.
func (z *Rat19) Inv(y *Rat19) *Rat19 {
	if zero := new(Rat19); y.Equals(zero) {
		panic(zeroInverse)
	}
	a := y.Quad()
	a.Inv(a)
	return z.Scale(z.Conj(y), a)
}

// Quo sets z equal to the quotient of x and y, and returns z. If y is zero,
// then Quo panics.
func (z *Rat19) Quo(x, y *Rat19) *Rat19 {
	if zero := new(Rat19); y.Equals(zero) {
		panic(zeroDenominator)
	}
	q := y.Quad()
	a, b, temp := new(big.Rat), new(big.Rat), new(big.Rat)
	a.Add(
		a.Mul(&x.l, &y.l),
		temp.Mul(&y.r, temp.Mul(h19, &x.r)),
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
func (z *Rat19) CrossRatio(v, w, x, y *Rat19) *Rat19 {
	temp := new(Rat19)
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

// Möbius sets z equal to the Möbius (fractional linear) transform of y:
// 		(a*y + b) * Inv(c*y + d)
// Then it returns z.
func (z *Rat19) Möbius(y, a, b, c, d *Rat19) *Rat19 {
	z.Mul(a, y)
	z.Add(z, b)
	temp := new(Rat19)
	temp.Mul(c, y)
	temp.Add(temp, d)
	temp.Inv(temp)
	return z.Mul(z, temp)
}

// Maclaurin sets z equal to the value of the Maclaurin polynomial p evaluated
// at y, and returns z. Horner's method is used.
func (z *Rat19) Maclaurin(y *Rat19, p *maclaurin.Rat) *Rat19 {
	if p.Len() == 0 {
		z = new(Rat19)
		return z
	}
	n := p.Degree
	var a *big.Rat
	if n == 0 {
		z = new(Rat19)
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

// Padé sets z equal to the value of the Padé approximant r evaluated at y,
// and returns z.
func (z *Rat19) Padé(y *Rat19, r *pade.Rat) *Rat19 {
	p, q := new(Rat19), new(Rat19)
	p.Maclaurin(y, &r.P)
	q.Maclaurin(y, &r.Q)
	return z.Quo(p, q)
}

// Generate returns a random Rat value for quick.Check testing.
func (z *Rat19) Generate(rand *rand.Rand, size int) reflect.Value {
	randomRat19 := &Rat19{
		*big.NewRat(rand.Int63(), rand.Int63()),
		*big.NewRat(rand.Int63(), rand.Int63()),
	}
	return reflect.ValueOf(randomRat19)
}
