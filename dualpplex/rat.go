// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package dualpplex

import (
	"math/big"
	"math/rand"
	"reflect"
	"strings"

	"github.com/meirizarrygelpi/numbers/pplex"
	"github.com/meirizarrygelpi/numbers/vec3"
)

// A Rat is a dual-perplex number with big.Rat components.
type Rat struct {
	l, r pplex.Rat
}

// One sets z equal to 1, and then returns z.
func (z *Rat) One() *Rat {
	z.l.One()
	z.r.Set(new(pplex.Rat))
	return z
}

// Real returns the real part of z.
func (z *Rat) Real() *big.Rat {
	return z.l.Real()
}

// Unreal returns the unreal part of z.
func (z *Rat) Unreal() *vec3.Rat {
	v := new(vec3.Rat)
	v[0] = z.l.Unreal()
	v[1] = z.r.Real()
	v[2] = z.r.Unreal()
	return v
}

// String returns the string version of a Rat value.
//
// If z corresponds to a+bs+cΓ+dsΓ, then the string is "⦗a+bs+cΓ+dsΓ⦘", similar
// to complex128 values.
func (z *Rat) String() string {
	v := z.Unreal()
	a := make([]string, 9)
	a[0] = leftBracket
	a[1] = z.l.Real().RatString()
	i := 2
	for j, u := range unitNames {
		if v[j].Sign() < 0 {
			a[i] = v[j].RatString()
		} else {
			a[i] = "+" + v[j].RatString()
		}
		a[i+1] = u
		i += 2
	}
	a[8] = rightBracket
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Rat) Equals(y *Rat) bool {
	return z.l.Equals(&y.l) && z.r.Equals(&y.r)
}

// Set sets z equal to y, and returns z.
func (z *Rat) Set(y *Rat) *Rat {
	z.l.Set(&y.l)
	z.r.Set(&y.r)
	return z
}

// SetPair sets z equal to a dual-perplex number made with a given pair, and
// then it returns z.
func (z *Rat) SetPair(a, b *pplex.Rat) *Rat {
	z.l.Set(a)
	z.r.Set(b)
	return z
}

// NewRat returns a pointer to the Rat value a+bs+cΓ+dsΓ.
func NewRat(a, b, c, d *big.Rat) *Rat {
	z := new(Rat)
	z.l.SetPair(a, b)
	z.r.SetPair(c, d)
	return z
}

// Scale sets z equal to y scaled by a, and returns z.
func (z *Rat) Scale(y *Rat, a *big.Rat) *Rat {
	z.l.Scale(&y.l, a)
	z.r.Scale(&y.r, a)
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
	z.l.Conj(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Bar sets z equal to the s-conjugate of y, and returns z.
func (z *Rat) Bar(y *Rat) *Rat {
	z.l.Conj(&y.l)
	z.r.Conj(&y.r)
	return z
}

// Tilde sets z equal to the Γ-conjugate of y, and returns z.
func (z *Rat) Tilde(y *Rat) *Rat {
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
	a, b, temp := new(pplex.Rat), new(pplex.Rat), new(pplex.Rat)
	a.Mul(&x.l, &y.l)
	b.Add(
		b.Mul(&x.l, &y.r),
		temp.Mul(&x.r, &y.l),
	)
	z.SetPair(a, b)
	return z
}

// Quad returns the quadrance of z. If z = a+bs+cΓ+dsΓ, then the quadrance is
//     a² + b² + 2abs
// Note that this is a perplex number.
func (z *Rat) Quad() *pplex.Rat {
	q := new(pplex.Rat)
	return q.Mul(&z.l, &z.l)
}

// Norm returns the norm of z. If z = a+bs+cΓ+dsΓ, then the norm is
// 		(a² - b²)²
// This is always non-negative.
func (z *Rat) Norm() *big.Rat {
	return z.Quad().Quad()
}

// IsZeroDivisor returns true if z is a zero divisor.
func (z *Rat) IsZeroDivisor() bool {
	return z.Quad().IsZeroDivisor()
}

// Inv sets z equal to the inverse of y, and returns z. If y is zero, then Inv
// panics.
func (z *Rat) Inv(y *Rat) *Rat {
	if y.IsZeroDivisor() {
		panic(zeroDivisorInverse)
	}
	a := y.Quad()
	a.Inv(a)
	z.l.Mul(&y.l, a)
	z.r.Mul(&y.r, a)
	return z.Tilde(z)
}

// Quo sets z equal to the quotient of x and y, and returns z. If y is zero,
// then Quo panics.
func (z *Rat) Quo(x, y *Rat) *Rat {
	if y.IsZeroDivisor() {
		panic(zeroDivisorDenominator)
	}
	return z.Mul(x, z.Inv(y))
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

// Möbius sets z equal to the Möbius (fractional linear) transform of y:
// 		(a*y + b) * Inv(c*y + d)
// Then it returns z.
func (z *Rat) Möbius(y, a, b, c, d *Rat) *Rat {
	z.Mul(a, y)
	z.Add(z, b)
	temp := new(Rat)
	temp.Mul(c, y)
	temp.Add(temp, d)
	temp.Inv(temp)
	return z.Mul(z, temp)
}

// Generate returns a random Rat value for quick.Check testing.
func (z *Rat) Generate(rand *rand.Rand, size int) reflect.Value {
	randomRat := &Rat{
		*pplex.NewRat(
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
		),
		*pplex.NewRat(
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
		),
	}
	return reflect.ValueOf(randomRat)
}
