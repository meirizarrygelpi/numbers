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

// A Float is a four-dimensional Grassmann number with big.Float components.
type Float struct {
	l, r grassmann3.Float
}

// Acc returns the accuracy of the real part of z.
func (z *Float) Acc() big.Accuracy {
	return z.l.Acc()
}

// Mode returns the rounding mode of the real part of z.
func (z *Float) Mode() big.RoundingMode {
	return z.l.Mode()
}

// Prec returns the precision in bits of the real part of z.
func (z *Float) Prec() uint {
	return z.l.Prec()
}

// SetMode sets the rounding mode of z, and then it returns z.
func (z *Float) SetMode(mode big.RoundingMode) *Float {
	z.l.SetMode(mode)
	z.r.SetMode(mode)
	return z
}

// SetPrec sets the precision of z, and then it returns z.
func (z *Float) SetPrec(prec uint) *Float {
	z.l.SetPrec(prec)
	z.r.SetPrec(prec)
	return z
}

// Zero sets z equal to 0, and then it returns z. Each component has 64 bits of
// precision.
func (z *Float) Zero() *Float {
	z.l.Zero()
	z.r.Zero()
	return z
}

// One sets z equal to 1, and then returns z.
func (z *Float) One() *Float {
	z.l.One()
	z.r.Zero()
	return z
}

// Real returns the real part of z.
func (z *Float) Real() *big.Float {
	return z.l.Real()
}

