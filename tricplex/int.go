// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package tricplex

import (
	"math/big"
	"math/rand"
	"reflect"
	"strings"

	"github.com/meirizarrygelpi/numbers/bicplex"
	"github.com/meirizarrygelpi/numbers/cplex"
	"github.com/meirizarrygelpi/numbers/vec7"
)

// An Int is a tri-complex number with big.Int components.
type Int struct {
	l, r bicplex.Int
}

// One sets z equal to 1, and then returns z.
func (z *Int) One() *Int {
	z.l.One()
	z.r.Set(new(bicplex.Int))
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
// If z corresponds to a+bi+cJ+diJ+eK+fiK+gJK+hiJK, then the string is
// "(a+bi+cJ+diJ+eK+fiK+gJK+hiJK)", similar to complex128 values.
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

// SetPair sets z equal to a tri-complex number made with a given pair, and
// then it returns z.
func (z *Int) SetPair(a, b *bicplex.Int) *Int {
	z.l.Set(a)
	z.r.Set(b)
	return z
}

// NewInt returns a pointer to the Int value a+bi+cJ+diJ+eK+fiK+gJK+hiJK.
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

// Bar sets z equal to the i-conjugate of y, and returns z.
func (z *Int) Bar(y *Int) *Int {
	z.l.Bar(&y.l)
	z.r.Bar(&y.r)
	return z
}

// Tilde sets z equal to the J-conjugate of y, and returns z.
func (z *Int) Tilde(y *Int) *Int {
	z.l.Tilde(&y.l)
	z.r.Tilde(&y.r)
	return z
}

// Star sets z equal to the K-conjugate of y, and returns z.
func (z *Int) Star(y *Int) *Int {
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
	a, b, temp := new(bicplex.Int), new(bicplex.Int), new(bicplex.Int)
	a.Sub(
		a.Mul(&x.l, &y.l),
		temp.Mul(&x.r, &y.r),
	)
	b.Add(
		b.Mul(&x.l, &y.r),
		temp.Mul(&x.r, &y.l),
	)
	z.SetPair(a, b)
	return z
}

// Quad returns the quadrance of z. If z = a+bi+cJ+diJ+eK+fiK+gJK+hiJK, then
// the quadrance is
// 		a² - b² + c² - d² + 2(ab + cd)i
// Note that this is a complex number.
func (z *Int) Quad() *bicplex.Int {
	q := new(bicplex.Int)
	return q.Add(q.Mul(&z.l, &z.l), new(bicplex.Int).Mul(&z.r, &z.r))
}

// Norm returns the norm of z. If z = a+bi+cJ+diJ+eK+fiK+gJK+hiJK, then the
// norm is
// 		(a² - b² + c² - d²)² + 4(ab + cd)²
// There is another way to write the norm as a sum of two squares:
// 		(a² + b² - c² - d²)² + 4(ac + bd)²
// Alternatively, it can also be written as a difference of two squares:
//		(a² + b² + c² + d²)² - 4(ad - bc)²
// Finally, you have the factorized form:
// 		((a - d)² + (b + c)²)((a + d)² + (b - c)²)
// In this form it is clear that the norm is always non-negative.
func (z *Int) Norm() *big.Int {
	return z.Quad().Norm()
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
	a := y.Quad()
	z.Mul(x, z.Star(y))
	z.l.Quo(&z.l, a)
	z.r.Quo(&z.r, a)
	return z
}

// Generate returns a random Int value for quick.Check testing.
func (z *Int) Generate(rand *rand.Rand, size int) reflect.Value {
	randomInt := &Int{
		*bicplex.NewInt(
			big.NewInt(rand.Int63()),
			big.NewInt(rand.Int63()),
			big.NewInt(rand.Int63()),
			big.NewInt(rand.Int63()),
		),
		*bicplex.NewInt(
			big.NewInt(rand.Int63()),
			big.NewInt(rand.Int63()),
			big.NewInt(rand.Int63()),
			big.NewInt(rand.Int63()),
		),
	}
	return reflect.ValueOf(randomInt)
}
