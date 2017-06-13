// Copyright (c) 2017 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package comhamilton

import (
	"math/rand"
	"reflect"
	"strings"

	"fmt"

	"github.com/meirizarrygelpi/numbers/cplex"
	"github.com/meirizarrygelpi/numbers/hamilton"
	"github.com/meirizarrygelpi/numbers/vec7"
)

// An Int64 is an complex-Hamilton quaternion with int64 components.
type Int64 struct {
	l, r hamilton.Int64
}

// One sets z equal to 1, and then returns z.
func (z *Int64) One() *Int64 {
	z.l.One()
	z.r.Set(new(hamilton.Int64))
	return z
}

// Real returns the real part of z.
func (z *Int64) Real() int64 {
	return z.l.Real()
}

// Unreal returns the unreal part of z.
func (z *Int64) Unreal() *vec7.Int64 {
	v := new(vec7.Int64)
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

// String returns the string version of a Int64 value.
//
// If z corresponds to a+bi+cj+dk+eH+fiH+gjH+hkH, then the string is
// "(a+bi+cj+dk+eH+fiH+gjH+hkH)", similar to complex128 values.
func (z *Int64) String() string {
	v := z.Unreal()
	a := make([]string, 17)
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
	a[16] = rightBracket
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

// SetPair sets z equal to an complex-Hamilton quaternion made with a given pair, and
// then it returns z.
func (z *Int64) SetPair(a, b *hamilton.Int64) *Int64 {
	z.l.Set(a)
	z.r.Set(b)
	return z
}

// NewInt64 returns a pointer to the Int64 value a+bi+cj+dk+eH+fiH+gjH+hkH.
func NewInt64(a, b, c, d, e, f, g, h int64) *Int64 {
	z := new(Int64)
	z.l.SetPair(
		cplex.NewInt64(a, b),
		cplex.NewInt64(c, d),
	)
	z.r.SetPair(
		cplex.NewInt64(e, f),
		cplex.NewInt64(g, h),
	)
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

// Bar sets z equal to the Hamilton conjugate of y, and returns z.
func (z *Int64) Bar(y *Int64) *Int64 {
	z.l.Conj(&y.l)
	z.r.Conj(&y.r)
	return z
}

// Tilde sets z equal to the H-conjugate of y, and returns z.
func (z *Int64) Tilde(y *Int64) *Int64 {
	z.l.Set(&y.l)
	z.r.Neg(&y.r)
	return z
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
	a, b, temp := new(hamilton.Int64), new(hamilton.Int64), new(hamilton.Int64)
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

// Commutator sets z equal to the commutator of x and y:
// 		Mul(x, y) - Mul(y, x)
// Then it returns z.
func (z *Int64) Commutator(x, y *Int64) *Int64 {
	return z.Sub(
		z.Mul(x, y),
		new(Int64).Mul(y, x),
	)
}

// Quad returns the quadrance of z. This is a complex number.
func (z *Int64) Quad() *cplex.Int64 {
	q := new(hamilton.Int64)
	q.Mul(&z.l, q.Conj(&z.r))
	q.Dilate(q, 2)
	return cplex.NewInt64(z.l.Quad()-z.r.Quad(), q.Real())
}

// Norm returns the norm of z. This is always non-negative.
func (z *Int64) Norm() int64 {
	return z.Quad().Quad()
}

// IsZeroDivisor returns true if z is a zero divisor.
func (z *Int64) IsZeroDivisor() bool {
	zero := new(hamilton.Int64)
	return z.l.Equals(zero)
}

// QuoL sets z equal to the left quotient of x and y:
// 		Mul(Inv(y), x)
// Then it returns z. If y is a zero divisor, then QuoL panics.
func (z *Int64) QuoL(x, y *Int64) *Int64 {
	if y.IsZeroDivisor() {
		panic(zeroDivisorDenominator)
	}
	n := y.Norm()
	temp := new(Int64)
	z.Mul(temp.Bar(y), x)
	z.Mul(temp.Tilde(y), z)
	z.Mul(temp.Bar(temp), z)
	return z.Divide(z, n)
}

// QuoR sets z equal to the right quotient of x and y:
// 		Mul(x, Inv(y))
// Then it returns z. If y is a zero divisor, then QuoR panics.
func (z *Int64) QuoR(x, y *Int64) *Int64 {
	if y.IsZeroDivisor() {
		panic(zeroDivisorDenominator)
	}
	n := y.Norm()
	temp := new(Int64)
	z.Mul(x, temp.Bar(y))
	z.Mul(z, temp.Tilde(y))
	z.Mul(z, temp.Bar(temp))
	return z.Divide(z, n)
}

// Generate returns a random Int64 value for quick.Check testing.
func (z *Int64) Generate(rand *rand.Rand, size int) reflect.Value {
	randomInt64 := &Int64{
		*hamilton.NewInt64(
			rand.Int63(),
			rand.Int63(),
			rand.Int63(),
			rand.Int63(),
		),
		*hamilton.NewInt64(
			rand.Int63(),
			rand.Int63(),
			rand.Int63(),
			rand.Int63(),
		),
	}
	return reflect.ValueOf(randomInt64)
}
