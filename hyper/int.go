// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package hyper

import (
	"math/big"
	"math/rand"
	"reflect"
	"strings"

	"github.com/meirizarrygelpi/numbers/nplex"
	"github.com/meirizarrygelpi/numbers/vec3"
)

// An Int is a hyper number with big.Int components.
type Int struct {
	l, r nplex.Int
}

// One sets z equal to 1, and then returns z.
func (z *Int) One() *Int {
	z.l.One()
	z.r.Set(new(nplex.Int))
	return z
}

// Real returns the real part of z.
func (z *Int) Real() *big.Int {
	return z.l.Real()
}

// Unreal returns the unreal part of z, a three-dimensional vector.
func (z *Int) Unreal() *vec3.Int {
	v := new(vec3.Int)
	v[0] = z.l.Unreal()
	v[1] = z.r.Real()
	v[2] = z.r.Unreal()
	return v
}

// String returns the string version of an Int value.
//
// If z corresponds to a+bα+cΓ+dαΓ, then the string is "⦗a+bα+cΓ+dαΓ⦘", similar
// to complex128 values.
func (z *Int) String() string {
	v := z.Unreal()
	a := make([]string, 9)
	a[0] = leftBracket
	a[1] = z.l.Real().String()
	if v[0].Sign() < 0 {
		a[2] = v[0].String()
	} else {
		a[2] = "+" + v[0].String()
	}
	a[3] = unit1
	if v[1].Sign() < 0 {
		a[4] = v[1].String()
	} else {
		a[4] = "+" + v[1].String()
	}
	a[5] = unit2
	if v[2].Sign() < 0 {
		a[6] = v[2].String()
	} else {
		a[6] = "+" + v[2].String()
	}
	a[7] = unit3
	a[8] = rightBracket
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Int) Equals(y *Int) bool {
	return z.l.Equals(&y.l) && z.r.Equals(&y.r)
}

// Set sets z equal to y, and returns z.
func (z *Int) Set(y *Int) *Int {
	z.l.Set(&y.l)
	z.r.Set(&y.r)
	return z
}

// SetPair sets z equal to a hyper number made with a given pair, and
// then it returns z.
func (z *Int) SetPair(a, b *nplex.Int) *Int {
	z.l.Set(a)
	z.r.Set(b)
	return z
}

// NewInt returns a pointer to the Int value a+bα+cΓ+dαΓ.
func NewInt(a, b, c, d *big.Int) *Int {
	z := new(Int)
	z.l.SetPair(a, b)
	z.r.SetPair(c, d)
	return z
}

// Dilate sets z equal to y dilated by a, and returns z.
func (z *Int) Dilate(y *Int, a *big.Int) *Int {
	z.l.Dilate(&y.l, a)
	z.r.Dilate(&y.r, a)
	return z
}

// Divide sets z equal to y contracted by a, and returns z.
func (z *Int) Divide(y *Int, a *big.Int) *Int {
	z.l.Divide(&y.l, a)
	z.r.Divide(&y.r, a)
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Int) Neg(y *Int) *Int {
	z.l.Neg(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Star1 sets z equal to the α-conjugate of y, and returns z.
func (z *Int) Star1(y *Int) *Int {
	z.l.Conj(&y.l)
	z.r.Conj(&y.r)
	return z
}

// Star2 sets z equal to the Γ-conjugate of y, and returns z.
func (z *Int) Star2(y *Int) *Int {
	z.l.Set(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Add sets z equal to x+y, and returns z.
func (z *Int) Add(x, y *Int) *Int {
	z.l.Add(&x.l, &y.l)
	z.r.Add(&x.r, &y.r)
	return z
}

// Sub sets z equal to x-y, and returns z.
func (z *Int) Sub(x, y *Int) *Int {
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
func (z *Int) Mul(x, y *Int) *Int {
	a, b, temp := new(nplex.Int), new(nplex.Int), new(nplex.Int)
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
func (z *Int) Quad() *nplex.Int {
	q := new(nplex.Int)
	return q.Mul(&z.l, &z.l)
}

// Norm returns the norm of z. If z = a+bα+cΓ+dαΓ, then the norm is
// 		(a²)²
// This is always non-negative.
func (z *Int) Norm() *big.Int {
	return z.Quad().Quad()
}

// IsZeroDivisor returns true if z is a zero divisor.
func (z *Int) IsZeroDivisor() bool {
	return z.Quad().IsZeroDivisor()
}

// Quo sets z equal to the quotient of x and y, and returns z. If y is zero,
// then Quo panics.
func (z *Int) Quo(x, y *Int) *Int {
	if y.IsZeroDivisor() {
		panic(zeroDivisorDenominator)
	}
	z.Mul(x, z.Star2(y))
	a := y.Quad()
	z.l.Quo(&z.l, a)
	z.r.Quo(&z.r, a)
	return z
}

// Generate returns a random Int value for quick.Check testing.
func (z *Int) Generate(rand *rand.Rand, size int) reflect.Value {
	randomInt := &Int{
		*nplex.NewInt(
			big.NewInt(rand.Int63()),
			big.NewInt(rand.Int63()),
		),
		*nplex.NewInt(
			big.NewInt(rand.Int63()),
			big.NewInt(rand.Int63()),
		),
	}
	return reflect.ValueOf(randomInt)
}
