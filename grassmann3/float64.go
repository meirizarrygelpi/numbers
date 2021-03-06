// Copyright (c) 2016-2017 Melvin Eloy Irizarry-Gelpí
// Licensed under the MIT License.

package grassmann3

import (
	"math/rand"
	"reflect"
	"strings"

	"fmt"

	"math"

	"github.com/meirizarrygelpi/numbers/nplex"
	"github.com/meirizarrygelpi/numbers/grassmann2"
	"github.com/meirizarrygelpi/numbers/vec7"
)

// A Float64 is a three-dimensional Grassmann number with float64 components.
type Float64 struct {
	l, r grassmann2.Float64
}

// One sets z equal to 1, and then returns z.
func (z *Float64) One() *Float64 {
	z.l.One()
	z.r.Set(new(grassmann2.Float64))
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
// If z corresponds to a+bW+cX+dWX+eY+fWY+gXY+h(WX)Y, then the string is
// "(a+bW+cX+dWX+eY+fWY+gXY+h(WX)Y)", similar to complex128 values.
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

// SetPair sets z equal to a three-dimensional Grassmann number made with a given pair, and
// then it returns z.
func (z *Float64) SetPair(a, b *grassmann2.Float64) *Float64 {
	z.l.Set(a)
	z.r.Set(b)
	return z
}

// Set0Blade sets the 0-blade of z equal to a0, and then it returns z.
func (z *Float64) Set0Blade(a0 float64) *Float64 {
	z.l.Set0Blade(a0)
	return z
}

// Set1Blades sets the 1-blades of z equal to aW, aX, and aY, and then it returns
// z.
func (z *Float64) Set1Blades(aW, aX, aY float64) *Float64 {
	z.l.Set1Blades(aW, aX)
	z.r.Set0Blade(aY)
	return z
}

// Set2Blades sets the 2-blades of z equal to aWX, aWY, and aXY, and then it
// returns z.
func (z *Float64) Set2Blades(aWX, aWY, aXY float64) *Float64 {
	z.l.Set2Blade(aWX)
	z.r.Set1Blades(aWY, aXY)
	return z
}

// Set3Blade sets the 0-blade of z equal to aWXY, and then it returns z.
func (z *Float64) Set3Blade(aWXY float64) *Float64 {
	z.r.Set2Blade(aWXY)
	return z
}

// NewFloat64 returns a pointer to the Float64 value
// a+bW+cX+dWX+eY+fWY+gXY+h(WX)Y.
func NewFloat64(a, b, c, d, e, f, g, h float64) *Float64 {
	z := new(Float64)
	z.l.SetPair(
		nplex.NewFloat64(a, b),
		nplex.NewFloat64(c, d),
	)
	z.r.SetPair(
		nplex.NewFloat64(e, f),
		nplex.NewFloat64(g, h),
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

// Dagger sets z equal to the dagger conjugate of y, and returns z.
func (z *Float64) Dagger(y *Float64) *Float64 {
	z.l.Dagger(&y.l)
	z.r.Dagger(&y.r)
	z.r.Neg(&z.r)
	return z
}

// Hodge sets z equal to the Hodge conjugate of y, and returns z.
func (z *Float64) Hodge(y *Float64) *Float64 {
	a, b := new(grassmann2.Float64), new(grassmann2.Float64)
	a.Set(&y.l)
	b.Set(&y.r)
	return z.SetPair(b.Dagger(b.Hodge(b)), a.Hodge(a))
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
	a, b, temp := new(grassmann2.Float64), new(grassmann2.Float64), new(grassmann2.Float64)
	a.Mul(&x.l, &y.l)
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
func (z *Float64) Commutator(x, y *Float64) *Float64 {
	return z.Sub(
		z.Mul(x, y),
		new(Float64).Mul(y, x),
	)
}

// Associator sets z equal to the associator of w, x, and y:
// 		Mul(Mul(w, x), y) - Mul(w, Mul(x, y))
// Then it returns z.
func (z *Float64) Associator(w, x, y *Float64) *Float64 {
	temp := new(Float64)
	return z.Sub(
		z.Mul(z.Mul(w, x), y),
		temp.Mul(w, temp.Mul(x, y)),
	)
}

// Quad returns the quadrance of z. If z = a+bW+cX+dWX+eY+fWY+gXY+h(WX)Y, then
// the quadrance is
// 		a²
// This is always non-negative.
func (z *Float64) Quad() float64 {
	return z.l.Quad()
}

// IsZeroDivisor returns true if z is a zero divisor.
func (z *Float64) IsZeroDivisor() bool {
	return z.l.IsZeroDivisor()
}

// Inv sets z equal to the inverse of y, and returns z. If y is a zero divisor,
// then Inv panics.
func (z *Float64) Inv(y *Float64) *Float64 {
	if y.IsZeroDivisor() {
		panic(zeroDivisorInverse)
	}
	return z.Divide(z.Conj(y), y.Quad())
}

// QuoL sets z equal to the left quotient of x and y:
// 		Mul(Inv(y), x)
// Then it returns z. If y is a zero divisor, then QuoL panics.
func (z *Float64) QuoL(x, y *Float64) *Float64 {
	if y.IsZeroDivisor() {
		panic(zeroDivisorDenominator)
	}
	return z.Mul(z.Inv(y), x)
}

// QuoR sets z equal to the right quotient of x and y:
// 		Mul(x, Inv(y))
// Then it returns z. If y is a zero divisor, then QuoR panics.
func (z *Float64) QuoR(x, y *Float64) *Float64 {
	if y.IsZeroDivisor() {
		panic(zeroDivisorDenominator)
	}
	return z.Mul(x, z.Inv(y))
}

// SelfDual sets z equal to the self-dual part of y. If z is self-dual, then
//     Hodge(z) = z
// Then it returns z.
func (z *Float64) SelfDual(y *Float64) *Float64 {
	sd := new(Float64)
	sd.Hodge(y)
	sd.Add(y, sd)
	sd.Divide(sd, 2)
	return z.Set(sd)
}

// AntiSelfDual sets z equal to the anti-self-dual part of y. If z is
// anti-self-dual, then
//     Hodge(z) = -z
// Then it returns z.
func (z *Float64) AntiSelfDual(y *Float64) *Float64 {
	asd := new(Float64)
	asd.Hodge(y)
	asd.Sub(y, asd)
	asd.Divide(asd, 2)
	return z.Set(asd)
}

// Generate returns a random Float64 value for quick.Check testing.
func (z *Float64) Generate(rand *rand.Rand, size int) reflect.Value {
	randomFloat64 := &Float64{
		*grassmann2.NewFloat64(
			rand.Float64(),
			rand.Float64(),
			rand.Float64(),
			rand.Float64(),
		),
		*grassmann2.NewFloat64(
			rand.Float64(),
			rand.Float64(),
			rand.Float64(),
			rand.Float64(),
		),
	}
	return reflect.ValueOf(randomFloat64)
}
