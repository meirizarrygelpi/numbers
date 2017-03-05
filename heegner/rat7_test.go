// Copyright (c) 2017 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package heegner

import (
	"math/big"
	"testing"
	"testing/quick"
)

// Commutativity

func TestAddCommutativeRat7(t *testing.T) {
	f := func(x, y *Rat7) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(Rat7).Add(x, y)
		r := new(Rat7).Add(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulCommutativeRat7(t *testing.T) {
	f := func(x, y *Rat7) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(Rat7).Mul(x, y)
		r := new(Rat7).Mul(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestNegConjCommutativeRat7(t *testing.T) {
	f := func(x *Rat7) bool {
		// t.Logf("x = %v", x)
		l, r := new(Rat7), new(Rat7)
		l.Neg(l.Conj(x))
		r.Conj(r.Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-commutativity

func TestSubAntiCommutativeRat7(t *testing.T) {
	f := func(x, y *Rat7) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat7), new(Rat7)
		l.Sub(x, y)
		r.Sub(y, x)
		r.Neg(r)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Associativity

func TestAddAssociativeRat7(t *testing.T) {
	f := func(x, y, z *Rat7) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Rat7), new(Rat7)
		l.Add(l.Add(x, y), z)
		r.Add(x, r.Add(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulAssociativeRat7(t *testing.T) {
	f := func(x, y, z *Rat7) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Rat7), new(Rat7)
		l.Mul(l.Mul(x, y), z)
		r.Mul(x, r.Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Identity

func TestAddZeroRat7(t *testing.T) {
	zero := new(Rat7)
	f := func(x *Rat7) bool {
		// t.Logf("x = %v", x)
		l := new(Rat7).Add(x, zero)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulOneRat7(t *testing.T) {
	one := &Rat7{
		l: *big.NewRat(1, 1),
	}
	f := func(x *Rat7) bool {
		// t.Logf("x = %v", x)
		l := new(Rat7).Mul(x, one)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulInvOneRat7(t *testing.T) {
	one := &Rat7{
		l: *big.NewRat(1, 1),
	}
	f := func(x *Rat7) bool {
		// t.Logf("x = %v", x)
		l := new(Rat7)
		l.Mul(x, l.Inv(x))
		return l.Equals(one)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddNegSubRat7(t *testing.T) {
	f := func(x, y *Rat7) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat7), new(Rat7)
		l.Sub(x, y)
		r.Add(x, r.Neg(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddScaleDoubleRat7(t *testing.T) {
	f := func(x *Rat7) bool {
		// t.Logf("x = %v", x)
		l, r := new(Rat7), new(Rat7)
		l.Add(x, x)
		r.Scale(x, big.NewRat(2, 1))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Involutivity

func TestInvInvolutiveRat7(t *testing.T) {
	f := func(x *Rat7) bool {
		// t.Logf("x = %v", x)
		l := new(Rat7)
		l.Inv(l.Inv(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestNegInvolutiveRat7(t *testing.T) {
	f := func(x *Rat7) bool {
		// t.Logf("x = %v", x)
		l := new(Rat7)
		l.Neg(l.Neg(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestConjInvolutiveRat7(t *testing.T) {
	f := func(x *Rat7) bool {
		// t.Logf("x = %v", x)
		l := new(Rat7)
		l.Conj(l.Conj(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-distributivity

func TestMulConjAntiDistributiveRat7(t *testing.T) {
	f := func(x, y *Rat7) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat7), new(Rat7)
		l.Conj(l.Mul(x, y))
		r.Mul(r.Conj(y), new(Rat7).Conj(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulInvAntiDistributiveRat7(t *testing.T) {
	f := func(x, y *Rat7) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat7), new(Rat7)
		l.Inv(l.Mul(x, y))
		r.Mul(r.Inv(y), new(Rat7).Inv(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Distributivity

func TestAddConjDistributiveRat7(t *testing.T) {
	f := func(x, y *Rat7) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat7), new(Rat7)
		l.Add(x, y)
		l.Conj(l)
		r.Add(r.Conj(x), new(Rat7).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubConjDistributiveRat7(t *testing.T) {
	f := func(x, y *Rat7) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat7), new(Rat7)
		l.Sub(x, y)
		l.Conj(l)
		r.Sub(r.Conj(x), new(Rat7).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddScaleDistributiveRat7(t *testing.T) {
	f := func(x, y *Rat7) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(Rat7), new(Rat7)
		l.Scale(l.Add(x, y), a)
		r.Add(r.Scale(x, a), new(Rat7).Scale(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubScaleDistributiveRat7(t *testing.T) {
	f := func(x, y *Rat7) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(Rat7), new(Rat7)
		l.Scale(l.Sub(x, y), a)
		r.Sub(r.Scale(x, a), new(Rat7).Scale(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddMulDistributiveRat7(t *testing.T) {
	f := func(x, y, z *Rat7) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Rat7), new(Rat7)
		l.Mul(l.Add(x, y), z)
		r.Add(r.Mul(x, z), new(Rat7).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubMulDistributiveRat7(t *testing.T) {
	f := func(x, y, z *Rat7) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Rat7), new(Rat7)
		l.Mul(l.Sub(x, y), z)
		r.Sub(r.Mul(x, z), new(Rat7).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Positivity

func TestQuadPositiveRat7(t *testing.T) {
	f := func(x *Rat7) bool {
		// t.Logf("x = %v", x)
		return x.Quad().Sign() > 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Composition

func TestCompositionRat7(t *testing.T) {
	f := func(x, y *Rat7) bool {
		// t.Logf("x = %v, y = %v", x, y)
		p := new(Rat7)
		a, b := new(big.Rat), new(big.Rat)
		p.Mul(x, y)
		a.Set(p.Quad())
		b.Mul(x.Quad(), y.Quad())
		return a.Cmp(b) == 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Quotient

func TestQuotientsRat7(t *testing.T) {
	f := func(x, y *Rat7) bool {
		// t.Logf("x = %v, y = %v", x, y)
		z, p, q := new(Rat7), new(Rat7), new(Rat7)
		z.Mul(x, y)
		p.Quo(z, y)
		q.Quo(z, x)
		return p.Equals(x) && q.Equals(y)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
