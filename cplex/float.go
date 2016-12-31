// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package cplex

import (
	"math/big"
	"math/rand"
	"reflect"
	"strings"
)

// A Float is a complex number with big.Float components.
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

// Prec returns the precision of z in bits of the real part of z.
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
// If z corresponds to a + bi, then the string is "⦗a+bi⦘", similar to
// complex128 values.
func (z *Float) String() string {
	a := make([]string, 5)
	a[0] = leftBracket
	a[1] = z.l.String()
	a[2] = sprintFloat(&z.r)
	a[3] = unitName
	a[4] = rightBracket
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Float) Equals(y *Float) bool {
	return z.l.Cmp(&y.l) == 0 && z.r.Cmp(&y.r) == 0
}

// Set sets z equal to y, and returns z.
func (z *Float) Set(y *Float) *Float {
	z.l.Set(&y.l)
	z.r.Set(&y.r)
	return z
}

// SetPair sets z equal to a complex number made with a given pair, and then
// it returns z.
func (z *Float) SetPair(a, b *big.Float) *Float {
	z.l.Set(a)
	z.r.Set(b)
	return z
}

// NewFloat returns a pointer to the Float value a+bi.
func NewFloat(a, b *big.Float) *Float {
	z := new(Float)
	z.SetPair(a, b)
	return z
}

// Dilate sets z equal to y dilated by a, and returns z.
func (z *Float) Dilate(y *Float, a *big.Float) *Float {
	z.l.Mul(&y.l, a)
	z.r.Mul(&y.r, a)
	return z
}

// Divide sets z equal to y contracted by a, and returns z.
func (z *Float) Divide(y *Float, a *big.Float) *Float {
	z.l.Quo(&y.l, a)
	z.r.Quo(&y.r, a)
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
	z.l.Set(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Add sets z equal to x+y, and returns z.
func (z *Float) Add(x, y *Float) *Float {
	z.l.Add(&x.l, &y.l)
	z.r.Add(&x.r, &y.r)
	return z
}

// Sub sets z equal to x-y, and returns z.
func (z *Float) Sub(x, y *Float) *Float {
	z.l.Sub(&x.l, &y.l)
	z.r.Sub(&x.r, &y.r)
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
func (z *Float) Mul(x, y *Float) *Float {
	a, b, temp := new(big.Float), new(big.Float), new(big.Float)
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

// Quad returns the quadrance of z. If z = a+bi, then the quadrance is
// 		a² + b²
// This is always non-negative.
func (z *Float) Quad() *big.Float {
	quad := new(big.Float)
	return quad.Add(
		quad.Mul(&z.l, &z.l),
		new(big.Float).Mul(&z.r, &z.r),
	)
}

// Inv sets z equal to the inverse of y, and returns z. If y is zero, then Inv
// panics.
func (z *Float) Inv(y *Float) *Float {
	if zero := new(Float).Zero(); y.Equals(zero) {
		panic(zeroInverse)
	}
	return z.Divide(z.Conj(y), y.Quad())
}

// Quo sets z equal to the quotient of x and y, and returns z. If y is zero,
// then Quo panics.
func (z *Float) Quo(x, y *Float) *Float {
	if zero := new(Float).Zero(); y.Equals(zero) {
		panic(zeroDenominator)
	}
	q := y.Quad()
	a, b, temp := new(big.Float), new(big.Float), new(big.Float)
	a.Add(
		a.Mul(&x.l, &y.l),
		temp.Mul(&y.r, &x.r),
	)
	b.Sub(
		temp.Mul(&x.r, &y.l),
		b.Mul(&y.r, &x.l),
	)
	z.SetPair(a, b)
	return z.Divide(z, q)
}

// Gauss sets z equal to the Gaussian integer a+bi, and returns z.
func (z *Float) Gauss(a, b *big.Int) *Float {
	z.l.SetInt(a)
	z.r.SetInt(b)
	return z
}

// CrossRatio sets z equal to the cross-ratio of v, w, x, and y:
// 		Inv(w - x) * (v - x) * Inv(v - y) * (w - y)
// Then it returns z.
func (z *Float) CrossRatio(v, w, x, y *Float) *Float {
	temp := new(Float)
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
func (z *Float) Möbius(y, a, b, c, d *Float) *Float {
	z.Mul(a, y)
	z.Add(z, b)
	temp := new(Float)
	temp.Mul(c, y)
	temp.Add(temp, d)
	temp.Inv(temp)
	return z.Mul(z, temp)
}

// Generate returns a random Float value for quick.Check testing.
func (z *Float) Generate(rand *rand.Rand, size int) reflect.Value {
	randomFloat := &Float{
		*big.NewFloat(rand.Float64()),
		*big.NewFloat(rand.Float64()),
	}
	return reflect.ValueOf(randomFloat)
}
