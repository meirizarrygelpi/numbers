// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licensed under the MIT License.

package cayley

import (
	"math/big"
	"math/rand"
	"reflect"
	"strings"

	"github.com/meirizarrygelpi/numbers/cplex"
	"github.com/meirizarrygelpi/numbers/hamilton"
	"github.com/meirizarrygelpi/numbers/vec3"
	"github.com/meirizarrygelpi/numbers/vec7"
)

// An Int is a Cayley octonion with big.Int components.
type Int struct {
	l, r hamilton.Int
}

// One sets z equal to 1, and then returns z.
func (z *Int) One() *Int {
	z.l.One()
	z.r.Set(new(hamilton.Int))
	return z
}

// Real returns the real part of z.
func (z *Int) Real() *big.Int {
	return z.l.Real()
}

// Unreal returns the unreal part of z.
func (z *Int) Unreal() *vec7.Int {
	v := new(vec7.Int)
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

// String returns the string version of an Int value.
//
// If z corresponds to a+bi+cj+dk+em+fn+gp+hq, then the string is
// "(a+bi+cj+dk+em+fn+gp+hq)", similar to complex128 values.
func (z *Int) String() string {
	v := z.Unreal()
	a := make([]string, 17)
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
	a[16] = rightBracket
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

// SetPair sets z equal to a Cayley octonion made with a given pair, and
// then it returns z.
func (z *Int) SetPair(a, b *hamilton.Int) *Int {
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
func (z *Int) SetUnreal(v *vec7.Int) *Int {
	u := vec3.Int([3]*big.Int{v[0], v[1], v[2]})
	z.l.SetUnreal(&u)
	z.r.SetReal(v[3])
	u = vec3.Int([3]*big.Int{v[4], v[5], v[6]})
	z.r.SetUnreal(&u)
	return z
}

// NewInt returns a pointer to the Int value a+bi+cj+dk+em+fn+gp+hq.
func NewInt(a, b, c, d, e, f, g, h *big.Int) *Int {
	z := new(Int)
	z.l.SetPair(
		cplex.NewInt(a, b),
		cplex.NewInt(c, d),
	)
	z.r.SetPair(
		cplex.NewInt(e, f),
		cplex.NewInt(g, h),
	)
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
	a, b, temp := new(hamilton.Int), new(hamilton.Int), new(hamilton.Int)
	a.Sub(
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

// Quad returns the quadrance of z. If z = a+bi+cj+dk+em+fn+gp+hq, then the
// quadrance is
// 		a² + b² + c² + d² + e² + f² + g² + h²
// This is always non-negative.
func (z *Int) Quad() *big.Int {
	return new(big.Int).Add(
		z.l.Quad(),
		z.r.Quad(),
	)
}

// QuoL sets z equal to the left quotient of x and y:
// 		Mul(Inv(y), x)
// Then it returns z. If y is zero, then QuoL panics.
func (z *Int) QuoL(x, y *Int) *Int {
	if zero := new(Int); y.Equals(zero) {
		panic(zeroDenominator)
	}
	return z.Divide(z.Mul(z.Conj(y), x), y.Quad())
}

// QuoR sets z equal to the right quotient of x and y:
// 		Mul(x, Inv(y))
// Then it returns z. If y is zero, then QuoR panics.
func (z *Int) QuoR(x, y *Int) *Int {
	if zero := new(Int); y.Equals(zero) {
		panic(zeroDenominator)
	}
	return z.Divide(z.Mul(x, z.Conj(y)), y.Quad())
}

// Generate returns a random Int value for quick.Check testing.
func (z *Int) Generate(rand *rand.Rand, size int) reflect.Value {
	randomInt := &Int{
		*hamilton.NewInt(
			big.NewInt(rand.Int63()),
			big.NewInt(rand.Int63()),
			big.NewInt(rand.Int63()),
			big.NewInt(rand.Int63()),
		),
		*hamilton.NewInt(
			big.NewInt(rand.Int63()),
			big.NewInt(rand.Int63()),
			big.NewInt(rand.Int63()),
			big.NewInt(rand.Int63()),
		),
	}
	return reflect.ValueOf(randomInt)
}
