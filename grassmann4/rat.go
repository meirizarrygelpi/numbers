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

// A Rat is a four-dimensional Grassmann number with big.Rat components.
type Rat struct {
	l, r grassmann3.Rat
}

// One sets z equal to 1, and then returns z.
func (z *Rat) One() *Rat {
	z.l.One()
	z.r.Set(new(grassmann3.Rat))
	return z
}

// Real returns the real part of z.
func (z *Rat) Real() *big.Rat {
	return z.l.Real()
}

// Unreal returns the unreal part of z.
func (z *Rat) Unreal() *[15]*big.Rat {
	var v [15]*big.Rat
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

// String returns the string version of a Rat value.
func (z *Rat) String() string {
	v := z.Unreal()
	a := make([]string, 33)
	a[0] = leftBracket
	a[1] = z.l.Real().RatString()
	i := 2
	for j, u := range unitNames {
		if v[j].Sign() < 0 {
			a[i] = v[j].RatString()
		} else {
			a[i] = "+" + v[j].RatString()
		}
		a[i+1] = u
		i += 2
	}
	a[32] = rightBracket
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Rat) Equals(y *Rat) bool {
	return z.l.Equals(&y.l) && z.r.Equals(&y.r)
}

// Set sets z equal to y, and returns z.
func (z *Rat) Set(y *Rat) *Rat {
	z.l.Set(&y.l)
	z.r.Set(&y.r)
	return z
}

// SetPair sets z equal to a four-dimensional Grassmann number made with a given pair, and
// then it returns z.
func (z *Rat) SetPair(a, b *grassmann3.Rat) *Rat {
	z.l.Set(a)
	z.r.Set(b)
	return z
}

// Set0Blade sets the 0-blade of z equal to a0, and then it returns z.
func (z *Rat) Set0Blade(a0 *big.Rat) *Rat {
	z.l.Set0Blade(a0)
	return z
}

// Set1Blades sets the 1-blades of z equal to aW, aX, aY, and aZ, and then it
// returns z.
func (z *Rat) Set1Blades(aW, aX, aY, aZ *big.Rat) *Rat {
	z.l.Set1Blades(aW, aX, aY)
	z.r.Set0Blade(aZ)
	return z
}

// Set2Blades sets the 2-blades of z equal to aWX, aWY, aXY, aWZ, aXZ, and aYZ,
// then it returns z.
func (z *Rat) Set2Blades(aWX, aWY, aXY, aWZ, aXZ, aYZ *big.Rat) *Rat {
	z.l.Set2Blades(aWX, aWY, aXY)
	z.r.Set1Blades(aWZ, aXZ, aYZ)
	return z
}

// Set3Blades sets the 3-blade of z equal to aWXY, and then it returns z.
func (z *Rat) Set3Blades(aWXY, aWXZ, aWYZ, aXYZ *big.Rat) *Rat {
	z.l.Set3Blade(aWXY)
	z.r.Set2Blades(aWXZ, aWYZ, aXYZ)
	return z
}

// Set4Blade sets the 4-blade of z equal to aWXY, and then it returns z.
func (z *Rat) Set4Blade(aWXYZ *big.Rat) *Rat {
	z.r.Set3Blade(aWXYZ)
	return z
}

// Scale sets z equal to y scaled by a, and returns z.
func (z *Rat) Scale(y *Rat, a *big.Rat) *Rat {
	z.l.Scale(&y.l, a)
	z.r.Scale(&y.r, a)
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Rat) Neg(y *Rat) *Rat {
	z.l.Neg(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Rat) Conj(y *Rat) *Rat {
	z.l.Conj(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Dagger sets z equal to the dagger conjugate of y, and returns z.
func (z *Rat) Dagger(y *Rat) *Rat {
	z.l.Dagger(&y.l)
	z.r.Dagger(&y.r)
	z.r.Neg(&z.r)
	return z
}

// Hodge sets z equal to the Hodge conjugate of y, and returns z.
func (z *Rat) Hodge(y *Rat) *Rat {
	a, b := new(grassmann3.Rat), new(grassmann3.Rat)
	a.Set(&y.l)
	b.Set(&y.r)
	return z.SetPair(b.Dagger(b.Hodge(b)), a.Hodge(a))
}

// Add sets z equal to x+y, and returns z.
func (z *Rat) Add(x, y *Rat) *Rat {
	z.l.Add(&x.l, &y.l)
	z.r.Add(&x.r, &y.r)
	return z
}

// Sub sets z equal to x-y, and returns z.
func (z *Rat) Sub(x, y *Rat) *Rat {
	z.l.Sub(&x.l, &y.l)
	z.r.Sub(&x.r, &y.r)
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
func (z *Rat) Mul(x, y *Rat) *Rat {
	a, b, temp := new(grassmann3.Rat), new(grassmann3.Rat), new(grassmann3.Rat)
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
func (z *Rat) Commutator(x, y *Rat) *Rat {
	return z.Sub(
		z.Mul(x, y),
		new(Rat).Mul(y, x),
	)
}

// Associator sets z equal to the associator of w, x, and y:
// 		Mul(Mul(w, x), y) - Mul(w, Mul(x, y))
// Then it returns z.
func (z *Rat) Associator(w, x, y *Rat) *Rat {
	temp := new(Rat)
	return z.Sub(
		z.Mul(z.Mul(w, x), y),
		temp.Mul(w, temp.Mul(x, y)),
	)
}

// Quad returns the quadrance of z. If z = a+bW+cX+dWX+eY+fWY+gXY+h(WX)Y, then
// the quadrance is
//     a²
// This is always non-negative.
func (z *Rat) Quad() *big.Rat {
	return z.l.Quad()
}

// IsZeroDivisor returns true if z is a zero divisor.
func (z *Rat) IsZeroDivisor() bool {
	return z.l.IsZeroDivisor()
}

// Inv sets z equal to the inverse of y, and returns z. If y is a zero divisor,
// then Inv panics.
func (z *Rat) Inv(y *Rat) *Rat {
	if y.IsZeroDivisor() {
		panic(zeroDivisorInverse)
	}
	a := y.Quad()
	a.Inv(a)
	return z.Scale(z.Conj(y), a)
}

// QuoL sets z equal to the left quotient of x and y:
//     Mul(Inv(y), x)
// Then it returns z. If y is a zero divisor, then QuoL panics.
func (z *Rat) QuoL(x, y *Rat) *Rat {
	if y.IsZeroDivisor() {
		panic(zeroDivisorDenominator)
	}
	return z.Mul(z.Inv(y), x)
}

// QuoR sets z equal to the right quotient of x and y:
//     Mul(x, Inv(y))
// Then it returns z. If y is a zero divisor, then QuoR panics.
func (z *Rat) QuoR(x, y *Rat) *Rat {
	if y.IsZeroDivisor() {
		panic(zeroDivisorDenominator)
	}
	return z.Mul(x, z.Inv(y))
}

// SelfDual sets z equal to the self-dual part of y. If z is self-dual, then
//     Hodge(z) = z
// Then it returns z.
func (z *Rat) SelfDual(y *Rat) *Rat {
	sd := new(Rat)
	sd.Hodge(y)
	sd.Add(y, sd)
	sd.Scale(sd, big.NewRat(1, 2))
	return z.Set(sd)
}

// AntiSelfDual sets z equal to the anti-self-dual part of y. If z is
// anti-self-dual, then
//     Hodge(z) = -z
// Then it returns z.
func (z *Rat) AntiSelfDual(y *Rat) *Rat {
	asd := new(Rat)
	asd.Hodge(y)
	asd.Sub(y, asd)
	asd.Scale(asd, big.NewRat(1, 2))
	return z.Set(asd)
}

// Generate returns a random Rat value for quick.Check testing.
func (z *Rat) Generate(rand *rand.Rand, size int) reflect.Value {
	randomRat := &Rat{
		*grassmann3.NewRat(
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
		),
		*grassmann3.NewRat(
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
		),
	}
	return reflect.ValueOf(randomRat)
}
