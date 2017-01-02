// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package ultra

import (
	"math/big"
	"math/rand"
	"reflect"
	"strings"

	"github.com/meirizarrygelpi/numbers/nplex"
	"github.com/meirizarrygelpi/numbers/supra"
	"github.com/meirizarrygelpi/numbers/vec7"
)

// A Rat is an ultra number with big.Rat components.
type Rat struct {
	l, r supra.Rat
}

// One sets z equal to 1, and then returns z.
func (z *Rat) One() *Rat {
	z.l.One()
	z.r.Set(new(supra.Rat))
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
// If z corresponds to a+bα+cβ+dγ+eδ+fε+gζ+hη, then the string is
// "⦗a+bα+cβ+dγ+eδ+fε+gζ+hη⦘", similar to complex128 values.
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

// SetPair sets z equal to an ultra number made with a given pair, and
// then it returns z.
func (z *Rat) SetPair(a, b *supra.Rat) *Rat {
	z.l.Set(a)
	z.r.Set(b)
	return z
}

// NewRat returns a pointer to the Rat value a+bα+cβ+dγ+eδ+fε+gζ+hη.
func NewRat(a, b, c, d, e, f, g, h *big.Rat) *Rat {
	z := new(Rat)
	z.l.SetPair(
		nplex.NewRat(a, b),
		nplex.NewRat(c, d),
	)
	z.r.SetPair(
		nplex.NewRat(e, f),
		nplex.NewRat(g, h),
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
	a, b, temp := new(supra.Rat), new(supra.Rat), new(supra.Rat)
	a.Mul(&x.l, &y.l)
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
func (z *Rat) Commutator(x, y *Rat) *Rat {
	return z.Sub(
		z.Mul(x, y),
		new(Rat).Mul(y, x),
	)
}

// Associator sets z equal to the associator of w, x, and y:
// 		Mul(Mul(w, x), y) - Mul(w, Mul(x, y))
// Then it returns z.
func (z *Rat) Associator(w, x, y *Rat) *Rat {
	temp := new(Rat)
	return z.Sub(
		z.Mul(z.Mul(w, x), y),
		temp.Mul(w, temp.Mul(x, y)),
	)
}

// Quad returns the quadrance of z. If z = a+bα+cβ+dγ+eδ+fε+gζ+hη, then the
// quadrance is
// 		a²
// This is always non-negative.
func (z *Rat) Quad() *big.Rat {
	return z.l.Quad()
}

// IsZeroDivisor returns true if z is a zero divisor.
func (z *Rat) IsZeroDivisor() bool {
	return z.l.IsZeroDivisor()
}

// Inv sets z equal to the inverse of y, and returns z. If y is a zero divisor,
// then Inv panics.
func (z *Rat) Inv(y *Rat) *Rat {
	if y.IsZeroDivisor() {
		panic(zeroDivisorInverse)
	}
	a := y.Quad()
	a.Inv(a)
	return z.Scale(z.Conj(y), a)
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

// Generate returns a random Rat value for quick.Check testing.
func (z *Rat) Generate(rand *rand.Rand, size int) reflect.Value {
	randomRat := &Rat{
		*supra.NewRat(
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
		),
		*supra.NewRat(
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
		),
	}
	return reflect.ValueOf(randomRat)
}
