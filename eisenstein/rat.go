// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package eisenstein

import (
	"fmt"
	"math/big"
	"math/rand"
	"reflect"
	"strings"
)

// A Rat represents an arbitrary-precision Eisenstein rational.
type Rat struct {
	l, r big.Rat
}

// One sets z equal to 1, and then it returns z.
func (z *Rat) One() *Rat {
	z.l.SetInt64(1)
	z.r.SetInt64(0)
	return z
}

// Omega sets z equal to the Eisenstein unit ω, and then it returns z.
func (z *Rat) Omega() *Rat {
	z.l.SetInt64(0)
	z.r.SetInt64(1)
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
func (z *Rat) String() string {
	a := make([]string, 5)
	a[0] = leftBracket
	a[1] = fmt.Sprintf("%v", (&z.l).RatString())
	if (&z.r).Sign() == -1 {
		a[2] = fmt.Sprintf("%v", (&z.r).RatString())
	} else {
		a[2] = fmt.Sprintf("+%v", (&z.r).RatString())
	}
	a[3] = omega
	a[4] = rightBracket
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Rat) Equals(y *Rat) bool {
	if z.l.Cmp(&y.l) != 0 || z.r.Cmp(&y.r) != 0 {
		return false
	}
	return true
}

// Set sets z equal to y, and returns z.
func (z *Rat) Set(y *Rat) *Rat {
	z.l.Set(&y.l)
	z.r.Set(&y.r)
	return z
}

// SetPair sets z equal to an Eisenstein number made with a given pair, and then
// it returns z.
func (z *Rat) SetPair(a, b *big.Rat) *Rat {
	z.l.Set(a)
	z.r.Set(b)
	return z
}

// NewRat returns a pointer to the Rat value a+bω.
func NewRat(a, b *big.Rat) *Rat {
	z := new(Rat)
	z.SetPair(a, b)
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
	z.l.Sub(&y.l, &y.r)
	z.r.Neg(&y.r)
	return z
}

// Add sets z equal to the sum of x and y, and returns z.
func (z *Rat) Add(x, y *Rat) *Rat {
	z.l.Add(&x.l, &y.l)
	z.r.Add(&x.r, &y.r)
	return z
}

// Sub sets z equal to the difference of x and y, and returns z.
func (z *Rat) Sub(x, y *Rat) *Rat {
	z.l.Sub(&x.l, &y.l)
	z.r.Sub(&x.r, &y.r)
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rule is:
// 		Mul(ω, ω) + ω + 1 = 0
// This binary operation is commutative and associative.
func (z *Rat) Mul(x, y *Rat) *Rat {
	a, b, temp := new(big.Rat), new(big.Rat), new(big.Rat)
	a.Sub(
		a.Mul(&x.l, &y.l),
		temp.Mul(&y.r, &x.r),
	)
	b.Add(
		b.Mul(&y.r, &x.l),
		temp.Mul(&x.r, &y.l),
	)
	b.Sub(
		b,
		temp.Mul(&x.r, &y.r),
	)
	z.SetPair(a, b)
	return z
}

// Quad returns the quadrance of z. If z = a+bω, then the quadrance is
// 		Mul(a, a) + Mul(b, b) - Mul(a, b)
// This is always non-negative.
func (z *Rat) Quad() *big.Rat {
	q, temp := new(big.Rat), new(big.Rat)
	q.Add(
		q.Mul(&z.l, &z.l),
		temp.Mul(&z.r, &z.r),
	)
	q.Sub(
		q,
		temp.Mul(&z.l, &z.r),
	)
	return q
}

// Inv sets z equal to the inverse of y, and then it returns z. If y is zero,
// then Quo will panic.
func (z *Rat) Inv(y *Rat) *Rat {
	if zero := new(Rat); y.Equals(zero) {
		panic(zeroInverse)
	}
	q := y.Quad()
	q.Inv(q)
	z.Conj(y)
	z.Scale(z, q)
	return z
}

// Quo sets z equal to the quotient of x and y, and returns z.
func (z *Rat) Quo(x, y *Rat) *Rat {
	if zero := new(Rat); y.Equals(zero) {
		panic(zeroDenominator)
	}
	z.Inv(y)
	z.Mul(x, z)
	return z
}

// Associates returns the six associates of z.
func (z *Rat) Associates() (a, b, c, d, e, f *Rat) {
	a.Set(z)
	b.Neg(z)
	unit := new(Rat)
	unit.Omega()
	c.Mul(z, unit)
	unit.Neg(unit)
	d.Mul(z, unit)
	unit.Mul(unit, unit)
	e.Mul(z, unit)
	unit.Neg(unit)
	f.Mul(z, unit)
	return
}

// Generate a random Rat value for quick.Check testing.
func (z *Rat) Generate(rand *rand.Rand, size int) reflect.Value {
	randomRat := &Rat{
		*big.NewRat(rand.Int63(), rand.Int63()),
		*big.NewRat(rand.Int63(), rand.Int63()),
	}
	return reflect.ValueOf(randomRat)
}
