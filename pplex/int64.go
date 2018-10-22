// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licensed under the MIT License.

package pplex

import (
	"fmt"
	"math/rand"
	"reflect"
	"strings"

	"github.com/meirizarrygelpi/numbers/maclaurin"
	"github.com/meirizarrygelpi/numbers/pade"
)

// An Int64 is a perplex number with int64 components.
type Int64 struct {
	l, r int64
}

// One sets z equal to 1, and then it returns z.
func (z *Int64) One() *Int64 {
	z.l = 1
	z.r = 0
	return z
}

// Real returns the real part of z.
func (z *Int64) Real() int64 {
	return z.l
}

// Unreal returns the unreal part of z.
func (z *Int64) Unreal() int64 {
	return z.r
}

// String returns the string version of an Int64 value.
//
// If z corresponds to a + bs, then the string is "(a+bs)", similar to
// complex128 values.
func (z *Int64) String() string {
	a := make([]string, 5)
	a[0] = leftBracket
	a[1] = fmt.Sprint(z.l)
	if z.r < 0 {
		a[2] = fmt.Sprint(z.r)
	} else {
		a[2] = "+" + fmt.Sprint(z.r)
	}
	a[3] = unitName
	a[4] = rightBracket
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Int64) Equals(y *Int64) bool {
	return z.l == y.l && z.r == y.r
}

// Set sets z equal to y, and returns z.
func (z *Int64) Set(y *Int64) *Int64 {
	z.l = y.l
	z.r = y.r
	return z
}

// SetPair sets z equal to a perplex number made with a given pair, and then
// it returns z.
func (z *Int64) SetPair(a, b int64) *Int64 {
	z.l = a
	z.r = b
	return z
}

// SetReal sets the real part of z equal to a, and then it returns z.
func (z *Int64) SetReal(a int64) *Int64 {
	z.l = a
	return z
}

// SetUnreal sets the unreal part of z equal to b, and then it returns z.
func (z *Int64) SetUnreal(b int64) *Int64 {
	z.r = b
	return z
}

// NewInt64 returns a pointer to the Int64 value a+bs.
func NewInt64(a, b int64) *Int64 {
	z := new(Int64)
	z.SetPair(a, b)
	return z
}

// Plus sets z equal to y+a, with a real, and returns z.
func (z *Int64) Plus(y *Int64, a int64) *Int64 {
	z.l = y.l + a
	z.r = y.r
	return z
}

// Minus sets z equal to y-a, with a real, and returns z.
func (z *Int64) Minus(y *Int64, a int64) *Int64 {
	z.l = y.l - a
	z.r = y.r
	return z
}

// Dilate sets z equal to y dilated by a, and returns z.
func (z *Int64) Dilate(y *Int64, a int64) *Int64 {
	z.l = y.l * a
	z.r = y.r * a
	return z
}

// Divide sets z equal to y contracted by a, and returns z.
func (z *Int64) Divide(y *Int64, a int64) *Int64 {
	z.l = y.l / a
	z.r = y.r / a
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Int64) Neg(y *Int64) *Int64 {
	z.l = -y.l
	z.r = -y.r
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Int64) Conj(y *Int64) *Int64 {
	z.l = y.l
	z.r = -y.r
	return z
}

// Add sets z equal to x+y, and returns z.
func (z *Int64) Add(x, y *Int64) *Int64 {
	z.l = x.l + y.l
	z.r = x.r + y.r
	return z
}

// Sub sets z equal to x-y, and returns z.
func (z *Int64) Sub(x, y *Int64) *Int64 {
	z.l = x.l - y.l
	z.r = x.r - y.r
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
func (z *Int64) Mul(x, y *Int64) *Int64 {
	a := (x.l * y.l) + (y.r * x.r)
	b := (x.r * y.l) + (y.r * x.l)
	z.SetPair(a, b)
	return z
}

// Dot returns the dot product of y and z. If z = a+bs and y = c+ds, then the
// dot product is
// 		ac - bd
// This can be positive, negative, or zero. The dot product is equivalent to
// 		½(Mul(Conj(z), y) + Mu(Conj(y), z))
// In this form it is clear that Dot is symmetric.
func (z *Int64) Dot(y *Int64) int64 {
	return (z.l * y.l) - (z.r * y.r)
}

// Quad returns the quadrance of z. If z = a+bs, then the quadrance is
// 		a² - b²
// This can be positive, negative, or zero.
func (z *Int64) Quad() int64 {
	return z.Dot(z)
}

// Cross returns the cross product of y and z. If z = a+bs and y = c+ds, then
// the cross product is
// 		ad - bc
// This can be positive, negative, or zero. The cross product is equivalent to
// the unreal part of
// 		½(Mul(Conj(z), y) - Mu(Conj(y), z))
// In this form it is clear that Cross is anti-symmetric.
func (z *Int64) Cross(y *Int64) int64 {
	return (z.l * y.r) - (z.r * y.l)
}

// IsZeroDivisor returns true if z is a zero divisor. This is equivalent to z
// having zero quadrance.
func (z *Int64) IsZeroDivisor() bool {
	return z.l == z.r || z.l == -z.r
}

// Quo sets z equal to the quotient of x and y, and returns z. If y is a zero
// divisor, then Quo panics.
func (z *Int64) Quo(x, y *Int64) *Int64 {
	if y.IsZeroDivisor() {
		panic(zeroDivisorDenominator)
	}
	q := y.Quad()
	a := (x.l * y.l) - (y.r * x.r)
	b := (x.r * y.l) - (y.r * x.l)
	z.SetPair(a, b)
	return z.Divide(z, q)
}

// Maclaurin sets z equal to the value of the Maclaurin polynomial p evaluated
// at y, and returns z. Horner's method is used.
func (z *Int64) Maclaurin(y *Int64, p *maclaurin.Int64) *Int64 {
	if p.Len() == 0 {
		z = new(Int64)
		return z
	}
	n := p.Degree
	var a int64
	if n == 0 {
		z = new(Int64)
		a, _ = p.Coeff(n)
		z.SetReal(a)
		return z
	}
	a, _ = p.Coeff(n)
	z.Dilate(y, a)
	for n > 1 {
		n--
		if a, ok := p.Coeff(n); ok {
			z.Plus(z, a)
		}
		z.Mul(z, y)
	}
	if a, ok := p.Coeff(0); ok {
		z.Plus(z, a)
	}
	return z
}

// Padé sets z equal to the value of the Padé approximant r evaluated at y,
// and returns z.
func (z *Int64) Padé(y *Int64, r *pade.Int64) *Int64 {
	p, q := new(Int64), new(Int64)
	p.Maclaurin(y, &r.P)
	q.Maclaurin(y, &r.Q)
	return z.Quo(p, q)
}

// Generate returns a random Int64 value for quick.Check testing.
func (z *Int64) Generate(rand *rand.Rand, size int) reflect.Value {
	randomInt64 := &Int64{
		rand.Int63(),
		rand.Int63(),
	}
	return reflect.ValueOf(randomInt64)
}
