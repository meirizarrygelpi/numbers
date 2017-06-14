// Copyright (c) 2016-2017-2017 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package grassmann4

import (
	"math/big"
	"math/rand"
	"reflect"
	"strings"

	"github.com/meirizarrygelpi/numbers/grassmann3"
)

// An Int is a four-dimensional Grassmann number with big.Int components.
type Int struct {
	l, r grassmann3.Int
}

// One sets z equal to 1, and then returns z.
func (z *Int) One() *Int {
	z.l.One()
	z.r.Set(new(grassmann3.Int))
	return z
}

// Real returns the real part of z.
func (z *Int) Real() *big.Int {
	return z.l.Real()
}

// Unreal returns the unreal part of z.
func (z *Int) Unreal() *[15]*big.Int {
	var v [15]*big.Int
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

// String returns the string version of an Int value.
func (z *Int) String() string {
	v := z.Unreal()
	a := make([]string, 33)
	a[0] = leftBracket
	a[1] = z.l.Real().String()
	i := 2
	for j, u := range unitNames {
		if v[j].Sign() < 0 {
			a[i] = v[j].String()
		} else {
			a[i] = "+" + v[j].String()
		}
		a[i+1] = u
		i += 2
	}
	a[32] = rightBracket
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

// SetPair sets z equal to a four-dimensional Grassmann number made with a given pair, and
// then it returns z.
func (z *Int) SetPair(a, b *grassmann3.Int) *Int {
	z.l.Set(a)
	z.r.Set(b)
	return z
}

// Set0Blade sets the 0-blade of z equal to a0, and then it returns z.
func (z *Int) Set0Blade(a0 *big.Int) *Int {
	z.l.Set0Blade(a0)
	return z
}

// Set1Blades sets the 1-blades of z equal to aW, aX, aY, and aZ, and then it
// returns z.
func (z *Int) Set1Blades(aW, aX, aY, aZ *big.Int) *Int {
	z.l.Set1Blades(aW, aX, aY)
	z.r.Set0Blade(aZ)
	return z
}

// Set2Blades sets the 2-blades of z equal to aWX, aWY, aXY, aWZ, aXZ, and aYZ,
// then it returns z.
func (z *Int) Set2Blades(aWX, aWY, aXY, aWZ, aXZ, aYZ *big.Int) *Int {
	z.l.Set2Blades(aWX, aWY, aXY)
	z.r.Set1Blades(aWZ, aXZ, aYZ)
	return z
}

// Set3Blades sets the 3-blade of z equal to aWXY, and then it returns z.
func (z *Int) Set3Blades(aWXY, aWXZ, aWYZ, aXYZ *big.Int) *Int {
	z.l.Set3Blade(aWXY)
	z.r.Set2Blades(aWXZ, aWYZ, aXYZ)
	return z
}

// Set4Blade sets the 4-blade of z equal to aWXY, and then it returns z.
func (z *Int) Set4Blade(aWXYZ *big.Int) *Int {
	z.r.Set3Blade(aWXYZ)
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

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Int) Conj(y *Int) *Int {
	z.l.Conj(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Dagger sets z equal to the dagger conjugate of y, and returns z.
func (z *Int) Dagger(y *Int) *Int {
	z.l.Dagger(&y.l)
	z.r.Dagger(&y.r)
	z.r.Neg(&z.r)
	return z
}

// Hodge sets z equal to the Hodge conjugate of y, and returns z.
func (z *Int) Hodge(y *Int) *Int {
	a, b := new(grassmann3.Int), new(grassmann3.Int)
	a.Set(&y.l)
	b.Set(&y.r)
	return z.SetPair(b.Dagger(b.Hodge(b)), a.Hodge(a))
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
	a, b, temp := new(grassmann3.Int), new(grassmann3.Int), new(grassmann3.Int)
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
func (z *Int) Commutator(x, y *Int) *Int {
	return z.Sub(
		z.Mul(x, y),
		new(Int).Mul(y, x),
	)
}

// Associator sets z equal to the associator of w, x, and y:
// 		Mul(Mul(w, x), y) - Mul(w, Mul(x, y))
// Then it returns z.
func (z *Int) Associator(w, x, y *Int) *Int {
	temp := new(Int)
	return z.Sub(
		z.Mul(z.Mul(w, x), y),
		temp.Mul(w, temp.Mul(x, y)),
	)
}

// Quad returns the quadrance of z. If z = a+..., then
// the quadrance is
// 		a²
// This is always non-negative.
func (z *Int) Quad() *big.Int {
	return z.l.Quad()
}

// IsZeroDivisor returns true if z is a zero divisor.
func (z *Int) IsZeroDivisor() bool {
	return z.l.IsZeroDivisor()
}

// QuoL sets z equal to the left quotient of x and y:
// 		Mul(Inv(y), x)
// Then it returns z. If y is a zero divisor, then QuoL panics.
func (z *Int) QuoL(x, y *Int) *Int {
	if y.IsZeroDivisor() {
		panic(zeroDivisorDenominator)
	}
	return z.Divide(z.Mul(z.Conj(y), x), y.Quad())
}

// QuoR sets z equal to the right quotient of x and y:
// 		Mul(x, Inv(y))
// Then it returns z. If y is a zero divisor, then QuoR panics.
func (z *Int) QuoR(x, y *Int) *Int {
	if y.IsZeroDivisor() {
		panic(zeroDivisorDenominator)
	}
	return z.Divide(z.Mul(x, z.Conj(y)), y.Quad())
}

// Generate returns a random Int value for quick.Check testing.
func (z *Int) Generate(rand *rand.Rand, size int) reflect.Value {
	randomInt := &Int{
		*grassmann3.NewInt(
			big.NewInt(rand.Int63()),
			big.NewInt(rand.Int63()),
			big.NewInt(rand.Int63()),
			big.NewInt(rand.Int63()),
			big.NewInt(rand.Int63()),
			big.NewInt(rand.Int63()),
			big.NewInt(rand.Int63()),
			big.NewInt(rand.Int63()),
		),
		*grassmann3.NewInt(
			big.NewInt(rand.Int63()),
			big.NewInt(rand.Int63()),
			big.NewInt(rand.Int63()),
			big.NewInt(rand.Int63()),
			big.NewInt(rand.Int63()),
			big.NewInt(rand.Int63()),
			big.NewInt(rand.Int63()),
			big.NewInt(rand.Int63()),
		),
	}
	return reflect.ValueOf(randomInt)
}
