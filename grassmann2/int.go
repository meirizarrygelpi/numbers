// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package grassmann2

import (
	"math/big"
	"math/rand"
	"reflect"
	"strings"

	"github.com/meirizarrygelpi/numbers/nplex"
	"github.com/meirizarrygelpi/numbers/vec3"
)

// An Int is a two-dimensional Grassmann number with big.Int components.
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
// If z corresponds to a+bW+cX+dWX, then the string is "(a+bW+cX+dWX)", similar
// to complex128 values.
func (z *Int) String() string {
	v := z.Unreal()
	a := make([]string, 9)
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

// SetPair sets z equal to a two-dimensional Grassmann number made with a given pair, and
// then it returns z.
func (z *Int) SetPair(a, b *nplex.Int) *Int {
	z.l.Set(a)
	z.r.Set(b)
	return z
}

// SetReal sets the real part of z equal to a, and then it returns z.
func (z *Int) SetReal(a *big.Int) *Int {
	z.l.SetReal(a)
	return z
}

// SetUnreal sets the unreal part of z equal to v, and then it returns z.
func (z *Int) SetUnreal(v *vec3.Int) *Int {
	z.l.SetUnreal(v[0])
	z.r.SetReal(v[1])
	z.r.SetUnreal(v[2])
	return z
}

// Set0Blade sets the 0-blade of z equal to a0, and then it returns z.
func (z *Int) Set0Blade(a0 *big.Int) *Int {
	z.l.SetReal(a0)
	return z
}

// Set1Blades sets the 1-blades of z equal to aW and aX, and then it returns z.
func (z *Int) Set1Blades(aW, aX *big.Int) *Int {
	z.l.SetUnreal(aW)
	z.r.SetReal(aX)
	return z
}

// Set2Blade sets the 2-blade of z equal to aWX, and then it returns z.
func (z *Int) Set2Blade(aWX *big.Int) *Int {
	z.r.SetUnreal(aWX)
	return z
}

// NewInt returns a pointer to the Int value a+bW+cX+dWX.
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

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Int) Conj(y *Int) *Int {
	z.l.Conj(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Dagger sets z equal to the dagger conjugate of y, and returns z.
func (z *Int) Dagger(y *Int) *Int {
	z.l.Conj(&y.l)
	z.r.Minus(&y.r)
	return z
}

// Hodge sets z equal to the Hodge conjugate of y, and returns z.
func (z *Int) Hodge(y *Int) *Int {
	a, b := new(nplex.Int), new(nplex.Int)
	a.Set(&y.l)
	b.Set(&y.r)
	return z.SetPair(b.Conj(b.Hodge(b)), a.Hodge(a))
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
	a, b, temp := new(nplex.Int), new(nplex.Int), new(nplex.Int)
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

// Quad returns the quadrance of z. If z = a+bW+cX+dWX, then the quadrance is
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
// Then it returns z. If y is zero, then QuoL panics.
func (z *Int) QuoL(x, y *Int) *Int {
	if y.IsZeroDivisor() {
		panic(zeroDivisorDenominator)
	}
	return z.Divide(z.Mul(z.Conj(y), x), y.Quad())
}

// QuoR sets z equal to the right quotient of x and y:
// 		Mul(x, Inv(y))
// Then it returns z. If y is zero, then QuoR panics.
func (z *Int) QuoR(x, y *Int) *Int {
	if y.IsZeroDivisor() {
		panic(zeroDivisorDenominator)
	}
	return z.Divide(z.Mul(x, z.Conj(y)), y.Quad())
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
