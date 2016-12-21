// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package cockle

import (
	"math/big"
	"math/rand"
	"reflect"
	"strings"

	"github.com/meirizarrygelpi/numbers/cplex"
	"github.com/meirizarrygelpi/numbers/vec3"
)

// A Float is a Cockle quaternion with big.Float components.
type Float struct {
	l, r cplex.Float
}

// Acc returns the accuracy of the real part of z.
func (z *Float) Acc() big.Accuracy {
	return z.l.Acc()
}

// Mode returns the accuracy of the real part of z.
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
func (z *Float) Unreal() *vec3.Float {
	v := new(vec3.Float)
	v[0] = z.l.Unreal()
	v[1] = z.r.Real()
	v[2] = z.r.Unreal()
	return v
}

// String returns the string version of a Float value.
//
// If z corresponds to a+bi+ct+du, then the string is "⦗a+bi+ct+du⦘", similar
// to complex128 values.
func (z *Float) String() string {
	v := z.Unreal()
	a := make([]string, 9)
	a[0] = leftBracket
	a[1] = z.l.Real().String()
	if v[0].Sign() < 0 {
		a[2] = v[0].String()
	} else {
		a[2] = "+" + v[0].String()
	}
	a[3] = unit1
	if v[1].Sign() < 0 {
		a[4] = v[1].String()
	} else {
		a[4] = "+" + v[1].String()
	}
	a[5] = unit2
	if v[2].Sign() < 0 {
		a[6] = v[2].String()
	} else {
		a[6] = "+" + v[2].String()
	}
	a[7] = unit3
	a[8] = rightBracket
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

// SetPair sets z equal to a Cockle quaternion made with a given pair, and
// then it returns z.
func (z *Float) SetPair(a, b *cplex.Float) *Float {
	z.l.Set(a)
	z.r.Set(b)
	return z
}

// NewFloat returns a pointer to the Float value a+bi+ct+du.
func NewFloat(a, b, c, d *big.Float) *Float {
	z := new(Float)
	z.l.SetPair(a, b)
	z.r.SetPair(c, d)
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
// The multiplication rule is:
// 		Mul(i, i) = Mul(j, j) = Mul(k, k) = -1
// 		Mul(i, j) = -Mul(j, i) = k
// 		Mul(j, k) = -Mul(k, j) = i
// 		Mul(k, i) = -Mul(i, k) = j
// This binary operation is non-commutative but associative.
func (z *Float) Mul(x, y *Float) *Float {
	a, b, temp := new(cplex.Float), new(cplex.Float), new(cplex.Float)
	a.Add(
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
func (z *Float) Commutator(x, y *Float) *Float {
	return z.Sub(
		z.Mul(x, y),
		new(Float).Mul(y, x),
	)
}

// Quad returns the quadrance of z. If z = a+bi+ct+du, then the quadrance is
// 		a² + b² - c² - d²
// This can be positive, negative, or zero.
func (z *Float) Quad() *big.Float {
	return new(big.Float).Sub(
		z.l.Quad(),
		z.r.Quad(),
	)
}

// IsZeroDivisor returns true if z is a zero divisor.
func (z *Float) IsZeroDivisor() bool {
	return z.l.Quad().Cmp(z.r.Quad()) == 0
}

// Inv sets z equal to the inverse of y, and returns z. If y is zero, then Inv
// panics.
func (z *Float) Inv(y *Float) *Float {
	if y.IsZeroDivisor() {
		panic(zeroDivisorInverse)
	}
	return z.Divide(z.Conj(y), y.Quad())
}

// QuoL sets z equal to the left quotient of x and y:
// 		Mul(Inv(y), x)
// Then it returns z. If y is zero, then QuoL panics.
func (z *Float) QuoL(x, y *Float) *Float {
	if y.IsZeroDivisor() {
		panic(zeroDivisorDenominator)
	}
	return z.Mul(z.Inv(y), x)
}

// QuoR sets z equal to the right quotient of x and y:
// 		Mul(x, Inv(y))
// Then it returns z. If y is zero, then QuoR panics.
func (z *Float) QuoR(x, y *Float) *Float {
	if y.IsZeroDivisor() {
		panic(zeroDivisorDenominator)
	}
	return z.Mul(x, z.Inv(y))
}

// CrossRatioL sets z equal to the left cross-ratio of v, w, x, and y:
// 		Inv(w - x) * (v - x) * Inv(v - y) * (w - y)
// Then it returns z.
func (z *Float) CrossRatioL(v, w, x, y *Float) *Float {
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

// CrossRatioR sets z equal to the right cross-ratio of v, w, x, and y:
// 		(v - x) * Inv(w - x) * (w - y) * Inv(v - y)
// Then it returns z.
func (z *Float) CrossRatioR(v, w, x, y *Float) *Float {
	temp := new(Float)
	z.Sub(v, x)
	temp.Sub(w, x)
	temp.Inv(temp)
	z.Mul(z, temp)
	temp.Sub(w, y)
	z.Mul(z, temp)
	temp.Sub(v, y)
	temp.Inv(temp)
	return z.Mul(z, temp)
}

// MöbiusL sets z equal to the left Möbius (fractional linear) transform of y:
// 		Inv(y*c + d) * (y*a + b)
// Then it returns z.
func (z *Float) MöbiusL(y, a, b, c, d *Float) *Float {
	z.Mul(y, a)
	z.Add(z, b)
	temp := new(Float)
	temp.Mul(y, c)
	temp.Add(temp, d)
	temp.Inv(temp)
	return z.Mul(temp, z)
}

// MöbiusR sets z equal to the right Möbius (fractional linear) transform of y:
// 		(a*y + b) * Inv(c*y + d)
// Then it returns z.
func (z *Float) MöbiusR(y, a, b, c, d *Float) *Float {
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
		*cplex.NewFloat(
			big.NewFloat(rand.Float64()),
			big.NewFloat(rand.Float64()),
		),
		*cplex.NewFloat(
			big.NewFloat(rand.Float64()),
			big.NewFloat(rand.Float64()),
		),
	}
	return reflect.ValueOf(randomFloat)
}
