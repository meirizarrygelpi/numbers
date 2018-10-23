// Copyright (c) 2017 Melvin Eloy Irizarry-Gelpí
// Licensed under the MIT License.

package percockle

import (
	"math/big"
	"math/rand"
	"reflect"
	"strings"

	"github.com/meirizarrygelpi/numbers/cockle"
	"github.com/meirizarrygelpi/numbers/cplex"
	"github.com/meirizarrygelpi/numbers/pplex"
	"github.com/meirizarrygelpi/numbers/vec7"
)

// A Rat is an perplex-Cockle quaternion with big.Rat components.
type Rat struct {
	l, r cockle.Rat
}

// One sets z equal to 1, and then returns z.
func (z *Rat) One() *Rat {
	z.l.One()
	z.r.Set(new(cockle.Rat))
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
// If z corresponds to a+bi+ct+du+eS+fiS+gtS+huS, then the string is
// "(a+bi+ct+du+eS+fiS+gtS+huS)", similar to complex128 values.
func (z *Rat) String() string {
	v := z.Unreal()
	a := make([]string, 17)
	a[0] = leftBracket
	a[1] = z.l.Real().RatString()
	i := 2
	for j, u := range unitNames {
		if v[j].Sign() < 0 {
			a[i] = v[j].RatString()
		} else {
			a[i] = "+" + v[j].RatString()
		}
		a[i+1] = u
		i += 2
	}
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

// SetPair sets z equal to an perplex-Cockle quaternion made with a given pair, and
// then it returns z.
func (z *Rat) SetPair(a, b *cockle.Rat) *Rat {
	z.l.Set(a)
	z.r.Set(b)
	return z
}

// NewRat returns a pointer to the Rat value a+bi+ct+du+eS+fiS+gtS+huS.
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

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Rat) Conj(y *Rat) *Rat {
	z.l.Conj(&y.l)
	z.r.Neg(&y.r)
	return z
}

// Bar sets z equal to the Cockle conjugate of y, and returns z.
func (z *Rat) Bar(y *Rat) *Rat {
	z.l.Conj(&y.l)
	z.r.Conj(&y.r)
	return z
}

// Tilde sets z equal to the S-conjugate of y, and returns z.
func (z *Rat) Tilde(y *Rat) *Rat {
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
func (z *Rat) Mul(x, y *Rat) *Rat {
	a, b, temp := new(cockle.Rat), new(cockle.Rat), new(cockle.Rat)
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

// Commutator sets z equal to the commutator of x and y:
// 		Mul(x, y) - Mul(y, x)
// Then it returns z.
func (z *Rat) Commutator(x, y *Rat) *Rat {
	return z.Sub(
		z.Mul(x, y),
		new(Rat).Mul(y, x),
	)
}

// Quad returns the quadrance of z. This is a perplex number.
func (z *Rat) Quad() *pplex.Rat {
	q := new(cockle.Rat)
	quad := z.l.Quad()
	q.Mul(&z.l, q.Conj(&z.r))
	q.Scale(q, big.NewRat(2, 1))
	return pplex.NewRat(quad.Add(quad, z.r.Quad()), q.Real())
}

// Norm returns the norm of z. This can be positive, negative, or zero.
func (z *Rat) Norm() *big.Rat {
	return z.Quad().Quad()
}

// IsZeroDivisor returns true if z is a zero divisor.
func (z *Rat) IsZeroDivisor() bool {
	zero := new(cockle.Rat)
	return z.l.Equals(zero)
}

// Inv sets z equal to the inverse of y, and returns z. If y is a zero divisor,
// then Inv panics.
func (z *Rat) Inv(y *Rat) *Rat {
	if y.IsZeroDivisor() {
		panic(zeroDivisorInverse)
	}
	q := y.Quad()
	q.Inv(q)
	zero := new(big.Rat)
	return z.Mul(
		NewRat(q.Real(), zero, zero, zero,
			q.Unreal(), zero, zero, zero),
		new(Rat).Bar(y),
	)
}

// QuoL sets z equal to the left quotient of x and y:
// 		Mul(Inv(y), x)
// Then it returns z. If y is a zero divisor, then QuoL panics.
func (z *Rat) QuoL(x, y *Rat) *Rat {
	if y.IsZeroDivisor() {
		panic(zeroDivisorDenominator)
	}
	return z.Mul(z.Inv(y), x)
}

// QuoR sets z equal to the right quotient of x and y:
// 		Mul(x, Inv(y))
// Then it returns z. If y is a zero divisor, then QuoR panics.
func (z *Rat) QuoR(x, y *Rat) *Rat {
	if y.IsZeroDivisor() {
		panic(zeroDivisorDenominator)
	}
	return z.Mul(x, z.Inv(y))
}

// CrossRatioL sets z equal to the left cross-ratio of v, w, x, and y:
// 		Inv(w - x) * (v - x) * Inv(v - y) * (w - y)
// Then it returns z.
func (z *Rat) CrossRatioL(v, w, x, y *Rat) *Rat {
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

// CrossRatioR sets z equal to the right cross-ratio of v, w, x, and y:
// 		(v - x) * Inv(w - x) * (w - y) * Inv(v - y)
// Then it returns z.
func (z *Rat) CrossRatioR(v, w, x, y *Rat) *Rat {
	temp := new(Rat)
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

// MobiusL sets z equal to the left Möbius (fractional linear) transform of y:
// 		Inv(y*c + d) * (y*a + b)
// Then it returns z.
func (z *Rat) MobiusL(y, a, b, c, d *Rat) *Rat {
	z.Mul(y, a)
	z.Add(z, b)
	temp := new(Rat)
	temp.Mul(y, c)
	temp.Add(temp, d)
	temp.Inv(temp)
	return z.Mul(temp, z)
}

// MobiusR sets z equal to the right Möbius (fractional linear) transform of y:
// 		(a*y + b) * Inv(c*y + d)
// Then it returns z.
func (z *Rat) MobiusR(y, a, b, c, d *Rat) *Rat {
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
		*cockle.NewRat(
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
		),
		*cockle.NewRat(
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
		),
	}
	return reflect.ValueOf(randomRat)
}