// Unreal returns the unreal part of z.
func (z *Float) Unreal() *[15]*big.Float {
	var v [15]*big.Float
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

func sprintFloat(a *big.Float) string {
	if a.Signbit() {
		return a.String()
	}
	if a.IsInf() {
		return "+Inf"
	}
	return "+" + a.String()
}

// String returns the string version of a Float value.
func (z *Float) String() string {
	v := z.Unreal()
	a := make([]string, 33)
	a[0] = leftBracket
	a[1] = z.l.Real().String()
	i := 2
	for j, u := range unitNames {
		a[i] = sprintFloat(v[j])
		a[i+1] = u
		i += 2
	}
	a[32] = rightBracket
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Float) Equals(y *Float) bool {
	return z.l.Equals(&y.l) && z.r.Equals(&y.r)
}

// Set sets z equal to y, and returns z.
func (z *Float) Set(y *Float) *Float {
	z.l.Set(&y.l)
	z.r.Set(&y.r)
	return z
}

// SetPair sets z equal to a four-dimensional Grassmann number made with a given pair, and
// then it returns z.
func (z *Float) SetPair(a, b *grassmann3.Float) *Float {
	z.l.Set(a)
	z.r.Set(b)
	return z
}

// Set0Blade sets the 0-blade of z equal to a0, and then it returns z.
func (z *Float) Set0Blade(a0 *big.Float) *Float {
	z.l.Set0Blade(a0)
	return z
}

// Set1Blades sets the 1-blades of z equal to aW, aX, aY, and aZ, and then it
// returns z.
func (z *Float) Set1Blades(aW, aX, aY, aZ *big.Float) *Float {
	z.l.Set1Blades(aW, aX, aY)
	z.r.Set0Blade(aZ)
	return z
}

// Set2Blades sets the 2-blades of z equal to aWX, aWY, aXY, aWZ, aXZ, and aYZ,
// then it returns z.
func (z *Float) Set2Blades(aWX, aWY, aXY, aWZ, aXZ, aYZ *big.Float) *Float {
	z.l.Set2Blades(aWX, aWY, aXY)
	z.r.Set1Blades(aWZ, aXZ, aYZ)
	return z
}

// Set3Blades sets the 3-blade of z equal to aWXY, and then it returns z.
func (z *Float) Set3Blades(aWXY, aWXZ, aWYZ, aXYZ *big.Float) *Float {
	z.l.Set3Blade(aWXY)
	z.r.Set2Blades(aWXZ, aWYZ, aXYZ)
	return z
}

// Set4Blade sets the 4-blade of z equal to aWXY, and then it returns z.
func (z *Float) Set4Blade(aWXYZ *big.Float) *Float {
	z.r.Set3Blade(aWXYZ)
	return z
}

// Dilate sets z equal to y dilated by a, and returns z.
func (z *Float) Dilate(y *Float, a *big.Float) *Float {
	z.l.Dilate(&y.l, a)
	z.r.Dilate(&y.r, a)
	return z
}

// Divide sets z equal to y contracted by a, and returns z.
func (z *Float) Divide(y *Float, a *big.Float) *Float {
	z.l.Divide(&y.l, a)
	z.r.Divide(&y.r, a)
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Float) Neg(y *Float) *Float {
	z.l.Neg(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Float) Conj(y *Float) *Float {
	z.l.Conj(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Dagger sets z equal to the dagger conjugate of y, and returns z.
func (z *Float) Dagger(y *Float) *Float {
	z.l.Dagger(&y.l)
	z.r.Dagger(&y.r)
	z.r.Neg(&z.r)
	return z
}

// Hodge sets z equal to the Hodge conjugate of y, and returns z.
func (z *Float) Hodge(y *Float) *Float {
	a, b := new(grassmann3.Float), new(grassmann3.Float)
	a.Set(&y.l)
	b.Set(&y.r)
	return z.SetPair(b.Dagger(b.Hodge(b)), a.Hodge(a))
}

// Add sets z equal to x+y, and returns z.
func (z *Float) Add(x, y *Float) *Float {
	z.l.Add(&x.l, &y.l)
	z.r.Add(&x.r, &y.r)
	return z
}

// Sub sets z equal to x-y, and returns z.
func (z *Float) Sub(x, y *Float) *Float {
	z.l.Sub(&x.l, &y.l)
	z.r.Sub(&x.r, &y.r)
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
func (z *Float) Mul(x, y *Float) *Float {
	a, b, temp := new(grassmann3.Float), new(grassmann3.Float), new(grassmann3.Float)
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
func (z *Float) Commutator(x, y *Float) *Float {
	return z.Sub(
		z.Mul(x, y),
		new(Float).Mul(y, x),
	)
}

// Associator sets z equal to the associator of w, x, and y:
// 		Mul(Mul(w, x), y) - Mul(w, Mul(x, y))
// Then it returns z.
func (z *Float) Associator(w, x, y *Float) *Float {
	temp := new(Float)
	return z.Sub(
		z.Mul(z.Mul(w, x), y),
		temp.Mul(w, temp.Mul(x, y)),
	)
}

// Quad returns the quadrance of z. If z = a+..., then
// the quadrance is
// 		a²
// This is always non-negative.
func (z *Float) Quad() *big.Float {
	return z.l.Quad()
}

// IsZeroDivisor returns true if z is a zero divisor.
func (z *Float) IsZeroDivisor() bool {
	return z.l.IsZeroDivisor()
}

// Inv sets z equal to the inverse of y, and returns z. If y is a zero divisor,
// then Inv panics.
func (z *Float) Inv(y *Float) *Float {
	if zero := new(Float); y.Equals(zero) {
		panic(zeroDivisorInverse)
	}
	return z.Divide(z.Conj(y), y.Quad())
}

// QuoL sets z equal to the left quotient of x and y:
// 		Mul(Inv(y), x)
// Then it returns z. If y is a zero divisor, then QuoL panics.
func (z *Float) QuoL(x, y *Float) *Float {
	if y.IsZeroDivisor() {
		panic(zeroDivisorDenominator)
	}
	return z.Mul(z.Inv(y), x)
}

// QuoR sets z equal to the right quotient of x and y:
// 		Mul(x, Inv(y))
// Then it returns z. If y is a zero divisor, then QuoR panics.
func (z *Float) QuoR(x, y *Float) *Float {
	if y.IsZeroDivisor() {
		panic(zeroDivisorDenominator)
	}
	return z.Mul(x, z.Inv(y))
}

// SelfDual sets z equal to the self-dual part of y. If z is self-dual, then
//     Hodge(z) = z
// Then it returns z.
func (z *Float) SelfDual(y *Float) *Float {
	sd := new(Float)
	sd.Hodge(y)
	sd.Add(y, sd)
	sd.Divide(sd, big.NewFloat(2))
	return z.Set(sd)
}

// AntiSelfDual sets z equal to the anti-self-dual part of y. If z is
// anti-self-dual, then
//     Hodge(z) = -z
// Then it returns z.
func (z *Float) AntiSelfDual(y *Float) *Float {
	asd := new(Float)
	asd.Hodge(y)
	asd.Sub(y, asd)
	asd.Divide(asd, big.NewFloat(2))
	return z.Set(asd)
}

// Generate returns a random Float value for quick.Check testing.
func (z *Float) Generate(rand *rand.Rand, size int) reflect.Value {
	randomFloat := &Float{
		*grassmann3.NewFloat(
			big.NewFloat(rand.Float64()),
			big.NewFloat(rand.Float64()),
			big.NewFloat(rand.Float64()),
			big.NewFloat(rand.Float64()),
			big.NewFloat(rand.Float64()),
			big.NewFloat(rand.Float64()),
			big.NewFloat(rand.Float64()),
			big.NewFloat(rand.Float64()),
		),
		*grassmann3.NewFloat(
			big.NewFloat(rand.Float64()),
			big.NewFloat(rand.Float64()),
			big.NewFloat(rand.Float64()),
			big.NewFloat(rand.Float64()),
			big.NewFloat(rand.Float64()),
			big.NewFloat(rand.Float64()),
			big.NewFloat(rand.Float64()),
			big.NewFloat(rand.Float64()),
		),
	}
	return reflect.ValueOf(randomFloat)
}
