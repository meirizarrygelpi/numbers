// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licensed under the MIT License.

package cplex

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"strings"

	"github.com/meirizarrygelpi/numbers/maclaurin"
	"github.com/meirizarrygelpi/numbers/pade"
)

// A Float64 is a complex number with float64 components.
type Float64 struct {
	l, r float64
}

// One sets z equal to 1, and then it returns z.
func (z *Float64) One() *Float64 {
	z.l = 1
	z.r = 0
	return z
}

// Inf sets z equal to a complex infinity, and then it returns z.
func (z *Float64) Inf(s1, s2 int) *Float64 {
	return z.SetPair(math.Inf(s1), math.Inf(s2))
}

// IsInf returns true if z is an infinity.
func (z *Float64) IsInf(s1, s2 int) bool {
	return math.IsInf(z.Real(), s1) && math.IsInf(z.Unreal(), s2)
}

// HasInf returns true if z has an infinite component.
func (z *Float64) HasInf() bool {
	return math.IsInf(z.Real(), 0) || math.IsInf(z.Unreal(), 0)
}

// NaN sets z equal to a complex NaN, and then it returns z.
func (z *Float64) NaN() *Float64 {
	nan := math.NaN()
	return z.SetPair(nan, nan)
}

// IsNaN returns true if a component of z is NaN and none is infinite.
func (z *Float64) IsNaN() bool {
	if math.IsInf(z.Real(), 0) || math.IsInf(z.Unreal(), 0) {
		return false
	}
	if math.IsNaN(z.Real()) || math.IsNaN(z.Unreal()) {
		return true
	}
	return false
}

// Real returns the real part of z.
func (z *Float64) Real() float64 {
	return z.l
}

// Unreal returns the unreal part of z.
func (z *Float64) Unreal() float64 {
	return z.r
}

