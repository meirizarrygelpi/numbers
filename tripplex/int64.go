// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licensed under the MIT License.

package tripplex

import (
	"math/rand"
	"reflect"
	"strings"

	"fmt"

	"github.com/meirizarrygelpi/numbers/bipplex"
	"github.com/meirizarrygelpi/numbers/pplex"
	"github.com/meirizarrygelpi/numbers/vec7"
)

// An Int64 is a tri-perplex number with int64 components.
type Int64 struct {
	l, r bipplex.Int64
}

// One sets z equal to 1, and then returns z.
func (z *Int64) One() *Int64 {
	z.l.One()
	z.r.Set(new(bipplex.Int64))
	return z
}

// Real returns the real part of z.
func (z *Int64) Real() int64 {
	return z.l.Real()
}

// Unreal returns the unreal part of z.
func (z *Int64) Unreal() *vec7.Int64 {
	v := new(vec7.Int64)
	w := z.l.Unreal()
	v[0] = w[0]
	v[1] = w[1]
	v[2] = w[2]
	v[3] = z.r.Real()
	w = z.r.Unreal()
	v[4] = w[0]
	v[5] = w[1]
	v[6] = w[2]
	return v
}

// String returns the string version of a Int64 value.
//
// If z corresponds to a+bs+cT+dsT+eU+fsU+gTU+hsTU, then the string is
// "(a+bs+cT+dsT+eU+fsU+gTU+hsTU)", similar to complex128 values.
func (z *Int64) String() string {
	v := z.Unreal()
	a := make([]string, 17)
	a[0] = leftBracket
	a[1] = fmt.Sprint(z.l.Real())
	i := 2
	for j, u := range unitNames {
		if v[j] < 0 {
			a[i] = fmt.Sprint(v[j])
		} else {
			a[i] = "+" + fmt.Sprint(v[j])
		}
		a[i+1] = u
		i += 2
	}
	a[16] = rightBracket
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

// SetPair sets z equal to a tri-perplex number made with a given pair, and
// then it returns z.
func (z *Int64) SetPair(a, b *bipplex.Int64) *Int64 {
	z.l.Set(a)
	z.r.Set(b)
	return z
}

// NewInt64 returns a pointer to the Int64 value a+bs+cT+dsT+eU+fsU+gTU+hsTU.
func NewInt64(a, b, c, d, e, f, g, h int64) *Int64 {
	z := new(Int64)
	z.l.SetPair(
		pplex.NewInt64(a, b),
		pplex.NewInt64(c, d),
	)
	z.r.SetPair(
		pplex.NewInt64(e, f),
		pplex.NewInt64(g, h),
	)
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

// Bar sets z equal to the i-conjugate of y, and returns z.
func (z *Int64) Bar(y *Int64) *Int64 {
	z.l.Bar(&y.l)
	z.r.Bar(&y.r)
	return z
}

// Tilde sets z equal to the J-conjugate of y, and returns z.
func (z *Int64) Tilde(y *Int64) *Int64 {
	z.l.Tilde(&y.l)
	z.r.Tilde(&y.r)
	return z
}

// Star sets z equal to the K-conjugate of y, and returns z.
func (z *Int64) Star(y *Int64) *Int64 {
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
func (z *Int64) Mul(x, y *Int64) *Int64 {
	a, b, temp := new(bipplex.Int64), new(bipplex.Int64), new(bipplex.Int64)
	a.Add(
		a.Mul(&x.l, &y.l),
		temp.Mul(&x.r, &y.r),
	)
	b.Add(
		b.Mul(&x.l, &y.r),
		temp.Mul(&x.r, &y.l),
	)
	z.SetPair(a, b)
	return z
}

// Quad returns the quadrance of z. If z = a+bs+cT+dsT+eU+fsU+gTU+hsTU, then
// the quadrance is
// 		a² - b² + c² - d² + 2(ab + cd)i
// Note that this is a perplex number.
func (z *Int64) Quad() *bipplex.Int64 {
	q := new(bipplex.Int64)
	return q.Sub(q.Mul(&z.l, &z.l), new(bipplex.Int64).Mul(&z.r, &z.r))
}

// Norm returns the norm of z. If z = a+bs+cT+dsT+eU+fsU+gTU+hsTU, then the
// norm is
// 		(a² - b² + c² - d²)² + 4(ab + cd)²
// There is another way to write the norm as a sum of two squares:
// 		(a² + b² - c² - d²)² + 4(ac + bd)²
// Alternatively, it can also be written as a difference of two squares:
//		(a² + b² + c² + d²)² - 4(ad - bc)²
// Finally, you have the factorized form:
// 		((a - d)² + (b + c)²)((a + d)² + (b - c)²)
// In this form it is clear that the norm is always non-negative.
func (z *Int64) Norm() int64 {
	return z.Quad().Norm()
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
	a := y.Quad()
	z.Mul(x, z.Star(y))
	z.l.Quo(&z.l, a)
	z.r.Quo(&z.r, a)
	return z
}

// Generate returns a random Int64 value for quick.Check testing.
func (z *Int64) Generate(rand *rand.Rand, size int) reflect.Value {
	randomInt64 := &Int64{
		*bipplex.NewInt64(
			rand.Int63(),
			rand.Int63(),
			rand.Int63(),
			rand.Int63(),
		),
		*bipplex.NewInt64(
			rand.Int63(),
			rand.Int63(),
			rand.Int63(),
			rand.Int63(),
		),
	}
	return reflect.ValueOf(randomInt64)
}
