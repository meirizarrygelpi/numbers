// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package eisenstein

import (
	"math/big"
	"math/rand"
	"reflect"
	"strings"
)

// An Int represents an arbitrary-precision Eisenstein integer.
type Int struct {
	l, r big.Int
}

// One sets z equal to 1, and then it returns z.
func (z *Int) One() *Int {
	z.l.SetInt64(1)
	z.r.SetInt64(0)
	return z
}

// Omega sets z equal to the Eisenstein unit ω, and then it returns z.
func (z *Int) Omega() *Int {
	z.l.SetInt64(0)
	z.r.SetInt64(1)
	return z
}

// Real returns the real part of z.
func (z *Int) Real() *big.Int {
	return &z.l
}

// Unreal returns the unreal part of z.
func (z *Int) Unreal() *big.Int {
	return &z.r
}

// String returns the string version of an Int value.
//
// If z corresponds to a + bω, then the string is "⦗a+bω⦘", similar to
// complex128 values.
func (z *Int) String() string {
	a := make([]string, 5)
	a[0] = leftBracket
	a[1] = z.l.String()
	if z.r.Sign() < 0 {
		a[2] = z.r.String()
	} else {
		a[2] = "+" + z.r.String()
	}
	a[3] = omega
	a[4] = rightBracket
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Int) Equals(y *Int) bool {
	return z.l.Cmp(&y.l) == 0 && z.r.Cmp(&y.r) == 0
}

// Set sets z equal to y, and returns z.
func (z *Int) Set(y *Int) *Int {
	z.l.Set(&y.l)
	z.r.Set(&y.r)
	return z
}

// SetPair sets z equal to an Eisenstein number made with a given pair, and then
// it returns z.
func (z *Int) SetPair(a, b *big.Int) *Int {
	z.l.Set(a)
	z.r.Set(b)
	return z
}

// NewInt returns a pointer to the Int value a+bω.
func NewInt(a, b *big.Int) *Int {
	z := new(Int)
	z.SetPair(a, b)
	return z
}

// Scale sets z equal to y scaled by a, and returns z.
func (z *Int) Scale(y *Int, a *big.Int) *Int {
	z.l.Mul(&y.l, a)
	z.r.Mul(&y.r, a)
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Int) Neg(y *Int) *Int {
	z.l.Neg(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Int) Conj(y *Int) *Int {
	z.l.Sub(&y.l, &y.r)
	z.r.Neg(&y.r)
	return z
}

// Add sets z equal to the sum of x and y, and returns z.
func (z *Int) Add(x, y *Int) *Int {
	z.l.Add(&x.l, &y.l)
	z.r.Add(&x.r, &y.r)
	return z
}

// Sub sets z equal to the difference of x and y, and returns z.
func (z *Int) Sub(x, y *Int) *Int {
	z.l.Sub(&x.l, &y.l)
	z.r.Sub(&x.r, &y.r)
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rule is:
// 		Mul(ω, ω) + ω + 1 = 0
// This binary operation is commutative and associative.
func (z *Int) Mul(x, y *Int) *Int {
	a, b, temp := new(big.Int), new(big.Int), new(big.Int)
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
func (z *Int) Quad() *big.Int {
	quad, temp := new(big.Int), new(big.Int)
	quad.Add(
		quad.Mul(&z.l, &z.l),
		temp.Mul(&z.r, &z.r),
	)
	quad.Sub(
		quad,
		temp.Mul(&z.l, &z.r),
	)
	return quad
}

// Quo sets z equal to the quotient of x and y, and returns z. Note that
// truncated division is used.
func (z *Int) Quo(x, y *Int) *Int {
	if zero := new(Int); y.Equals(zero) {
		panic(zeroDenominator)
	}
	quad := y.Quad()
	z.Conj(y)
	z.Mul(x, z)
	z.l.Quo(&z.l, quad)
	z.r.Quo(&z.r, quad)
	return z
}

// Associates returns the six associates of z.
func (z *Int) Associates() (a, b, c, d, e, f *Int) {
	a.Set(z)
	b.Neg(z)
	unit := new(Int)
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

// Generate a random Int value for quick.Check testing.
func (z *Int) Generate(rand *rand.Rand, size int) reflect.Value {
	randomInt := &Int{
		*big.NewInt(rand.Int63()),
		*big.NewInt(rand.Int63()),
	}
	return reflect.ValueOf(randomInt)
}
