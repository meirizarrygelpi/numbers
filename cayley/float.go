// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package cayley

import (
	"math/big"
	"math/rand"
	"reflect"
	"strings"

	"github.com/meirizarrygelpi/numbers/cplex"
	"github.com/meirizarrygelpi/numbers/hamilton"
	"github.com/meirizarrygelpi/numbers/vec7"
)

// A Float is a Cayley octonion with big.Float components.
type Float struct {
	l, r hamilton.Float
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

// String returns the string version of a Float value.
//
// If z corresponds to a+bi+cj+dk+em+fn+gp+hq, then the string is
// "⦗a+bi+cj+dk+em+fn+gp+hq⦘", similar to complex128 values.
func (z *Float) String() string {
	v := z.Unreal()
	a := make([]string, 17)
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
	if v[3].Sign() < 0 {
		a[8] = v[3].String()
	} else {
		a[8] = "+" + v[3].String()
	}
	a[9] = unit4
	if v[4].Sign() < 0 {
		a[10] = v[4].String()
	} else {
		a[10] = "+" + v[4].String()
	}
	a[11] = unit5
	if v[5].Sign() < 0 {
		a[12] = v[5].String()
	} else {
		a[12] = "+" + v[5].String()
	}
	a[13] = unit6
	if v[6].Sign() < 0 {
		a[14] = v[6].String()
	} else {
		a[14] = "+" + v[6].String()
	}
	a[15] = unit3
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

// SetPair sets z equal to a Cayley octonion made with a given pair, and
// then it returns z.
func (z *Float) SetPair(a, b *hamilton.Float) *Float {
	z.l.Set(a)
	z.r.Set(b)
	return z
}

// NewFloat returns a pointer to the Float value a+bi+cj+dk+em+fn+gp+hq.
func NewFloat(a, b, c, d, e, f, g, h *big.Float) *Float {
	z := new(Float)
	z.l.SetPair(
		cplex.NewFloat(a, b),
		cplex.NewFloat(c, d),
	)
	z.r.SetPair(
		cplex.NewFloat(e, f),
		cplex.NewFloat(g, h),
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
//     +-----+----+----+----+----+----+----+----+
//     | Mul | i  | j  | k  | m  | n  | p  | q  |
//     +-----+----+----+----+----+----+----+----+
//     | i   | -1 | +k | -j | +n | -m | -q | +p |
//     +-----+----+----+----+----+----+----+----+
//     | j   | -k | -1 | +i | +p | +q | -m | -n |
//     +-----+----+----+----+----+----+----+----+
//     | k   | +j | -i | -1 | +q | -p | +n | -m |
//     +-----+----+----+----+----+----+----+----+
//     | m   | -n | -p | -q | -1 | +i | +j | +k |
//     +-----+----+----+----+----+----+----+----+
//     | n   | +m | -q | +p | -i | -1 | -k | +j |
//     +-----+----+----+----+----+----+----+----+
//     | p   | +q | +m | -n | -j | +k | -1 | -i |
//     +-----+----+----+----+----+----+----+----+
//     | q   | -p | +n | +m | -k | -j | +i | -1 |
//     +-----+----+----+----+----+----+----+----+
// This binary operation is non-commutative and non-associative.
func (z *Float) Mul(x, y *Float) *Float {
	a, b, temp := new(hamilton.Float), new(hamilton.Float), new(hamilton.Float)
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
func (z *Float) Commutator(x, y *Float) *Float {
	return z.Sub(
		z.Mul(x, y),
		new(Float).Mul(y, x),
	)
}

// Associator sets z equal to the associator of w, x, and y:
// 		Mul(Mul(w, x), y) - Mul(w, Mul(x, y))
// Then it returns z.
func (z *Float) Associator(w, x, y *Float) *Float {
	temp := new(Float)
	return z.Sub(
		z.Mul(z.Mul(w, x), y),
		temp.Mul(w, temp.Mul(x, y)),
	)
}

// Quad returns the quadrance of z. If z = a+bi+cj+dk+em+fn+gp+hq, then the
// quadrance is
// 		a² + b² + c² + d² + e² + f² + g² + h²
// This is always non-negative.
func (z *Float) Quad() *big.Float {
	return new(big.Float).Add(
		z.l.Quad(),
		z.r.Quad(),
	)
}

// Inv sets z equal to the inverse of y, and returns z. If y is zero, then Inv
// panics.
func (z *Float) Inv(y *Float) *Float {
	if zero := new(Float); y.Equals(zero) {
		panic(zeroInverse)
	}
	return z.Divide(z.Conj(y), y.Quad())
}

// QuoL sets z equal to the left quotient of x and y:
// 		Mul(Inv(y), x)
// Then it returns z. If y is zero, then QuoL panics.
func (z *Float) QuoL(x, y *Float) *Float {
	if zero := new(Float); y.Equals(zero) {
		panic(zeroDenominator)
	}
	return z.Mul(z.Inv(y), x)
}

// QuoR sets z equal to the right quotient of x and y:
// 		Mul(x, Inv(y))
// Then it returns z. If y is zero, then QuoR panics.
func (z *Float) QuoR(x, y *Float) *Float {
	if zero := new(Float); y.Equals(zero) {
		panic(zeroDenominator)
	}
	return z.Mul(x, z.Inv(y))
}

// Graves sets z equal to the Gravesian integer a+bi+cj+dk+em+fn+gp+hq, and
// returns z.
func (z *Float) Graves(a, b, c, d, e, f, g, h *big.Int) *Float {
	z.l.Lipschitz(a, b, c, d)
	z.r.Lipschitz(e, f, g, h)
	return z
}

// Klein sets z equal to the Kleinian integer
// (a+½)+(b+½)i+(c+½)j+(d+½)k+(e+½)m+(f+½)n+(g+½)p+(h+½)q, and returns z.
func (z *Float) Klein(a, b, c, d, e, f, g, h *big.Int) *Float {
	z.Graves(a, b, c, d, e, f, g, h)
	half := big.NewFloat(0.5)
	return z.Add(z, NewFloat(half, half, half, half, half, half, half, half))
}

// Generate returns a random Float value for quick.Check testing.
func (z *Float) Generate(rand *rand.Rand, size int) reflect.Value {
	randomFloat := &Float{
		*hamilton.NewFloat(
			big.NewFloat(rand.Float64()),
			big.NewFloat(rand.Float64()),
			big.NewFloat(rand.Float64()),
			big.NewFloat(rand.Float64()),
		),
		*hamilton.NewFloat(
			big.NewFloat(rand.Float64()),
			big.NewFloat(rand.Float64()),
			big.NewFloat(rand.Float64()),
			big.NewFloat(rand.Float64()),
		),
	}
	return reflect.ValueOf(randomFloat)
}
