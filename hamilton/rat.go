// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package hamilton

import (
	"math/big"
	"math/rand"
	"reflect"
	"strings"

	"github.com/meirizarrygelpi/numbers/cplex"
	"github.com/meirizarrygelpi/numbers/vec3"
)

// A Rat is a Hamilton quaternion with big.Rat components.
type Rat struct {
	l, r cplex.Rat
}

// One sets z equal to 1, and then returns z.
func (z *Rat) One() *Rat {
	z.l.One()
	z.r.Set(new(cplex.Rat))
	return z
}

// Real returns the real part of z.
func (z *Rat) Real() *big.Rat {
	return z.l.Real()
}

// Unreal returns the unreal part of z.
func (z *Rat) Unreal() *vec3.Rat {
	v := new(vec3.Rat)
	v[0] = z.l.Unreal()
	v[1] = z.r.Real()
	v[2] = z.r.Unreal()
	return v
}

// String returns the string version of a Rat value.
//
// If z corresponds to a+bi+cj+dk, then the string is "⦗a+bi+cj+dk⦘", similar
// to complex128 values.
func (z *Rat) String() string {
	v := z.Unreal()
	a := make([]string, 9)
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
	a[8] = rightBracket
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

// SetPair sets z equal to a Hamilton quaternion made with a given pair, and
// then it returns z.
func (z *Rat) SetPair(a, b *cplex.Rat) *Rat {
	z.l.Set(a)
	z.r.Set(b)
	return z
}

// NewRat returns a pointer to the Rat value a+bi+cj+dk.
func NewRat(a, b, c, d *big.Rat) *Rat {
	z := new(Rat)
	z.l.SetPair(a, b)
	z.r.SetPair(c, d)
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
//
// The multiplication rule is:
// 		Mul(i, i) = Mul(j, j) = Mul(k, k) = -1
// 		Mul(i, j) = -Mul(j, i) = k
// 		Mul(j, k) = -Mul(k, j) = i
// 		Mul(k, i) = -Mul(i, k) = j
// This binary operation is non-commutative but associative.
func (z *Rat) Mul(x, y *Rat) *Rat {
	a, b, temp := new(cplex.Rat), new(cplex.Rat), new(cplex.Rat)
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
func (z *Rat) Commutator(x, y *Rat) *Rat {
	return z.Sub(
		z.Mul(x, y),
		new(Rat).Mul(y, x),
	)
}

// Quad returns the quadrance of z. If z = a+bi+cj+dk, then the quadrance is
// 		a² + b² + c² + d²
// This is always non-negative.
func (z *Rat) Quad() *big.Rat {
	return new(big.Rat).Add(
		z.l.Quad(),
		z.r.Quad(),
	)
}

// Inv sets z equal to the inverse of y, and returns z. If y is zero, then Inv
// panics.
func (z *Rat) Inv(y *Rat) *Rat {
	if zero := new(Rat); y.Equals(zero) {
		panic(zeroInverse)
	}
	a := y.Quad()
	a.Inv(a)
	return z.Scale(z.Conj(y), a)
}

// QuoL sets z equal to the left quotient of x and y:
// 		Mul(Inv(y), x)
// Then it returns z. If y is zero, then QuoL panics.
func (z *Rat) QuoL(x, y *Rat) *Rat {
	if zero := new(Rat); y.Equals(zero) {
		panic(zeroDenominator)
	}
	return z.Mul(z.Inv(y), x)
}

// QuoR sets z equal to the right quotient of x and y:
// 		Mul(x, Inv(y))
// Then it returns z. If y is zero, then QuoR panics.
func (z *Rat) QuoR(x, y *Rat) *Rat {
	if zero := new(Rat); y.Equals(zero) {
		panic(zeroDenominator)
	}
	return z.Mul(x, z.Inv(y))
}

// Lipschitz sets z equal to the Lipschitz integer a+bi+cj+dk, and returns z.
func (z *Rat) Lipschitz(a, b, c, d *big.Int) *Rat {
	z.l.Gauss(a, b)
	z.r.Gauss(c, d)
	return z
}

// Hurwitz sets z equal to the Hurwitz integer (a+½)+(b+½)i+(c+½)j+(d+½)k,
// and returns z.
func (z *Rat) Hurwitz(a, b, c, d *big.Int) *Rat {
	z.Lipschitz(a, b, c, d)
	half := big.NewRat(1, 2)
	return z.Add(z, NewRat(half, half, half, half))
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

// MöbiusL sets z equal to the left Möbius (fractional linear) transform of y:
// 		Inv(y*c + d) * (y*a + b)
// Then it returns z.
func (z *Rat) MöbiusL(y, a, b, c, d *Rat) *Rat {
	z.Mul(y, a)
	z.Add(z, b)
	temp := new(Rat)
	temp.Mul(y, c)
	temp.Add(temp, d)
	temp.Inv(temp)
	return z.Mul(temp, z)
}

// MöbiusR sets z equal to the right Möbius (fractional linear) transform of y:
// 		(a*y + b) * Inv(c*y + d)
// Then it returns z.
func (z *Rat) MöbiusR(y, a, b, c, d *Rat) *Rat {
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
		*cplex.NewRat(
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
		),
		*cplex.NewRat(
			big.NewRat(rand.Int63(), rand.Int63()),
			big.NewRat(rand.Int63(), rand.Int63()),
		),
	}
	return reflect.ValueOf(randomRat)
}
