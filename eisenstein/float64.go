// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package eisenstein

import (
	"fmt"
	"math/rand"
	"reflect"
	"strings"
)

// A Float64 represents an Eisenstein number with float64 components.
type Float64 struct {
	l, r float64
}

// One sets z equal to 1, and then it returns z.
func (z *Float64) One() *Float64 {
	z.l = 1
	z.r = 0
	return z
}

// Omega sets z equal to the Eisenstein unit ω, and then it returns z.
func (z *Float64) Omega() *Float64 {
	z.l = 0
	z.r = 1
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

// String returns the string version of an Float64 value.
//
// If z corresponds to a + bω, then the string is "⦗a+bω⦘", similar to
// complex128 values.
func (z *Float64) String() string {
	a := make([]string, 5)
	a[0] = leftBracket
	a[1] = fmt.Sprintf("%v", z.l)
	if z.r < 0 {
		a[2] = fmt.Sprintf("%v", z.r)
	} else {
		a[2] = fmt.Sprintf("+%v", z.r)
	}
	a[3] = omega
	a[4] = rightBracket
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Float64) Equals(y *Float64) bool {
	return z.l == y.l && z.r == y.r
}

// Set sets z equal to y and then returns z.
func (z *Float64) Set(y *Float64) *Float64 {
	z.l = y.l
	z.r = y.r
	return z
}

// SetPair sets z equal to an Eisenstein number made with a given pair, and then
// it returns z.
func (z *Float64) SetPair(a, b float64) *Float64 {
	z.l = a
	z.r = b
	return z
}

// NewFloat64 returns a pointer to an Float64 with given components.
func NewFloat64(a, b float64) *Float64 {
	z := new(Float64)
	z.SetPair(a, b)
	return z
}

// Scale sets z equal to the scaling of y by a, and then it returns z.
func (z *Float64) Scale(y *Float64, a float64) *Float64 {
	z.l = y.l * a
	z.r = y.r * a
	return z
}

// Neg sets z equal to the negation of y, and then it returns z.
func (z *Float64) Neg(y *Float64) *Float64 {
	z.l = -y.l
	z.r = -y.r
	return z
}

// Conj sets z equal to the conjugate of y, and then it returns z.
func (z *Float64) Conj(y *Float64) *Float64 {
	z.l = y.l - y.r
	z.r = -y.r
	return z
}

// Add sets z equal to the addition of x and y, and then it returns z.
func (z *Float64) Add(x, y *Float64) *Float64 {
	z.l = x.l + y.l
	z.r = x.r + y.r
	return z
}

// Sub sets z equal to the subtraction of y from x, and then it returns z.
func (z *Float64) Sub(x, y *Float64) *Float64 {
	z.l = x.l - y.l
	z.r = x.r - y.r
	return z
}

// Mul sets z equal to the multiplication of x and y, and then it returns z.
func (z *Float64) Mul(x, y *Float64) *Float64 {
	a := (x.l * y.l) - (y.r * x.r)
	b := (y.r * x.l) + (x.r * (y.l - y.r))
	z.SetPair(a, b)
	return z
}

// Quad returns the quadrance of z, a non-negative number.
func (z *Float64) Quad() float64 {
	return (z.l * z.l) + (z.r * z.r) - (z.l * z.r)
}

// Inv sets z equal to the inverse of y, and then it returns z. If y is zero,
// then Quo will panic.
func (z *Float64) Inv(y *Float64) *Float64 {
	if zero := new(Float64); y.Equals(zero) {
		panic(zeroInverse)
	}
	q := y.Quad()
	z.Conj(y)
	z.l = z.l / q
	z.r = z.r / q
	return z
}

// Quo sets z equal to the quotient of x by y, and then it returns z. If y is
// zero, then Quo will panic.
func (z *Float64) Quo(x, y *Float64) *Float64 {
	if zero := new(Float64); y.Equals(zero) {
		panic(zeroDenominator)
	}
	z.Inv(y)
	z.Mul(x, z)
	return z
}

// Associates returns the six associates of z.
func (z *Float64) Associates() (a, b, c, d, e, f *Float64) {
	a.Set(z)
	b.Neg(z)
	unit := new(Float64)
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

// Generate a random Float64 value for quick.Check testing.
func (z *Float64) Generate(rand *rand.Rand, size int) reflect.Value {
	randomFloat64 := &Float64{rand.Float64(), rand.Float64()}
	return reflect.ValueOf(randomFloat64)
}
