// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package hypercplex

import (
	"math/rand"
	"reflect"
	"strings"

	"fmt"

	"math"

	"github.com/meirizarrygelpi/numbers/cplex"
	"github.com/meirizarrygelpi/numbers/dualcplex"
	"github.com/meirizarrygelpi/numbers/vec7"
)

// A Float64 is a hyper-complex number with float64 components.
type Float64 struct {
	l, r dualcplex.Float64
}

// One sets z equal to 1, and then returns z.
func (z *Float64) One() *Float64 {
	z.l.One()
	z.r.Set(new(dualcplex.Float64))
	return z
}

// Real returns the real part of z.
func (z *Float64) Real() float64 {
	return z.l.Real()
}

// Unreal returns the unreal part of z.
func (z *Float64) Unreal() *vec7.Float64 {
	v := new(vec7.Float64)
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
// If z corresponds to a+bi+cΓ+diΓ+eΛ+fiΛ+gΓΛ+hiΓΛ, then the string is
// "⦗a+bi+cΓ+diΓ+eΛ+fiΛ+gΓΛ+hiΓΛ⦘", similar to nilplex128 values.
func (z *Float64) String() string {
	v := z.Unreal()
	a := make([]string, 17)
	a[0] = leftBracket
	a[1] = fmt.Sprint(z.l.Real())
	i := 2
	for j, u := range unitNames {
		a[i] = sprintFloat64(v[j])
		a[i+1] = u
		i += 2
	}
	a[16] = rightBracket
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

// SetPair sets z equal to a hyper-complex number made with a given pair, and
// then it returns z.
func (z *Float64) SetPair(a, b *dualcplex.Float64) *Float64 {
	z.l.Set(a)
	z.r.Set(b)
	return z
}

// NewFloat64 returns a pointer to the Float64 value
// a+bi+cΓ+diΓ+eΛ+fiΛ+gΓΛ+hiΓΛ.
func NewFloat64(a, b, c, d, e, f, g, h float64) *Float64 {
	z := new(Float64)
	z.l.SetPair(
		cplex.NewFloat64(a, b),
		cplex.NewFloat64(c, d),
	)
	z.r.SetPair(
		cplex.NewFloat64(e, f),
		cplex.NewFloat64(g, h),
	)
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

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Float64) Conj(y *Float64) *Float64 {
	z.l.Conj(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Star1 sets z equal to the i-conjugate of y, and returns z.
func (z *Float64) Star1(y *Float64) *Float64 {
	z.l.Star1(&y.l)
	z.r.Star1(&y.r)
	return z
}

// Star2 sets z equal to the Γ-conjugate of y, and returns z.
func (z *Float64) Star2(y *Float64) *Float64 {
	z.l.Star2(&y.l)
	z.r.Star2(&y.r)
	return z
}

// Star3 sets z equal to the Λ-conjugate of y, and returns z.
func (z *Float64) Star3(y *Float64) *Float64 {
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
func (z *Float64) Mul(x, y *Float64) *Float64 {
	a, b, temp := new(dualcplex.Float64), new(dualcplex.Float64), new(dualcplex.Float64)
	a.Mul(&x.l, &y.l)
	b.Add(
		b.Mul(&x.l, &y.r),
		temp.Mul(&x.r, &y.l),
	)
	z.SetPair(a, b)
	return z
}

// Quad returns the quadrance of z. If z = a+bi+cΓ+diΓ+eΛ+fiΛ+gΓΛ+hiΓΛ, then
// the quadrance is
// 		a² + 2abα + 2acΓ + 2(ad + bc)αΓ
// Note that this is a dualcplex number.
func (z *Float64) Quad() *dualcplex.Float64 {
	q := new(dualcplex.Float64)
	return q.Mul(&z.l, &z.l)
}

// Norm returns the norm of z. If z = a+bi+cΓ+diΓ+eΛ+fiΛ+gΓΛ+hiΓΛ, then the
// norm is
// 		((a²)²)²
// In this form it is clear that the norm is always non-negative.
func (z *Float64) Norm() float64 {
	return z.Quad().Norm()
}

// IsZeroDivisor returns true if z is a zero divisor.
func (z *Float64) IsZeroDivisor() bool {
	return z.Quad().IsZeroDivisor()
}

// Inv sets z equal to the inverse of y, and returns z. If y is a zero divisor,
// then Inv panics.
func (z *Float64) Inv(y *Float64) *Float64 {
	if y.IsZeroDivisor() {
		panic(zeroDivisorInverse)
	}
	a := y.Quad()
	a.Inv(a)
	z.Star3(y)
	z.l.Mul(&z.l, a)
	z.r.Mul(&z.r, a)
	return z
}

// Quo sets z equal to the quotient of x and y, and returns z. If y is a zero
// divisor, then Quo panics.
func (z *Float64) Quo(x, y *Float64) *Float64 {
	if y.IsZeroDivisor() {
		panic(zeroDivisorDenominator)
	}
	a := y.Quad()
	a.Inv(a)
	z.Mul(x, z.Star3(y))
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
		*dualcplex.NewFloat64(
			rand.Float64(),
			rand.Float64(),
			rand.Float64(),
			rand.Float64(),
		),
		*dualcplex.NewFloat64(
			rand.Float64(),
			rand.Float64(),
			rand.Float64(),
			rand.Float64(),
		),
	}
	return reflect.ValueOf(randomFloat64)
}
