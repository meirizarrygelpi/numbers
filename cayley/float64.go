// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package cayley

import (
	"math/rand"
	"reflect"
	"strings"

	"fmt"

	"math"

	"github.com/meirizarrygelpi/numbers/cplex"
	"github.com/meirizarrygelpi/numbers/hamilton"
	"github.com/meirizarrygelpi/numbers/vec7"
)

// A Float64 is a Cayley octonion with float64 components.
type Float64 struct {
	l, r hamilton.Float64
}

// One sets z equal to 1, and then returns z.
func (z *Float64) One() *Float64 {
	z.l.One()
	z.r.Set(new(hamilton.Float64))
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
// If z corresponds to a+bi+cj+dk+em+fn+gp+hq, then the string is
// "⦗a+bi+cj+dk+em+fn+gp+hq⦘", similar to complex128 values.
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

// SetPair sets z equal to a Cayley octonion made with a given pair, and
// then it returns z.
func (z *Float64) SetPair(a, b *hamilton.Float64) *Float64 {
	z.l.Set(a)
	z.r.Set(b)
	return z
}

// NewFloat64 returns a pointer to the Float64 value a+bi+cj+dk+em+fn+gp+hq.
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
	a, b, temp := new(hamilton.Float64), new(hamilton.Float64), new(hamilton.Float64)
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

// Quad returns the quadrance of z. If z = a+bi+cj+dk+em+fn+gp+hq, then the
// quadrance is
// 		a² + b² + c² + d² + e² + f² + g² + h²
// This is always non-negative.
func (z *Float64) Quad() float64 {
	return z.l.Quad() + z.r.Quad()
}

// Inv sets z equal to the inverse of y, and returns z. If y is zero, then Inv
// panics.
func (z *Float64) Inv(y *Float64) *Float64 {
	if zero := new(Float64); y.Equals(zero) {
		panic(zeroInverse)
	}
	return z.Divide(z.Conj(y), y.Quad())
}

// QuoL sets z equal to the left quotient of x and y:
// 		Mul(Inv(y), x)
// Then it returns z. If y is zero, then QuoL panics.
func (z *Float64) QuoL(x, y *Float64) *Float64 {
	if zero := new(Float64); y.Equals(zero) {
		panic(zeroDenominator)
	}
	return z.Mul(z.Inv(y), x)
}

// QuoR sets z equal to the right quotient of x and y:
// 		Mul(x, Inv(y))
// Then it returns z. If y is zero, then QuoR panics.
func (z *Float64) QuoR(x, y *Float64) *Float64 {
	if zero := new(Float64); y.Equals(zero) {
		panic(zeroDenominator)
	}
	return z.Mul(x, z.Inv(y))
}

// Graves sets z equal to the Gravesian integer a+bi+cj+dk+em+fn+gp+hq, and
// returns z.
func (z *Float64) Graves(a, b, c, d, e, f, g, h int64) *Float64 {
	z.l.Lipschitz(a, b, c, d)
	z.r.Lipschitz(e, f, g, h)
	return z
}

// Klein sets z equal to the Kleinian integer
// (a+½)+(b+½)i+(c+½)j+(d+½)k+(e+½)m+(f+½)n+(g+½)p+(h+½)q, and returns z.
func (z *Float64) Klein(a, b, c, d, e, f, g, h int64) *Float64 {
	z.Graves(a, b, c, d, e, f, g, h)
	half := 0.5
	return z.Add(z, NewFloat64(half, half, half, half, half, half, half, half))
}

// Generate returns a random Float64 value for quick.Check testing.
func (z *Float64) Generate(rand *rand.Rand, size int) reflect.Value {
	randomFloat64 := &Float64{
		*hamilton.NewFloat64(
			rand.Float64(),
			rand.Float64(),
			rand.Float64(),
			rand.Float64(),
		),
		*hamilton.NewFloat64(
			rand.Float64(),
			rand.Float64(),
			rand.Float64(),
			rand.Float64(),
		),
	}
	return reflect.ValueOf(randomFloat64)
}