func sprintFloat64(a float64) string {
	if math.IsNaN(a) {
		return "+NaN"
	}
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
// If z corresponds to a + bi, then the string is "(a+bi)", similar to
// complex128 values.
func (z *Float64) String() string {
	a := make([]string, 5)
	a[0] = leftBracket
	a[1] = fmt.Sprintf("%g", z.l)
	a[2] = sprintFloat64(z.r)
	a[3] = unitName
	a[4] = rightBracket
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Float64) Equals(y *Float64) bool {
	return z.l == y.l && z.r == y.r
}

// Set sets z equal to y, and returns z.
func (z *Float64) Set(y *Float64) *Float64 {
	z.l = y.l
	z.r = y.r
	return z
}

// SetPair sets z equal to a complex number made with a given pair, and then
// it returns z.
func (z *Float64) SetPair(a, b float64) *Float64 {
	z.l = a
	z.r = b
	return z
}

// SetReal sets the real part of z equal to a, and then it returns z.
func (z *Float64) SetReal(a float64) *Float64 {
	z.l = a
	return z
}

// SetUnreal sets the unreal part of z equal to b, and then it returns z.
func (z *Float64) SetUnreal(b float64) *Float64 {
	z.r = b
	return z
}

// NewFloat64 returns a pointer to the Float64 value a+bi.
func NewFloat64(a, b float64) *Float64 {
	z := new(Float64)
	z.SetPair(a, b)
	return z
}

// Plus sets z equal to y+a, with a real, and returns z.
func (z *Float64) Plus(y *Float64, a float64) *Float64 {
	z.l = y.l + a
	z.r = y.r
	return z
}

// Minus sets z equal to y-a, with a real, and returns z.
func (z *Float64) Minus(y *Float64, a float64) *Float64 {
	z.l = y.l - a
	z.r = y.r
	return z
}

// Dilate sets z equal to y dilated by a, and returns z.
func (z *Float64) Dilate(y *Float64, a float64) *Float64 {
	z.l = y.l * a
	z.r = y.r * a
	return z
}

// Divide sets z equal to y contracted by a, and returns z.
func (z *Float64) Divide(y *Float64, a float64) *Float64 {
	z.l = y.l / a
	z.r = y.r / a
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Float64) Neg(y *Float64) *Float64 {
	z.l = -y.l
	z.r = -y.r
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Float64) Conj(y *Float64) *Float64 {
	z.l = y.l
	z.r = -y.r
	return z
}

// Add sets z equal to x+y, and returns z.
func (z *Float64) Add(x, y *Float64) *Float64 {
	z.l = x.l + y.l
	z.r = x.r + y.r
	return z
}

// Sub sets z equal to x-y, and returns z.
func (z *Float64) Sub(x, y *Float64) *Float64 {
	z.l = x.l - y.l
	z.r = x.r - y.r
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
func (z *Float64) Mul(x, y *Float64) *Float64 {
	a := (x.l * y.l) - (y.r * x.r)
	b := (x.r * y.l) + (y.r * x.l)
	z.SetPair(a, b)
	return z
}

// Dot returns the dot product of y and z. If z = a+bi and y = c+di, then the
// dot product is
// 		ac + bd
// This can be positive, negative, or zero. The dot product is equivalent to
// 		½(Mul(Conj(z), y) + Mu(Conj(y), z))
// In this form it is clear that Dot is symmetric.
func (z *Float64) Dot(y *Float64) float64 {
	return (z.l * y.l) + (z.r * y.r)
}

// Quad returns the quadrance of z. If z = a+bi, then the quadrance is
// 		a² + b²
// This is always non-negative.
func (z *Float64) Quad() float64 {
	return z.Dot(z)
}

// Cross returns the cross product of y and z. If z = a+bi and y = c+di, then
// the cross product is
// 		ad - bc
// This can be positive, negative, or zero. The cross product is equivalent to
// the unreal part of
// 		½(Mul(Conj(z), y) - Mu(Conj(y), z))
// In this form it is clear that Cross is anti-symmetric.
func (z *Float64) Cross(y *Float64) float64 {
	return (z.l * y.r) - (z.r * y.l)
}

// Lozenge sets z equal to the lozenge product of v, w, x, and y:
// 		Mul(v, Conj(w)) - Mul(x, Conj(y))
// Then it returns z.
func (z *Float64) Lozenge(v, w, x, y *Float64) *Float64 {
	a, b := new(Float64), new(Float64)
	a.Mul(v, a.Conj(w))
	b.Mul(x, b.Conj(y))
	return z.Sub(a, b)
}

// Inv sets z equal to the inverse of y, and returns z. If y is zero, then Inv
// panics.
func (z *Float64) Inv(y *Float64) *Float64 {
	if zero := new(Float64); y.Equals(zero) {
		panic(zeroInverse)
	}
	return z.Divide(z.Conj(y), y.Quad())
}

// Quo sets z equal to the quotient of x and y, and returns z. If y is zero,
// then Quo panics. Quo uses a naive division algorithm that can fail.
func (z *Float64) Quo(x, y *Float64) *Float64 {
	if zero := new(Float64); y.Equals(zero) {
		panic(zeroDenominator)
	}
	q := y.Quad()
	a := (x.l * y.l) + (y.r * x.r)
	b := (x.r * y.l) - (y.r * x.l)
	z.SetPair(a, b)
	return z.Divide(z, q)
}

var (
	ep = math.Nextafter(1.0, 2.0) - 1.0
	ov = math.MaxFloat64
	un = math.SmallestNonzeroFloat64 * math.Pow(2, 52)
)

func maxabs(a, b float64) float64 {
	return math.Max(math.Abs(a), math.Abs(b))
}

func f2(a, b, c, d, r, t float64) float64 {
	if r != 0 {
		br := b * r
		if br != 0 {
			return (a + br) * t
		}
		return (a * t) + (b*t)*r
	}
	return (a + d*(b/c)) * t
}

func f1(a, b, c, d float64) (e, f float64) {
	r := d / c
	t := 1 / (c + d*r)
	e = f2(a, b, c, d, r, t)
	f = f2(b, -a, c, d, r, t)
	return
}

func (z *Float64) robustQuo(x, y *Float64) *Float64 {
	a, b := x.Real(), x.Unreal()
	c, d := y.Real(), y.Unreal()

	if math.Abs(d) <= math.Abs(c) {
		e, f := f1(a, b, c, d)
		return z.SetPair(e, f)
	}

	e, f := f1(b, a, d, c)
	f = -f
	return z.SetPair(e, f)
}

// RobustQuo sets z equal to the quotient of x and y, and returns z. If y is
// zero, then RobustQuo panics. RobustQuo uses a more robust algorithm for
// complex division found in
//     M. Baudin and R.L. Smith, A Robust Complex Division in Scilab
//     (arXiv:1210.4539)
func (z *Float64) RobustQuo(x, y *Float64) *Float64 {
	if zero := new(Float64); y.Equals(zero) {
		panic(zeroDenominator)
	}
	p, q := new(Float64), new(Float64)

	p.Set(x)
	q.Set(y)

	a, b := x.Real(), x.Unreal()
	c, d := y.Real(), y.Unreal()

	ab := maxabs(a, b)
	cd := maxabs(c, d)

	base := 2.0
	s := 1.0

	be := base / (ep * ep)

	if ab >= ov/2.0 {
		p.Divide(p, 2.0)
		s = s * 2.0
	}
	if cd >= ov/2.0 {
		q.Divide(q, 2.0)
		s = s / 2.0
	}

	if ab <= (un*base)/ep {
		p.Dilate(p, be)
		s = s / be
	}
	if cd <= (un*base)/ep {
		q.Dilate(q, be)
		s = s * be
	}

	z.robustQuo(p, q)
	return z.Dilate(z, s)
}

// Gauss sets z equal to the Gaussian integer a+bi, and returns z.
func (z *Float64) Gauss(a, b int64) *Float64 {
	z.l = float64(a)
	z.r = float64(b)
	return z
}

// CrossRatio sets z equal to the cross-ratio of v, w, x, and y:
// 		Inv(w - x) * (v - x) * Inv(v - y) * (w - y)
// Then it returns z.
func (z *Float64) CrossRatio(v, w, x, y *Float64) *Float64 {
	temp := new(Float64)
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

// Mobius sets z equal to the Möbius (fractional linear) transform of y:
// 		(a*y + b) * Inv(c*y + d)
// Then it returns z.
func (z *Float64) Mobius(y, a, b, c, d *Float64) *Float64 {
	z.Mul(a, y)
	z.Add(z, b)
	temp := new(Float64)
	temp.Mul(c, y)
	temp.Add(temp, d)
	temp.Inv(temp)
	return z.Mul(z, temp)
}

// Maclaurin sets z equal to the value of the Maclaurin polynomial p evaluated
// at y, and returns z. Horner's method is used.
func (z *Float64) Maclaurin(y *Float64, p *maclaurin.Float64) *Float64 {
	if p.Len() == 0 {
		z = new(Float64)
		return z
	}
	n := p.Degree
	var a float64
	if n == 0 {
		z = new(Float64)
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

// Pade sets z equal to the value of the Padé approximant r evaluated at y,
// and returns z.
func (z *Float64) Pade(y *Float64, r *pade.Float64) *Float64 {
	p, q := new(Float64), new(Float64)
	p.Maclaurin(y, &r.P)
	q.Maclaurin(y, &r.Q)
	return z.Quo(p, q)
}

// Generate returns a random Float64 value for quick.Check testing.
func (z *Float64) Generate(rand *rand.Rand, size int) reflect.Value {
	randomFloat64 := &Float64{
		rand.Float64(),
		rand.Float64(),
	}
	return reflect.ValueOf(randomFloat64)
}
