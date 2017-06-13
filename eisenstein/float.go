// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package eisenstein

import (
	"math/big"
	"math/rand"
	"reflect"
	"strings"
)

// A Float represents an arbitrary-precision Eisenstein float.
type Float struct {
	l, r big.Float
}

// Real returns the real part of z.
func (z *Float) Real() *big.Float {
	return &z.l
}

// Unreal returns the unreal part of z.
func (z *Float) Unreal() *big.Float {
	return &z.r
}

// Acc returns the accuracy of the real part of z.
func (z *Float) Acc() big.Accuracy {
	return z.l.Acc()
}

// Mode returns the rounding mode of the real part of z.
func (z *Float) Mode() big.RoundingMode {
	return z.l.Mode()
}

// Prec returns the precision in bits of the real part of z.
func (z *Float) Prec() uint {
	return z.l.Prec()
}

// SetMode sets the rounding mode of z, and then it returns z.
func (z *Float) SetMode(mode big.RoundingMode) *Float {
	z.l.SetMode(mode)
	z.r.SetMode(mode)
	return z
}

// SetPrec sets the precision of z, and then it returns z.
func (z *Float) SetPrec(prec uint) *Float {
	z.l.SetPrec(prec)
	z.r.SetPrec(prec)
	return z
}

// Zero sets z equal to 0, and then it returns z. Each component has 64 bits of
// precision.
func (z *Float) Zero() *Float {
	z.l.SetInt64(0)
	z.r.SetInt64(0)
	return z
}

// One sets z equal to 1, and then it returns z. Each component has 64 bits of
// precision.
func (z *Float) One() *Float {
	z.l.SetInt64(1)
	z.r.SetInt64(0)
	return z
}

// Omega sets z equal to the Eisenstein unit ω, and then it returns z.
func (z *Float) Omega() *Float {
	z.l.SetInt64(0)
	z.r.SetInt64(1)
	return z
}

func sprintFloat(a *big.Float) string {
	if a.Signbit() {
		return a.String()
	}
	if a.IsInf() {
		return "+Inf"
	}
	return "+" + a.String()
}

// String returns the string version of a Float value.
//
// If z corresponds to a + bω, then the string is "(a+bω)", similar to
// complex128 values.
func (z *Float) String() string {
	a := make([]string, 5)
	a[0] = leftBracket
	a[1] = z.l.String()
	a[2] = sprintFloat(&z.r)
	a[3] = omega
	a[4] = rightBracket
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Float) Equals(y *Float) bool {
	if z.l.Cmp(&y.l) != 0 || z.r.Cmp(&y.r) != 0 {
		return false
	}
	return true
}

// Set sets z equal to y, and then it returns z.
func (z *Float) Set(y *Float) *Float {
	z.l.Set(&y.l)
	z.r.Set(&y.r)
	return z
}

// SetPair sets z equal to an Eisenstein number made with a given pair, and then
// it returns z.
func (z *Float) SetPair(a, b *big.Float) *Float {
	z.l.Set(a)
	z.r.Set(b)
	return z
}

// NewFloat returns a pointer to the Float value a+bω.
func NewFloat(a, b *big.Float) *Float {
	z := new(Float)
	z.SetPair(a, b)
	return z
}

// Scale sets z equal to y scaled by a, and returns z.
func (z *Float) Scale(y *Float, a *big.Float) *Float {
	z.l.Mul(&y.l, a)
	z.r.Mul(&y.r, a)
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Float) Neg(y *Float) *Float {
	z.l.Neg(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Float) Conj(y *Float) *Float {
	z.l.Sub(&y.l, &y.r)
	z.r.Neg(&y.r)
	return z
}

// Add sets z equal to the sum of x and y, and returns z.
func (z *Float) Add(x, y *Float) *Float {
	z.l.Add(&x.l, &y.l)
	z.r.Add(&x.r, &y.r)
	return z
}

// Sub sets z equal to the difference of x and y, and returns z.
func (z *Float) Sub(x, y *Float) *Float {
	z.l.Sub(&x.l, &y.l)
	z.r.Sub(&x.r, &y.r)
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rule is:
// 		Mul(ω, ω) + ω + 1 = 0
// This binary operation is commutative and associative.
func (z *Float) Mul(x, y *Float) *Float {
	a, b, temp := new(big.Float), new(big.Float), new(big.Float)
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
func (z *Float) Quad() *big.Float {
	quad, temp := new(big.Float), new(big.Float)
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

// Inv sets z equal to the inverse of y, and returns z.
func (z *Float) Inv(y *Float) *Float {
	if zero := new(Float).Zero(); y.Equals(zero) {
		panic(zeroInverse)
	}
	quad := y.Quad()
	z.Conj(y)
	z.l.Quo(&z.l, quad)
	z.r.Quo(&z.r, quad)
	return z
}

// Quo sets z equal to the quotient of x and y, and returns z.
func (z *Float) Quo(x, y *Float) *Float {
	if zero := new(Float).Zero(); y.Equals(zero) {
		panic(zeroDenominator)
	}
	z.Inv(y)
	z.Mul(x, z)
	return z
}

// Associates returns the six associates of z.
func (z *Float) Associates() (a, b, c, d, e, f *Float) {
	a.Set(z)
	b.Neg(z)
	unit := new(Float)
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

// Generate a random Float value for quick.Check testing.
func (z *Float) Generate(rand *rand.Rand, size int) reflect.Value {
	randomFloat := &Float{
		*big.NewFloat(rand.Float64()),
		*big.NewFloat(rand.Float64()),
	}
	return reflect.ValueOf(randomFloat)
}
