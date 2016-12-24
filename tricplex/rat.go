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

// A Rat is a tri-complex number with big.Rat components.
type Rat struct {
	l, r bicplex.Rat
}

// One sets z equal to 1, and then returns z.
func (z *Rat) One() *Rat {
	z.l.One()
	z.r.Set(new(bicplex.Rat))
	return z
}

// Real returns the real part of z.
func (z *Rat) Real() *big.Rat {
	return z.l.Real()
}

// Unreal returns the unreal part of z.
func (z *Rat) Unreal() *vec7.Rat {
	v := new(vec7.Rat)
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

// String returns the string version of a Rat value.
//
// If z corresponds to a+bi+cJ+diJ+eK+fiK+gJK+hiJK, then the string is
// "⦗a+bi+cJ+diJ+eK+fiK+gJK+hiJK⦘", similar to complex128 values.
func (z *Rat) String() string {
	v := z.Unreal()
	a := make([]string, 17)
	a[0] = leftBracket
	a[1] = z.l.Real().RatString()
	if v[0].Sign() < 0 {
		a[2] = v[0].RatString()
	} else {
		a[2] = "+" + v[0].RatString()
	}
	a[3] = unit1
	if v[1].Sign() < 0 {
		a[4] = v[1].RatString()
	} else {
		a[4] = "+" + v[1].RatString()
	}
	a[5] = unit2
	if v[2].Sign() < 0 {
		a[6] = v[2].RatString()
	} else {
		a[6] = "+" + v[2].RatString()
	}
	a[7] = unit3
	if v[3].Sign() < 0 {
		a[8] = v[3].RatString()
	} else {
		a[8] = "+" + v[3].RatString()
	}
	a[9] = unit4
	if v[4].Sign() < 0 {
		a[10] = v[4].RatString()
	} else {
		a[10] = "+" + v[4].RatString()
	}
	a[11] = unit5
	if v[5].Sign() < 0 {
		a[12] = v[5].RatString()
	} else {
		a[12] = "+" + v[5].RatString()
	}
	a[13] = unit6
	if v[6].Sign() < 0 {
		a[14] = v[6].RatString()
	} else {
		a[14] = "+" + v[6].RatString()
	}
	a[15] = unit3
	a[16] = rightBracket
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

// SetPair sets z equal to a tri-complex number made with a given pair, and
// then it returns z.
func (z *Rat) SetPair(a, b *bicplex.Rat) *Rat {
	z.l.Set(a)
	z.r.Set(b)
	return z
}

// NewRat returns a pointer to the Rat value a+bi+cJ+diJ+eK+fiK+gJK+hiJK.
func NewRat(a, b, c, d, e, f, g, h *big.Rat) *Rat {
	z := new(Rat)
	z.l.SetPair(
		cplex.NewRat(a, b),
		cplex.NewRat(c, d),
	)
	z.r.SetPair(
		cplex.NewRat(e, f),
		cplex.NewRat(g, h),
	)
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

// Star1 sets z equal to the i-conjugate of y, and returns z.
func (z *Rat) Star1(y *Rat) *Rat {
	z.l.Star1(&y.l)
	z.r.Star1(&y.r)
	return z
}

// Star2 sets z equal to the J-conjugate of y, and returns z.
func (z *Rat) Star2(y *Rat) *Rat {
	z.l.Star2(&y.l)
	z.r.Star2(&y.r)
	return z
}

// Star3 sets z equal to the K-conjugate of y, and returns z.
func (z *Rat) Star3(y *Rat) *Rat {
	z.l.Set(&y.l)
	z.r.Neg(&y.r)
	return z
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
//
// The multiplication table is:
//     +-----+-----+-----+-----+-----+-----+-----+-----+
//     | Mul | i   | J   | iJ  | K   | iK  | JK  | iJK |
//     +-----+-----+-----+-----+-----+-----+-----+-----+
//     | i   | -1  | iJ  | -J  | iK  | -K  | iJK | -JK |
//     +-----+-----+-----+-----+-----+-----+-----+-----+
//     | J   | iJ  | -1  | -i  | JK  | iJK | -K  | -iK |
//     +-----+-----+-----+-----+-----+-----+-----+-----+
//     | iJ  | -J  | -i  | +1  | iJK | -JK | -iK | +K  |
//     +-----+-----+-----+-----+-----+-----+-----+-----+
//     | K   | iK  | JK  | iJK | -1  | -i  | -J  | -iJ |
//     +-----+-----+-----+-----+-----+-----+-----+-----+
//     | iK  | -K  | iJK | -JK | -i  | +1  | -iJ | +J  |
//     +-----+-----+-----+-----+-----+-----+-----+-----+
//     | JK  | iJK | -K  | -iK | -J  | -iK | +1  | +i  |
//     +-----+-----+-----+-----+-----+-----+-----+-----+
//     | iJK | -JK | -iK | +K  | -iJ | +J  | +i  | -1  |
//     +-----+-----+-----+-----+-----+-----+-----+-----+
// This binary operation is commutative and associative.
func (z *Rat) Mul(x, y *Rat) *Rat {
	a, b, temp := new(bicplex.Rat), new(bicplex.Rat), new(bicplex.Rat)
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
// 		...
// Note that this is a bi-complex number.
func (z *Rat) Quad() *bicplex.Rat {
	q := new(bicplex.Rat)
	return q.Add(q.Mul(&z.l, &z.l), new(bicplex.Rat).Mul(&z.r, &z.r))
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
func (z *Rat) Norm() *big.Rat {
	return z.Quad().Norm()
}

// IsZeroDivisor returns true if z is a zero divisor.
func (z *Rat) IsZeroDivisor() bool {
	zero := new(bicplex.Rat)
	return z.Quad().Equals(zero)
}

// Inv sets z equal to the inverse of y, and returns z. If y is zero, then Inv
// panics.
func (z *Rat) Inv(y *Rat) *Rat {
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
func (z *Rat) Quo(x, y *Rat) *Rat {
	if y.IsZeroDivisor() {
		panic(zeroDivisorDenominator)
	}
	return z.Mul(x, z.Inv(y))
}

// CrossRatio sets z equal to the cross-ratio of v, w, x, and y:
// 		Inv(w - x) * (v - x) * Inv(v - y) * (w - y)
// Then it returns z.
func (z *Rat) CrossRatio(v, w, x, y *Rat) *Rat {
	temp := new(Rat)
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
func (z *Rat) Möbius(y, a, b, c, d *Rat) *Rat {
	z.Mul(a, y)
	z.Add(z, b)
	temp := new(Rat)
	temp.Mul(c, y)
	temp.Add(temp, d)
	temp.Inv(temp)
	return z.Mul(z, temp)
}

// Generate returns a random Rat value for quick.Check testing.
func (z *Rat) Generate(rand *rand.Rand, size int) reflect.Value {
	randomRat := &Rat{
		*bicplex.NewRat(
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
		),
		*bicplex.NewRat(
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
		),
	}
	return reflect.ValueOf(randomRat)
}