// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package eisenstein

import (
	"fmt"
	"math/rand"
	"reflect"
	"strings"
)

// An Int64 represents an Eisenstein number with int64 components.
type Int64 struct {
	l, r int64
}

// One sets z equal to 1, and then it returns z.
func (z *Int64) One() *Int64 {
	z.l = 1
	z.r = 0
	return z
}

// Omega sets z equal to the Eisenstein unit ω, and then it returns z.
func (z *Int64) Omega() *Int64 {
	z.l = 0
	z.r = 1
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
// If z corresponds to a + bω, then the string is "(a+bω)", similar to
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
	a[3] = omega
	a[4] = rightBracket
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Int64) Equals(y *Int64) bool {
	return z.l == y.l && z.r == y.r
}

// Set sets z equal to y and then returns z.
func (z *Int64) Set(y *Int64) *Int64 {
	z.l = y.l
	z.r = y.r
	return z
}

// SetPair sets z equal to an Eisenstein number made with a given pair, and then
// it returns z.
func (z *Int64) SetPair(a, b int64) *Int64 {
	z.l = a
	z.r = b
	return z
}

// NewInt64 returns a pointer to an Int64 with given components.
func NewInt64(a, b int64) *Int64 {
	z := new(Int64)
	z.SetPair(a, b)
	return z
}

// Scale sets z equal to the scaling of y by a, and then it returns z.
func (z *Int64) Scale(y *Int64, a int64) *Int64 {
	z.l = y.l * a
	z.r = y.r * a
	return z
}

// Neg sets z equal to the negation of y, and then it returns z.
func (z *Int64) Neg(y *Int64) *Int64 {
	z.l = -y.l
	z.r = -y.r
	return z
}

// Conj sets z equal to the conjugate of y, and then it returns z.
func (z *Int64) Conj(y *Int64) *Int64 {
	z.l = y.l - y.r
	z.r = -y.r
	return z
}

// Add sets z equal to the addition of x and y, and then it returns z.
func (z *Int64) Add(x, y *Int64) *Int64 {
	z.l = x.l + y.l
	z.r = x.r + y.r
	return z
}

// Sub sets z equal to the subtraction of y from x, and then it returns z.
func (z *Int64) Sub(x, y *Int64) *Int64 {
	z.l = x.l - y.l
	z.r = x.r - y.r
	return z
}

// Mul sets z equal to the multiplication of x and y, and then it returns z.
func (z *Int64) Mul(x, y *Int64) *Int64 {
	a := (x.l * y.l) - (y.r * x.r)
	b := (y.r * x.l) + (x.r * (y.l - y.r))
	z.SetPair(a, b)
	return z
}

// Quad returns the quadrance of z, a non-negative number.
func (z *Int64) Quad() int64 {
	return (z.l * z.l) + (z.r * z.r) - (z.l * z.r)
}

// Quo sets z equal to the quotient of x by y, and then it returns z. If y is
// zero, then Quo will panic.
func (z *Int64) Quo(x, y *Int64) *Int64 {
	if zero := new(Int64); y.Equals(zero) {
		panic(zeroDenominator)
	}
	q := y.Quad()
	z.Conj(y)
	z.Mul(x, z)
	z.l = z.l / q
	z.r = z.r / q
	return z
}

// Associates returns the six associates of z.
func (z *Int64) Associates() (a, b, c, d, e, f *Int64) {
	a.Set(z)
	b.Neg(z)
	unit := new(Int64)
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

// Generate a random Int64 value for quick.Check testing.
func (z *Int64) Generate(rand *rand.Rand, size int) reflect.Value {
	randomInt64 := &Int64{rand.Int63(), rand.Int63()}
	return reflect.ValueOf(randomInt64)
}
