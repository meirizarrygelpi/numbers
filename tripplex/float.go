// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package tripplex

import (
	"math/big"
	"math/rand"
	"reflect"
	"strings"

	"github.com/meirizarrygelpi/numbers/bipplex"
	"github.com/meirizarrygelpi/numbers/pplex"
	"github.com/meirizarrygelpi/numbers/vec7"
)

// A Float is a tri-perplex number with big.Float components.
type Float struct {
	l, r bipplex.Float
}

// Acc returns the accuracy of the real part of z.
func (z *Float) Acc() big.Accuracy {
	return z.l.Acc()
}

// Mode returns the rounding mode of the real part of z.
func (z *Float) Mode() big.RoundingMode {
	return z.l.Mode()
}

// Prec returns the precision of z in bits of the real part of z.
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
func (z *Float) Unreal() *vec7.Float {
	v := new(vec7.Float)
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
//
// If z corresponds to a+bs+cT+dsT+eU+fsU+gTU+hsTU, then the string is
// "⦗a+bs+cT+dsT+eU+fsU+gTU+hsTU⦘", similar to perplex128 values.
func (z *Float) String() string {
	v := z.Unreal()
	a := make([]string, 17)
	a[0] = leftBracket
	a[1] = z.l.Real().String()
	i := 2
	for j, u := range [7]string{unit1, unit2, unit3, unit4, unit5, unit6, unit7} {
		a[i] = sprintFloat(v[j])
		a[i+1] = u
		i += 2
	}
	a[16] = rightBracket
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

// SetPair sets z equal to a tri-perplex number made with a given pair, and
// then it returns z.
func (z *Float) SetPair(a, b *bipplex.Float) *Float {
	z.l.Set(a)
	z.r.Set(b)
	return z
}

// NewFloat returns a pointer to the Float value a+bs+cT+dsT+eU+fsU+gTU+hsTU.
func NewFloat(a, b, c, d, e, f, g, h *big.Float) *Float {
	z := new(Float)
	z.l.SetPair(
		pplex.NewFloat(a, b),
		pplex.NewFloat(c, d),
	)
	z.r.SetPair(
		pplex.NewFloat(e, f),
		pplex.NewFloat(g, h),
	)
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

// Star1 sets z equal to the i-conjugate of y, and returns z.
func (z *Float) Star1(y *Float) *Float {
	z.l.Star1(&y.l)
	z.r.Star1(&y.r)
	return z
}

// Star2 sets z equal to the J-conjugate of y, and returns z.
func (z *Float) Star2(y *Float) *Float {
	z.l.Star2(&y.l)
	z.r.Star2(&y.r)
	return z
}

// Star3 sets z equal to the K-conjugate of y, and returns z.
func (z *Float) Star3(y *Float) *Float {
	z.l.Set(&y.l)
	z.r.Neg(&y.r)
	return z
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
//
// The multiplication table is:
//     +-----+-----+-----+-----+-----+-----+-----+-----+
//     | Mul | s   | T   | sT  | U   | sU  | TU  | sTU |
//     +-----+-----+-----+-----+-----+-----+-----+-----+
//     | s   | 1   | sT  | T   | sU  | U   | sTU | TU  |
//     +-----+-----+-----+-----+-----+-----+-----+-----+
//     | T   | sT  | 1   | s   | TU  | sTU | U   | sU  |
//     +-----+-----+-----+-----+-----+-----+-----+-----+
//     | sT  | T   | s   | 1   | sTU | TU  | sU  | U   |
//     +-----+-----+-----+-----+-----+-----+-----+-----+
//     | U   | sU  | TU  | sTU | 1   | s   | T   | sT  |
//     +-----+-----+-----+-----+-----+-----+-----+-----+
//     | sU  | U   | sTU | TU  | s   | 1   | sT  | T   |
//     +-----+-----+-----+-----+-----+-----+-----+-----+
//     | TU  | sTU | U   | sU  | T   | sT  | 1   | s   |
//     +-----+-----+-----+-----+-----+-----+-----+-----+
//     | sTU | TU  | sU  | U   | sT  | T   | s   | 1   |
//     +-----+-----+-----+-----+-----+-----+-----+-----+
// This binary operation is commutative and associative.
func (z *Float) Mul(x, y *Float) *Float {
	a, b, temp := new(bipplex.Float), new(bipplex.Float), new(bipplex.Float)
	a.Add(
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

// Quad returns the quadrance of z. If z = a+bs+cT+dsT+eU+fsU+gTU+hsTU, then
// the quadrance is
// 		a² - b² + c² - d² + 2(ab + cd)i
// Note that this is a perplex number.
func (z *Float) Quad() *bipplex.Float {
	q := new(bipplex.Float)
	return q.Sub(q.Mul(&z.l, &z.l), new(bipplex.Float).Mul(&z.r, &z.r))
}

// Norm returns the norm of z. If z = a+bs+cT+dsT+eU+fsU+gTU+hsTU, then the
// norm is
// 		(a² - b² + c² - d²)² + 4(ab + cd)²
// There is another way to write the norm as a sum of two squares:
// 		(a² + b² - c² - d²)² + 4(ac + bd)²
// Alternatively, it can also be written as a difference of two squares:
//		(a² + b² + c² + d²)² - 4(ad - bc)²
// Finally, you have the factorized form:
// 		((a - d)² + (b + c)²)((a + d)² + (b - c)²)
// In this form it is clear that the norm is always non-negative.
func (z *Float) Norm() *big.Float {
	return z.Quad().Norm()
}

// IsZeroDivisor returns true if z is a zero divisor.
func (z *Float) IsZeroDivisor() bool {
	return z.Quad().IsZeroDivisor()
}

// Inv sets z equal to the inverse of y, and returns z. If y is zero, then Inv
// panics.
func (z *Float) Inv(y *Float) *Float {
	if y.IsZeroDivisor() {
		panic(zeroDivisorInverse)
	}
	a := y.Quad()
	a.Inv(a)
	z.Star3(y)
	z.l.Mul(&z.l, a)
	z.r.Mul(&z.r, a)
	return z
}

// Quo sets z equal to the quotient of x and y, and returns z. If y is zero,
// then Quo panics.
func (z *Float) Quo(x, y *Float) *Float {
	if y.IsZeroDivisor() {
		panic(zeroDivisorDenominator)
	}
	a := y.Quad()
	a.Inv(a)
	z.Mul(x, z.Star3(y))
	z.l.Mul(&z.l, a)
	z.r.Mul(&z.r, a)
	return z
}

// CrossRatio sets z equal to the cross-ratio of v, w, x, and y:
// 		Inv(w - x) * (v - x) * Inv(v - y) * (w - y)
// Then it returns z.
func (z *Float) CrossRatio(v, w, x, y *Float) *Float {
	temp := new(Float)
	z.Sub(w, x)
	z.Inv(z)
	temp.Sub(v, x)
	z.Mul(z, temp)
	temp.Sub(v, y)
	temp.Inv(temp)
	z.Mul(z, temp)
	temp.Sub(w, y)
	return z.Mul(z, temp)
}

// Möbius sets z equal to the Möbius (fractional linear) transform of y:
// 		(a*y + b) * Inv(c*y + d)
// Then it returns z.
func (z *Float) Möbius(y, a, b, c, d *Float) *Float {
	z.Mul(a, y)
	z.Add(z, b)
	temp := new(Float)
	temp.Mul(c, y)
	temp.Add(temp, d)
	temp.Inv(temp)
	return z.Mul(z, temp)
}

// Generate returns a random Float value for quick.Check testing.
func (z *Float) Generate(rand *rand.Rand, size int) reflect.Value {
	randomFloat := &Float{
		*bipplex.NewFloat(
			big.NewFloat(rand.Float64()),
			big.NewFloat(rand.Float64()),
			big.NewFloat(rand.Float64()),
			big.NewFloat(rand.Float64()),
		),
		*bipplex.NewFloat(
			big.NewFloat(rand.Float64()),
			big.NewFloat(rand.Float64()),
			big.NewFloat(rand.Float64()),
			big.NewFloat(rand.Float64()),
		),
	}
	return reflect.ValueOf(randomFloat)
}
