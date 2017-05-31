// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package pplex

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"strings"
)

// A Float64 is a perplex number with float64 components.
type Float64 struct {
	l, r float64
}

// One sets z equal to 1, and then it returns z.
func (z *Float64) One() *Float64 {
	z.l = 1
	z.r = 0
	return z
}

// Real returns the real part of z.
func (z *Float64) Real() float64 {
	return z.l
}

// Unreal returns the unreal part of z.
func (z *Float64) Unreal() float64 {
	return z.r
}

func sprintFloat64(a float64) string {
	if math.Signbit(a) {
		return fmt.Sprintf("%g", a)
	}
	if math.IsInf(a, +1) {
		return "+Inf"
	}
	return fmt.Sprintf("+%g", a)
}

// String returns the string version of a Float64 value.
//
// If z corresponds to a + bs, then the string is "⦗a+bs⦘", similar to
// complex128 values.
func (z *Float64) String() string {
	a := make([]string, 5)
	a[0] = leftBracket
	a[1] = fmt.Sprintf("%g", z.l)
	a[2] = sprintFloat64(z.r)
	a[3] = unitName
	a[4] = rightBracket
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Float64) Equals(y *Float64) bool {
	return z.l == y.l && z.r == y.r
}

// Set sets z equal to y, and returns z.
func (z *Float64) Set(y *Float64) *Float64 {
	z.l = y.l
	z.r = y.r
	return z
}

// SetPair sets z equal to a perplex number made with a given pair, and then
// it returns z.
func (z *Float64) SetPair(a, b float64) *Float64 {
	z.l = a
	z.r = b
	return z
}

// NewFloat64 returns a pointer to the Float64 value a+bs.
func NewFloat64(a, b float64) *Float64 {
	z := new(Float64)
	z.SetPair(a, b)
	return z
}

// Dilate sets z equal to y dilated by a, and returns z.
func (z *Float64) Dilate(y *Float64, a float64) *Float64 {
	z.l = y.l * a
	z.r = y.r * a
	return z
}

// Divide sets z equal to y contracted by a, and returns z.
func (z *Float64) Divide(y *Float64, a float64) *Float64 {
	z.l = y.l / a
	z.r = y.r / a
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Float64) Neg(y *Float64) *Float64 {
	z.l = -y.l
	z.r = -y.r
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Float64) Conj(y *Float64) *Float64 {
	z.l = y.l
	z.r = -y.r
	return z
}

// Add sets z equal to x+y, and returns z.
func (z *Float64) Add(x, y *Float64) *Float64 {
	z.l = x.l + y.l
	z.r = x.r + y.r
	return z
}

// Sub sets z equal to x-y, and returns z.
func (z *Float64) Sub(x, y *Float64) *Float64 {
	z.l = x.l - y.l
	z.r = x.r - y.r
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
func (z *Float64) Mul(x, y *Float64) *Float64 {
	a := (x.l * y.l) + (y.r * x.r)
	b := (x.r * y.l) + (y.r * x.l)
	z.SetPair(a, b)
	return z
}

// Dot returns the dot product of y and z. If z = a+bs and y = c+ds, then the
// dot product is
// 		ac - bd
// This can be positive, negative, or zero. The dot product is equivalent to
// 		½(Mul(Conj(z), y) + Mu(Conj(y), z))
// In this form it is clear that Dot is symmetric.
func (z *Float64) Dot(y *Float64) float64 {
	return (z.l * y.l) - (z.r * y.r)
}

// Quad returns the quadrance of z. If z = a+bs, then the quadrance is
// 		a² - b²
// This can be positive, negative, or zero.
func (z *Float64) Quad() float64 {
	return z.Dot(z)
}

// Cross returns the cross product of y and z. If z = a+bs and y = c+ds, then
// the cross product is
// 		ad - bc
// This can be positive, negative, or zero. The cross product is equivalent to
// the unreal part of
// 		½(Mul(Conj(z), y) - Mu(Conj(y), z))
// In this form it is clear that Cross is anti-symmetric.
func (z *Float64) Cross(y *Float64) float64 {
	return (z.l * y.r) - (z.r * y.l)
}

// IsZeroDivisor returns true if z is a zero divisor. This is equivalent to z
// having zero quadrance.
func (z *Float64) IsZeroDivisor() bool {
	return z.l == z.r || z.l == -z.r
}

// Inv sets z equal to the inverse of y, and returns z. If y is a zero divisor,
// then Inv panics.
func (z *Float64) Inv(y *Float64) *Float64 {
	if y.IsZeroDivisor() {
		panic(zeroDivisorInverse)
	}
	return z.Divide(z.Conj(y), y.Quad())
}

// Quo sets z equal to the quotient of x and y, and returns z. If y is a zero
// divisor, then Quo panics.
func (z *Float64) Quo(x, y *Float64) *Float64 {
	if y.IsZeroDivisor() {
		panic(zeroDivisorDenominator)
	}
	q := y.Quad()
	a := (x.l * y.l) - (y.r * x.r)
	b := (x.r * y.l) - (y.r * x.l)
	z.SetPair(a, b)
	return z.Divide(z, q)
}

// CrossRatio sets z equal to the cross-ratio of v, w, x, and y:
// 		Inv(w - x) * (v - x) * Inv(v - y) * (w - y)
// Then it returns z.
func (z *Float64) CrossRatio(v, w, x, y *Float64) *Float64 {
	temp := new(Float64)
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
func (z *Float64) Möbius(y, a, b, c, d *Float64) *Float64 {
	z.Mul(a, y)
	z.Add(z, b)
	temp := new(Float64)
	temp.Mul(c, y)
	temp.Add(temp, d)
	temp.Inv(temp)
	return z.Mul(z, temp)
}

// Generate returns a random Float64 value for quick.Check testing.
func (z *Float64) Generate(rand *rand.Rand, size int) reflect.Value {
	randomFloat64 := &Float64{
		rand.Float64(),
		rand.Float64(),
	}
	return reflect.ValueOf(randomFloat64)
}
