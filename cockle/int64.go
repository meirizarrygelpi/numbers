// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package cockle

import (
	"math/rand"
	"reflect"
	"strings"

	"fmt"

	"github.com/meirizarrygelpi/numbers/cplex"
	"github.com/meirizarrygelpi/numbers/vec3"
)

// An Int64 is a Cockle quaternion with int64 components.
type Int64 struct {
	l, r cplex.Int64
}

// One sets z equal to 1, and then returns z.
func (z *Int64) One() *Int64 {
	z.l.One()
	z.r.Set(new(cplex.Int64))
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
// If z corresponds to a+bi+ct+du, then the string is "⦗a+bi+ct+du⦘", similar
// to complex128 values.
func (z *Int64) String() string {
	v := z.Unreal()
	a := make([]string, 9)
	a[0] = leftBracket
	a[1] = fmt.Sprint(z.l.Real())
	if v[0] < 0 {
		a[2] = fmt.Sprint(v[0])
	} else {
		a[2] = "+" + fmt.Sprint(v[0])
	}
	a[3] = unit1
	if v[1] < 0 {
		a[4] = fmt.Sprint(v[1])
	} else {
		a[4] = "+" + fmt.Sprint(v[1])
	}
	a[5] = unit2
	if v[2] < 0 {
		a[6] = fmt.Sprint(v[2])
	} else {
		a[6] = "+" + fmt.Sprint(v[2])
	}
	a[7] = unit3
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

// SetPair sets z equal to a Cockle quaternion made with a given pair, and
// then it returns z.
func (z *Int64) SetPair(a, b *cplex.Int64) *Int64 {
	z.l.Set(a)
	z.r.Set(b)
	return z
}

// NewInt64 returns a pointer to the Int64 value a+bi+ct+du.
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
// The multiplication rule is:
// 		Mul(i, i) = Mul(j, j) = Mul(k, k) = -1
// 		Mul(i, j) = -Mul(j, i) = k
// 		Mul(j, k) = -Mul(k, j) = i
// 		Mul(k, i) = -Mul(i, k) = j
// This binary operation is non-commutative but associative.
func (z *Int64) Mul(x, y *Int64) *Int64 {
	a, b, temp := new(cplex.Int64), new(cplex.Int64), new(cplex.Int64)
	a.Add(
		a.Mul(&x.l, &y.l),
		temp.Mul(temp.Conj(&y.r), &x.r),
	)
	b.Add(
		b.Mul(&y.r, &x.l),
		temp.Mul(&x.r, temp.Conj(&y.l)),
	)
	z.SetPair(a, b)
	return z
}

// Commutator sets z equal to the commutator of x and y:
// 		Mul(x, y) - Mul(y, x)
// Then it returns z.
func (z *Int64) Commutator(x, y *Int64) *Int64 {
	return z.Sub(
		z.Mul(x, y),
		new(Int64).Mul(y, x),
	)
}

// Quad returns the quadrance of z. If z = a+bi+ct+du, then the quadrance is
// 		a² + b² - c² - d²
// This can be positive, negative, or zero.
func (z *Int64) Quad() int64 {
	return z.l.Quad() - z.r.Quad()
}

// IsZeroDivisor returns true if z is a zero divisor.
func (z *Int64) IsZeroDivisor() bool {
	return z.l.Quad() == z.r.Quad()
}

// QuoL sets z equal to the left quotient of x and y:
// 		Mul(Inv(y), x)
// Then it returns z. If y is zero, then QuoL panics.
func (z *Int64) QuoL(x, y *Int64) *Int64 {
	if y.IsZeroDivisor() {
		panic(zeroDivisorDenominator)
	}
	return z.Divide(z.Mul(z.Conj(y), x), y.Quad())
}

// QuoR sets z equal to the right quotient of x and y:
// 		Mul(x, Inv(y))
// Then it returns z. If y is zero, then QuoR panics.
func (z *Int64) QuoR(x, y *Int64) *Int64 {
	if y.IsZeroDivisor() {
		panic(zeroDivisorDenominator)
	}
	return z.Divide(z.Mul(x, z.Conj(y)), y.Quad())
}

// Generate returns a random Int64 value for quick.Check testing.
func (z *Int64) Generate(rand *rand.Rand, size int) reflect.Value {
	randomInt64 := &Int64{
		*cplex.NewInt64(
			rand.Int63(),
			rand.Int63(),
		),
		*cplex.NewInt64(
			rand.Int63(),
			rand.Int63(),
		),
	}
	return reflect.ValueOf(randomInt64)
}
