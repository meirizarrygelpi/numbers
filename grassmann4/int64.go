// Copyright (c) 2016-2017-2017 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package grassmann4

import (
	"math/rand"
	"reflect"
	"strings"

	"fmt"

	"github.com/meirizarrygelpi/numbers/grassmann3"
)

// An Int64 is a four-dimensional Grassmann number with int64 components.
type Int64 struct {
	l, r grassmann3.Int64
}

// One sets z equal to 1, and then returns z.
func (z *Int64) One() *Int64 {
	z.l.One()
	z.r.Set(new(grassmann3.Int64))
	return z
}

// Real returns the real part of z.
func (z *Int64) Real() int64 {
	return z.l.Real()
}

// Unreal returns the unreal part of z.
func (z *Int64) Unreal() *[15]int64 {
	var v [15]int64
	w := z.l.Unreal()
	v[0] = w[0]
	v[1] = w[1]
	v[2] = w[2]
	v[3] = w[3]
	v[4] = w[4]
	v[5] = w[5]
	v[6] = w[6]
	v[7] = z.r.Real()
	w = z.r.Unreal()
	v[8] = w[0]
	v[9] = w[1]
	v[10] = w[2]
	v[11] = w[3]
	v[12] = w[4]
	v[13] = w[5]
	v[14] = w[6]
	return &v
}

// String returns the string version of a Int64 value.
func (z *Int64) String() string {
	v := z.Unreal()
	a := make([]string, 33)
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
	a[32] = rightBracket
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

// SetPair sets z equal to a four-dimensional Grassmann number made with a given pair, and
// then it returns z.
func (z *Int64) SetPair(a, b *grassmann3.Int64) *Int64 {
	z.l.Set(a)
	z.r.Set(b)
	return z
}

// Set0Blade sets the 0-blade of z equal to a0, and then it returns z.
func (z *Int64) Set0Blade(a0 int64) *Int64 {
	z.l.Set0Blade(a0)
	return z
}

// Set1Blades sets the 1-blades of z equal to aW, aX, aY, and aZ, and then it
// returns z.
func (z *Int64) Set1Blades(aW, aX, aY, aZ int64) *Int64 {
	z.l.Set1Blades(aW, aX, aY)
	z.r.Set0Blade(aZ)
	return z
}

// Set2Blades sets the 2-blades of z equal to aWX, aWY, aXY, aWZ, aXZ, and aYZ,
// then it returns z.
func (z *Int64) Set2Blades(aWX, aWY, aXY, aWZ, aXZ, aYZ int64) *Int64 {
	z.l.Set2Blades(aWX, aWY, aXY)
	z.r.Set1Blades(aWZ, aXZ, aYZ)
	return z
}

// Set3Blades sets the 3-blade of z equal to aWXY, and then it returns z.
func (z *Int64) Set3Blades(aWXY, aWXZ, aWYZ, aXYZ int64) *Int64 {
	z.l.Set3Blade(aWXY)
	z.r.Set2Blades(aWXZ, aWYZ, aXYZ)
	return z
}

// Set4Blade sets the 4-blade of z equal to aWXY, and then it returns z.
func (z *Int64) Set4Blade(aWXYZ int64) *Int64 {
	z.r.Set3Blade(aWXYZ)
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

// Dagger sets z equal to the dagger conjugate of y, and returns z.
func (z *Int64) Dagger(y *Int64) *Int64 {
	z.l.Dagger(&y.l)
	z.r.Dagger(&y.r)
	z.r.Neg(&z.r)
	return z
}

// Hodge sets z equal to the Hodge conjugate of y, and returns z.
func (z *Int64) Hodge(y *Int64) *Int64 {
	a, b := new(grassmann3.Int64), new(grassmann3.Int64)
	a.Set(&y.l)
	b.Set(&y.r)
	return z.SetPair(b.Dagger(b.Hodge(b)), a.Hodge(a))
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
	a, b, temp := new(grassmann3.Int64), new(grassmann3.Int64), new(grassmann3.Int64)
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
func (z *Int64) Commutator(x, y *Int64) *Int64 {
	return z.Sub(
		z.Mul(x, y),
		new(Int64).Mul(y, x),
	)
}

// Associator sets z equal to the associator of w, x, and y:
// 		Mul(Mul(w, x), y) - Mul(w, Mul(x, y))
// Then it returns z.
func (z *Int64) Associator(w, x, y *Int64) *Int64 {
	temp := new(Int64)
	return z.Sub(
		z.Mul(z.Mul(w, x), y),
		temp.Mul(w, temp.Mul(x, y)),
	)
}

// Quad returns the quadrance of z. If z = a+..., then
// the quadrance is
// 		a²
// This is always non-negative.
func (z *Int64) Quad() int64 {
	return z.l.Quad()
}

// IsZeroDivisor returns true if z is a zero divisor.
func (z *Int64) IsZeroDivisor() bool {
	return z.l.IsZeroDivisor()
}

// QuoL sets z equal to the left quotient of x and y:
// 		Mul(Inv(y), x)
// Then it returns z. If y is a zero divisor, then QuoL panics.
func (z *Int64) QuoL(x, y *Int64) *Int64 {
	if y.IsZeroDivisor() {
		panic(zeroDivisorDenominator)
	}
	return z.Divide(z.Mul(z.Conj(y), x), y.Quad())
}

// QuoR sets z equal to the right quotient of x and y:
// 		Mul(x, Inv(y))
// Then it returns z. If y is a zero divisor, then QuoR panics.
func (z *Int64) QuoR(x, y *Int64) *Int64 {
	if y.IsZeroDivisor() {
		panic(zeroDivisorDenominator)
	}
	return z.Divide(z.Mul(x, z.Conj(y)), y.Quad())
}

// Generate returns a random Int64 value for quick.Check testing.
func (z *Int64) Generate(rand *rand.Rand, size int) reflect.Value {
	randomInt64 := &Int64{
		*grassmann3.NewInt64(
			rand.Int63(),
			rand.Int63(),
			rand.Int63(),
			rand.Int63(),
			rand.Int63(),
			rand.Int63(),
			rand.Int63(),
			rand.Int63(),
		),
		*grassmann3.NewInt64(
			rand.Int63(),
			rand.Int63(),
			rand.Int63(),
			rand.Int63(),
			rand.Int63(),
			rand.Int63(),
			rand.Int63(),
			rand.Int63(),
		),
	}
	return reflect.ValueOf(randomInt64)
}
