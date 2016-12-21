// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package pplex

import (
	"fmt"
	"math/rand"
	"reflect"
	"strings"
)

// An Int64 is a perplex number with int64 components.
type Int64 struct {
	l, r int64
}

// One sets z equal to 1, and then it returns z.
func (z *Int64) One() *Int64 {
	z.l = 1
	z.r = 0
	return z
}

// Real returns the real part of z.
func (z *Int64) Real() int64 {
	return z.l
}

// Unreal returns the unreal part of z.
func (z *Int64) Unreal() int64 {
	return z.r
}

// String returns the string version of an Int64 value.
//
// If z corresponds to a + bs, then the string is "⦗a+bs⦘", similar to
// complex128 values.
func (z *Int64) String() string {
	a := make([]string, 5)
	a[0] = leftBracket
	a[1] = fmt.Sprint(z.l)
	if z.r < 0 {
		a[2] = fmt.Sprint(z.r)
	} else {
		a[2] = "+" + fmt.Sprint(z.r)
	}
	a[3] = unit
	a[4] = rightBracket
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Int64) Equals(y *Int64) bool {
	return z.l == y.l && z.r == y.r
}

// Set sets z equal to y, and returns z.
func (z *Int64) Set(y *Int64) *Int64 {
	z.l = y.l
	z.r = y.r
	return z
}

// SetPair sets z equal to a perplex number made with a given pair, and then
// it returns z.
func (z *Int64) SetPair(a, b int64) *Int64 {
	z.l = a
	z.r = b
	return z
}

// NewInt64 returns a pointer to the Int64 value a+bs.
func NewInt64(a, b int64) *Int64 {
	z := new(Int64)
	z.SetPair(a, b)
	return z
}

// Dilate sets z equal to y dilated by a, and returns z.
func (z *Int64) Dilate(y *Int64, a int64) *Int64 {
	z.l = y.l * a
	z.r = y.r * a
	return z
}

// Divide sets z equal to y contracted by a, and returns z.
func (z *Int64) Divide(y *Int64, a int64) *Int64 {
	z.l = y.l / a
	z.r = y.r / a
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Int64) Neg(y *Int64) *Int64 {
	z.l = -y.l
	z.r = -y.r
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Int64) Conj(y *Int64) *Int64 {
	z.l = y.l
	z.r = -y.r
	return z
}

// Add sets z equal to x+y, and returns z.
func (z *Int64) Add(x, y *Int64) *Int64 {
	z.l = x.l + y.l
	z.r = x.r + y.r
	return z
}

// Sub sets z equal to x-y, and returns z.
func (z *Int64) Sub(x, y *Int64) *Int64 {
	z.l = x.l - y.l
	z.r = x.r - y.r
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rule is:
// 		Mul(s, s) = +1
// This binary operation is commutative and associative.
func (z *Int64) Mul(x, y *Int64) *Int64 {
	a := (x.l * y.l) + (y.r * x.r)
	b := (y.r * x.l) + (x.r * y.l)
	z.SetPair(a, b)
	return z
}

// Quad returns the quadrance of z. If z = a+bs, then the quadrance is
// 		a² - b²
// This can be positive, negative, or zero.
func (z *Int64) Quad() int64 {
	return (z.l * z.l) - (z.r * z.r)
}

// IsZeroDivisor returns true if z is a zero divisor. This is equivalent to z
// having zero quadrance.
func (z *Int64) IsZeroDivisor() bool {
	return z.l == z.r || z.l == -z.r
}

// Quo sets z equal to the quotient of x and y, and returns z. If y is zero,
// then Quo panics.
func (z *Int64) Quo(x, y *Int64) *Int64 {
	if y.IsZeroDivisor() {
		panic(zeroDivisorDenominator)
	}
	return z.Divide(z.Mul(x, z.Conj(y)), y.Quad())
}

// Generate returns a random Int64 value for quick.Check testing.
func (z *Int64) Generate(rand *rand.Rand, size int) reflect.Value {
	randomInt64 := &Int64{
		rand.Int63(),
		rand.Int63(),
	}
	return reflect.ValueOf(randomInt64)
}
