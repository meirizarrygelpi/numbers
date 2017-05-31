// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package nplex

import (
	"math/big"
	"math/rand"
	"reflect"
	"strings"
)

// An Int is a nilplex number with big.Int components.
type Int struct {
	l, r big.Int
}

// One sets z equal to 1, and then it returns z.
func (z *Int) One() *Int {
	z.l.SetInt64(1)
	z.r.SetInt64(0)
	return z
}

// Real returns the real part of z.
func (z *Int) Real() *big.Int {
	return &z.l
}

// Unreal returns the unreal part of z.
func (z *Int) Unreal() *big.Int {
	return &z.r
}

// String returns the string version of an Int value.
//
// If z corresponds to a + bα, then the string is "⦗a+bα⦘", similar to
// complex128 values.
func (z *Int) String() string {
	a := make([]string, 5)
	a[0] = leftBracket
	a[1] = z.l.String()
	if z.r.Sign() < 0 {
		a[2] = z.r.String()
	} else {
		a[2] = "+" + z.r.String()
	}
	a[3] = unitName
	a[4] = rightBracket
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Int) Equals(y *Int) bool {
	return z.l.Cmp(&y.l) == 0 && z.r.Cmp(&y.r) == 0
}

// Set sets z equal to y, and returns z.
func (z *Int) Set(y *Int) *Int {
	z.l.Set(&y.l)
	z.r.Set(&y.r)
	return z
}

// SetPair sets z equal to a nilplex number made with a given pair, and then
// it returns z.
func (z *Int) SetPair(a, b *big.Int) *Int {
	z.l.Set(a)
	z.r.Set(b)
	return z
}

// NewInt returns a pointer to the Int value a+bα.
func NewInt(a, b *big.Int) *Int {
	z := new(Int)
	z.SetPair(a, b)
	return z
}

// Dilate sets z equal to y dilated by a, and returns z.
func (z *Int) Dilate(y *Int, a *big.Int) *Int {
	z.l.Mul(&y.l, a)
	z.r.Mul(&y.r, a)
	return z
}

// Divide sets z equal to y contracted by a, and returns z.
func (z *Int) Divide(y *Int, a *big.Int) *Int {
	z.l.Quo(&y.l, a)
	z.r.Quo(&y.r, a)
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Int) Neg(y *Int) *Int {
	z.l.Neg(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Int) Conj(y *Int) *Int {
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
func (z *Int) Mul(x, y *Int) *Int {
	a, b, temp := new(big.Int), new(big.Int), new(big.Int)
	a.Mul(&x.l, &y.l)
	b.Add(
		temp.Mul(&x.r, &y.l),
		b.Mul(&y.r, &x.l),
	)
	z.SetPair(a, b)
	return z
}

// Dot returns the dot product of y and z. If z = a+bα and y = c+dα, then the
// dot product is
// 		ac
// This can be positive, negative, or zero. The dot product is equivalent to
// 		½(Mul(Conj(z), y) + Mu(Conj(y), z))
// In this form it is clear that Dot is symmetric.
func (z *Int) Dot(y *Int) *big.Int {
	dot := new(big.Int)
	return dot.Mul(&z.l, &y.l)
}

// Quad returns the quadrance of z. If z = a+bα, then the quadrance is
// 		a²
// This is always non-negative.
func (z *Int) Quad() *big.Int {
	return z.Dot(z)
}

// Cross returns the cross product of y and z. If z = a+bα and y = c+dα, then
// the cross product is
// 		ad - bc
// This can be positive, negative, or zero. The cross product is equivalent to
// the unreal part of
// 		½(Mul(Conj(z), y) - Mu(Conj(y), z))
// In this form it is clear that Cross is anti-symmetric.
func (z *Int) Cross(y *Int) *big.Int {
	cross := new(big.Int)
	return cross.Sub(
		cross.Mul(&z.l, &y.r),
		new(big.Int).Mul(&z.r, &y.l),
	)
}

// IsZeroDivisor returns true if z is a zero divisor. This is equivalent to z
// having zero quadrance.
func (z *Int) IsZeroDivisor() bool {
	zero := new(big.Int)
	return z.l.Cmp(zero) == 0
}

// Quo sets z equal to the quotient of x and y, and returns z. Note that
// truncated division is used. If y is a zero divisor, then Quo panics.
func (z *Int) Quo(x, y *Int) *Int {
	if y.IsZeroDivisor() {
		panic(zeroDivisorDenominator)
	}
	q := y.Quad()
	a, b, temp := new(big.Int), new(big.Int), new(big.Int)
	a.Mul(&x.l, &y.l)
	b.Sub(
		temp.Mul(&x.r, &y.l),
		b.Mul(&y.r, &x.l),
	)
	z.SetPair(a, b)
	return z.Divide(z, q)
}

// Generate returns a random Int value for quick.Check testing.
func (z *Int) Generate(rand *rand.Rand, size int) reflect.Value {
	randomInt := &Int{
		*big.NewInt(rand.Int63()),
		*big.NewInt(rand.Int63()),
	}
	return reflect.ValueOf(randomInt)
}
