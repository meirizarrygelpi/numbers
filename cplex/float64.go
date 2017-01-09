// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package cplex

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"strings"

	"github.com/meirizarrygelpi/numbers/maclaurin"
)

// A Float64 is a complex number with float64 components.
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
// If z corresponds to a + bi, then the string is "⦗a+bi⦘", similar to
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

// SetPair sets z equal to a complex number made with a given pair, and then
// it returns z.
func (z *Float64) SetPair(a, b float64) *Float64 {
	z.l = a
	z.r = b
	return z
}

// SetReal sets the real part of z equal to a, and then it returns z.
func (z *Float64) SetReal(a float64) *Float64 {
	z.l = a
	return z
}

// SetUnreal sets the unreal part of z equal to b, and then it returns z.
func (z *Float64) SetUnreal(b float64) *Float64 {
	z.r = b
	return z
}

// NewFloat64 returns a pointer to the Float64 value a+bi.
func NewFloat64(a, b float64) *Float64 {
	z := new(Float64)
	z.SetPair(a, b)
	return z
}

// Plus sets z equal to y+a, with a real, and returns z.
func (z *Float64) Plus(y *Float64, a float64) *Float64 {
	z.l = y.l + a
	z.r = y.r
	return z
}

// Minus sets z equal to y-a, with a real, and returns z.
func (z *Float64) Minus(y *Float64, a float64) *Float64 {
	z.l = y.l - a
	z.r = y.r
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
	a := (x.l * y.l) - (y.r * x.r)
	b := (x.r * y.l) + (y.r * x.l)
	z.SetPair(a, b)
	return z
}

// Quad returns the quadrance of z. If z = a+bi, then the quadrance is
// 		a² + b²
// This is always non-negative.
func (z *Float64) Quad() float64 {
	return (z.l * z.l) + (z.r * z.r)
}

// Inv sets z equal to the inverse of y, and returns z. If y is zero, then Inv
// panics.
func (z *Float64) Inv(y *Float64) *Float64 {
	if zero := new(Float64); y.Equals(zero) {
		panic(zeroInverse)
	}
	return z.Divide(z.Conj(y), y.Quad())
}

// Quo sets z equal to the quotient of x and y, and returns z. If y is zero,
// then Quo panics.
func (z *Float64) Quo(x, y *Float64) *Float64 {
	if zero := new(Float64); y.Equals(zero) {
		panic(zeroDenominator)
	}
	q := y.Quad()
	a := (x.l * y.l) + (y.r * x.r)
	b := (x.r * y.l) - (y.r * x.l)
	z.SetPair(a, b)
	return z.Divide(z, q)
}

// Gauss sets z equal to the Gaussian integer a+bi, and returns z.
func (z *Float64) Gauss(a, b int64) *Float64 {
	z.l = float64(a)
	z.r = float64(b)
	return z
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

// Maclaurin sets z equal to the value of the Maclaurin polynomial p evaluated
// at y, and returns z. Horner's method is used.
func (z *Float64) Maclaurin(y *Float64, p maclaurin.Float64) *Float64 {
	if p.Len() == 0 {
		z = new(Float64)
		return z
	}
	n := p.Degree
	var a float64
	if n == 0 {
		z = new(Float64)
		a, _ = p.Coeff(n)
		z.SetReal(a)
		return z
	}
	a, _ = p.Coeff(n)
	z.Dilate(y, a)
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

// Generate returns a random Float64 value for quick.Check testing.
func (z *Float64) Generate(rand *rand.Rand, size int) reflect.Value {
	randomFloat64 := &Float64{
		rand.Float64(),
		rand.Float64(),
	}
	return reflect.ValueOf(randomFloat64)
}
