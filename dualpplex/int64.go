// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package dualpplex

import (
	"math/rand"
	"reflect"
	"strings"

	"fmt"

	"github.com/meirizarrygelpi/numbers/pplex"
	"github.com/meirizarrygelpi/numbers/vec3"
)

// An Int64 is a dual-perplex number with int64 components.
type Int64 struct {
	l, r pplex.Int64
}

// One sets z equal to 1, and then returns z.
func (z *Int64) One() *Int64 {
	z.l.One()
	z.r.Set(new(pplex.Int64))
	return z
}

// Real returns the real part of z.
func (z *Int64) Real() int64 {
	return z.l.Real()
}

// Unreal returns the unreal part of z.
func (z *Int64) Unreal() *vec3.Int64 {
	v := new(vec3.Int64)
	v[0] = z.l.Unreal()
	v[1] = z.r.Real()
	v[2] = z.r.Unreal()
	return v
}

// String returns the string version of a Int64 value.
//
// If z corresponds to a+bs+cΓ+dsΓ, then the string is "⦗a+bs+cΓ+dsΓ⦘", similar
// to complex128 values.
func (z *Int64) String() string {
	v := z.Unreal()
	a := make([]string, 9)
	a[0] = leftBracket
	a[1] = fmt.Sprint(z.l.Real())
	i := 2
	for j, u := range [3]string{unit1, unit2, unit3} {
		if v[j] < 0 {
			a[i] = fmt.Sprint(v[j])
		} else {
			a[i] = "+" + fmt.Sprint(v[j])
		}
		a[i+1] = u
		i += 2
	}
	a[8] = rightBracket
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Int64) Equals(y *Int64) bool {
	return z.l.Equals(&y.l) && z.r.Equals(&y.r)
}

// Set sets z equal to y, and returns z.
func (z *Int64) Set(y *Int64) *Int64 {
	z.l.Set(&y.l)
	z.r.Set(&y.r)
	return z
}

// SetPair sets z equal to a dual-perplex number made with a given pair, and
// then it returns z.
func (z *Int64) SetPair(a, b *pplex.Int64) *Int64 {
	z.l.Set(a)
	z.r.Set(b)
	return z
}

// NewInt64 returns a pointer to the Int64 value a+bs+cΓ+dsΓ.
func NewInt64(a, b, c, d int64) *Int64 {
	z := new(Int64)
	z.l.SetPair(a, b)
	z.r.SetPair(c, d)
	return z
}

// Dilate sets z equal to y dilated by a, and returns z.
func (z *Int64) Dilate(y *Int64, a int64) *Int64 {
	z.l.Dilate(&y.l, a)
	z.r.Dilate(&y.r, a)
	return z
}

// Divide sets z equal to y contracted by a, and returns z.
func (z *Int64) Divide(y *Int64, a int64) *Int64 {
	z.l.Divide(&y.l, a)
	z.r.Divide(&y.r, a)
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Int64) Neg(y *Int64) *Int64 {
	z.l.Neg(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Int64) Conj(y *Int64) *Int64 {
	z.l.Conj(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Star1 sets z equal to the s-conjugate of y, and returns z.
func (z *Int64) Star1(y *Int64) *Int64 {
	z.l.Conj(&y.l)
	z.r.Conj(&y.r)
	return z
}

// Star2 sets z equal to the Γ-conjugate of y, and returns z.
func (z *Int64) Star2(y *Int64) *Int64 {
	z.l.Set(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Add sets z equal to x+y, and returns z.
func (z *Int64) Add(x, y *Int64) *Int64 {
	z.l.Add(&x.l, &y.l)
	z.r.Add(&x.r, &y.r)
	return z
}

// Sub sets z equal to x-y, and returns z.
func (z *Int64) Sub(x, y *Int64) *Int64 {
	z.l.Sub(&x.l, &y.l)
	z.r.Sub(&x.r, &y.r)
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication table is:
//     +-----+----+----+----+
//     | Mul | s  | Γ  | sΓ |
//     +-----+----+----+----+
//     | s   | +1 | sΓ | +Γ |
//     +-----+----+----+----+
//     | Γ   | sΓ | 0  | 0  |
//     +-----+----+----+----+
//     | sΓ  | +Γ | 0  | 0  |
//     +-----+----+----+----+
// This binary operation is commutative and associative.
func (z *Int64) Mul(x, y *Int64) *Int64 {
	a, b, temp := new(pplex.Int64), new(pplex.Int64), new(pplex.Int64)
	a.Mul(&x.l, &y.l)
	b.Add(
		b.Mul(&x.l, &y.r),
		temp.Mul(&x.r, &y.l),
	)
	z.SetPair(a, b)
	return z
}

// Quad returns the quadrance of z. If z = a+bs+cΓ+dsΓ, then the quadrance is
//     a² + b² + 2abs
// Note that this is a perplex number.
func (z *Int64) Quad() *pplex.Int64 {
	q := new(pplex.Int64)
	return q.Mul(&z.l, &z.l)
}

// Norm returns the norm of z. If z = a+bs+cΓ+dsΓ, then the norm is
// 		(a² - b²)²
// This is always non-negative.
func (z *Int64) Norm() int64 {
	return z.Quad().Quad()
}

// IsZeroDivisor returns true if z is a zero divisor.
func (z *Int64) IsZeroDivisor() bool {
	return z.Quad().IsZeroDivisor()
}

// Quo sets z equal to the quotient of x and y, and returns z. If y is zero,
// then Quo panics.
func (z *Int64) Quo(x, y *Int64) *Int64 {
	if y.IsZeroDivisor() {
		panic(zeroDivisorDenominator)
	}
	n := y.Norm()
	temp := new(Int64)
	z.Mul(x, temp.Star1(y))
	z.Mul(z, temp.Star2(y))
	z.Mul(z, temp.Star1(temp))
	return z.Divide(z, n)
}

// Generate returns a random Int64 value for quick.Check testing.
func (z *Int64) Generate(rand *rand.Rand, size int) reflect.Value {
	randomInt64 := &Int64{
		*pplex.NewInt64(
			rand.Int63(),
			rand.Int63(),
		),
		*pplex.NewInt64(
			rand.Int63(),
			rand.Int63(),
		),
	}
	return reflect.ValueOf(randomInt64)
}
