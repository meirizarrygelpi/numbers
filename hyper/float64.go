// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package hyper

import (
	"math/rand"
	"reflect"
	"strings"

	"fmt"

	"math"

	"github.com/meirizarrygelpi/numbers/nplex"
	"github.com/meirizarrygelpi/numbers/vec3"
)

// A Float64 is a hyper number with float64 components.
type Float64 struct {
	l, r nplex.Float64
}

// One sets z equal to 1, and then returns z.
func (z *Float64) One() *Float64 {
	z.l.One()
	z.r.Set(new(nplex.Float64))
	return z
}

// Real returns the real part of z.
func (z *Float64) Real() float64 {
	return z.l.Real()
}

// Unreal returns the unreal part of z.
func (z *Float64) Unreal() *vec3.Float64 {
	v := new(vec3.Float64)
	v[0] = z.l.Unreal()
	v[1] = z.r.Real()
	v[2] = z.r.Unreal()
	return v
}

// String returns the string version of a Float64 value.
//
// If z corresponds to a+bα+cΓ+dαΓ, then the string is "⦗a+bα+cΓ+dαΓ⦘", similar
// to complex128 values.
func (z *Float64) String() string {
	v := z.Unreal()
	a := make([]string, 9)
	a[0] = leftBracket
	a[1] = fmt.Sprint(z.l.Real())
	if math.Signbit(v[0]) {
		a[2] = fmt.Sprint(v[0])
	} else {
		a[2] = "+" + fmt.Sprint(v[0])
	}
	a[3] = unit1
	if math.Signbit(v[1]) {
		a[4] = fmt.Sprint(v[1])
	} else {
		a[4] = "+" + fmt.Sprint(v[1])
	}
	a[5] = unit2
	if math.Signbit(v[2]) {
		a[6] = fmt.Sprint(v[2])
	} else {
		a[6] = "+" + fmt.Sprint(v[2])
	}
	a[7] = unit3
	a[8] = rightBracket
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Float64) Equals(y *Float64) bool {
	return z.l.Equals(&y.l) && z.r.Equals(&y.r)
}

// Set sets z equal to y, and returns z.
func (z *Float64) Set(y *Float64) *Float64 {
	z.l.Set(&y.l)
	z.r.Set(&y.r)
	return z
}

// SetPair sets z equal to a hyper number made with a given pair, and
// then it returns z.
func (z *Float64) SetPair(a, b *nplex.Float64) *Float64 {
	z.l.Set(a)
	z.r.Set(b)
	return z
}

// NewFloat64 returns a pointer to the Float64 value a+bα+cΓ+dαΓ.
func NewFloat64(a, b, c, d float64) *Float64 {
	z := new(Float64)
	z.l.SetPair(a, b)
	z.r.SetPair(c, d)
	return z
}

// Dilate sets z equal to y dilated by a, and returns z.
func (z *Float64) Dilate(y *Float64, a float64) *Float64 {
	z.l.Dilate(&y.l, a)
	z.r.Dilate(&y.r, a)
	return z
}

// Divide sets z equal to y contracted by a, and returns z.
func (z *Float64) Divide(y *Float64, a float64) *Float64 {
	z.l.Divide(&y.l, a)
	z.r.Divide(&y.r, a)
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Float64) Neg(y *Float64) *Float64 {
	z.l.Neg(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Star1 sets z equal to the α-conjugate of y, and returns z.
func (z *Float64) Star1(y *Float64) *Float64 {
	z.l.Conj(&y.l)
	z.r.Conj(&y.r)
	return z
}

// Star2 sets z equal to the Γ-conjugate of y, and returns z.
func (z *Float64) Star2(y *Float64) *Float64 {
	z.l.Set(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Add sets z equal to x+y, and returns z.
func (z *Float64) Add(x, y *Float64) *Float64 {
	z.l.Add(&x.l, &y.l)
	z.r.Add(&x.r, &y.r)
	return z
}

// Sub sets z equal to x-y, and returns z.
func (z *Float64) Sub(x, y *Float64) *Float64 {
	z.l.Sub(&x.l, &y.l)
	z.r.Sub(&x.r, &y.r)
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication table is:
//     +-----+----+----+----+
//     | Mul | α  | Γ  | αΓ |
//     +-----+----+----+----+
//     | α   | 0  | αΓ | 0  |
//     +-----+----+----+----+
//     | Γ   | αΓ | 0  | 0  |
//     +-----+----+----+----+
//     | αΓ  | 0  | 0  | 0  |
//     +-----+----+----+----+
// This binary operation is commutative and associative.
func (z *Float64) Mul(x, y *Float64) *Float64 {
	a, b, temp := new(nplex.Float64), new(nplex.Float64), new(nplex.Float64)
	a.Mul(&x.l, &y.l)
	b.Add(
		b.Mul(&x.l, &y.r),
		temp.Mul(&x.r, &y.l),
	)
	z.SetPair(a, b)
	return z
}

// Quad returns the quadrance of z. If z = a+bα+cΓ+dαΓ, then the quadrance is
//     a² + 2abα
// Note that this is a nilplex number.
func (z *Float64) Quad() *nplex.Float64 {
	q := new(nplex.Float64)
	return q.Mul(&z.l, &z.l)
}

// Norm returns the norm of z. If z = a+bα+cΓ+dαΓ, then the norm is
// 		(a²)²
// This is always non-negative.
func (z *Float64) Norm() float64 {
	return z.Quad().Quad()
}

// IsZeroDivisor returns true if z is a zero divisor.
func (z *Float64) IsZeroDivisor() bool {
	return z.Quad().IsZeroDivisor()
}

// Inv sets z equal to the inverse of y, and returns z. If y is zero, then Inv
// panics.
func (z *Float64) Inv(y *Float64) *Float64 {
	if y.IsZeroDivisor() {
		panic(zeroDivisorInverse)
	}
	a := y.Quad()
	a.Inv(a)
	z.Star2(y)
	z.l.Mul(&z.l, a)
	z.r.Mul(&z.r, a)
	return z
}

// Quo sets z equal to the quotient of x and y, and returns z. If y is zero,
// then Quo panics.
func (z *Float64) Quo(x, y *Float64) *Float64 {
	if y.IsZeroDivisor() {
		panic(zeroDivisorDenominator)
	}
	z.Mul(x, z.Star2(y))
	a := y.Quad()
	a.Inv(a)
	z.l.Mul(&z.l, a)
	z.r.Mul(&z.r, a)
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

// Generate returns a random Float64 value for quick.Check testing.
func (z *Float64) Generate(rand *rand.Rand, size int) reflect.Value {
	randomFloat64 := &Float64{
		*nplex.NewFloat64(
			rand.Float64(),
			rand.Float64(),
		),
		*nplex.NewFloat64(
			rand.Float64(),
			rand.Float64(),
		),
	}
	return reflect.ValueOf(randomFloat64)
}
