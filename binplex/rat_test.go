// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licensed under the MIT License.

package binplex

import (
	"math/big"
	"testing"
	"testing/quick"
)

// Commutativity

func TestAddCommutativeRat(t *testing.T) {
	f := func(x, y *Rat) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(Rat).Add(x, y)
		r := new(Rat).Add(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestNegBarCommutativeRat(t *testing.T) {
	f := func(x *Rat) bool {
		// t.Logf("x = %v", x)
		l, r := new(Rat), new(Rat)
		l.Neg(l.Bar(x))
		r.Bar(r.Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestNegTildeCommutativeRat(t *testing.T) {
	f := func(x *Rat) bool {
		// t.Logf("x = %v", x)
		l, r := new(Rat), new(Rat)
		l.Neg(l.Tilde(x))
		r.Tilde(r.Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-commutativity

func TestSubAntiCommutativeRat(t *testing.T) {
	f := func(x, y *Rat) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat), new(Rat)
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

func TestAddAssociativeRat(t *testing.T) {
	f := func(x, y, z *Rat) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Rat), new(Rat)
		l.Add(l.Add(x, y), z)
		r.Add(x, r.Add(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulAssociativeRat(t *testing.T) {
	f := func(x, y, z *Rat) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Rat), new(Rat)
		l.Mul(l.Mul(x, y), z)
		r.Mul(x, r.Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Identity

func TestAddZeroRat(t *testing.T) {
	zero := new(Rat)
	f := func(x *Rat) bool {
		// t.Logf("x = %v", x)
		l := new(Rat).Add(x, zero)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulOneRat(t *testing.T) {
	one := new(Rat).One()
	f := func(x *Rat) bool {
		// t.Logf("x = %v", x)
		l := new(Rat).Mul(x, one)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulInvOneRat(t *testing.T) {
	one := new(Rat).One()
	f := func(x *Rat) bool {
		// t.Logf("x = %v", x)
		l := new(Rat)
		l.Mul(x, l.Inv(x))
		return l.Equals(one)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddNegSubRat(t *testing.T) {
	f := func(x, y *Rat) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat), new(Rat)
		l.Sub(x, y)
		r.Add(x, r.Neg(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddScaleDoubleRat(t *testing.T) {
	f := func(x *Rat) bool {
		// t.Logf("x = %v", x)
		l, r := new(Rat), new(Rat)
		l.Add(x, x)
		r.Scale(x, big.NewRat(2, 1))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Involutivity

func TestInvInvolutiveRat(t *testing.T) {
	f := func(x *Rat) bool {
		// t.Logf("x = %v", x)
		l := new(Rat)
		l.Inv(l.Inv(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestNegInvolutiveRat(t *testing.T) {
	f := func(x *Rat) bool {
		// t.Logf("x = %v", x)
		l := new(Rat)
		l.Neg(l.Neg(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBarInvolutiveRat(t *testing.T) {
	f := func(x *Rat) bool {
		// t.Logf("x = %v", x)
		l := new(Rat)
		l.Bar(l.Bar(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestTildeInvolutiveRat(t *testing.T) {
	f := func(x *Rat) bool {
		// t.Logf("x = %v", x)
		l := new(Rat)
		l.Tilde(l.Tilde(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-distributivity

func TestMulBarAntiDistributiveRat(t *testing.T) {
	f := func(x, y *Rat) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat), new(Rat)
		l.Bar(l.Mul(x, y))
		r.Mul(r.Bar(y), new(Rat).Bar(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulTildeAntiDistributiveRat(t *testing.T) {
	f := func(x, y *Rat) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat), new(Rat)
		l.Tilde(l.Mul(x, y))
		r.Mul(r.Tilde(y), new(Rat).Tilde(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulInvAntiDistributiveRat(t *testing.T) {
	f := func(x, y *Rat) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat), new(Rat)
		l.Inv(l.Mul(x, y))
		r.Mul(r.Inv(y), new(Rat).Inv(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Distributivity

func TestAddBarDistributiveRat(t *testing.T) {
	f := func(x, y *Rat) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat), new(Rat)
		l.Add(x, y)
		l.Bar(l)
		r.Add(r.Bar(x), new(Rat).Bar(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddTildeDistributiveRat(t *testing.T) {
	f := func(x, y *Rat) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat), new(Rat)
		l.Add(x, y)
		l.Tilde(l)
		r.Add(r.Tilde(x), new(Rat).Tilde(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubBarDistributiveRat(t *testing.T) {
	f := func(x, y *Rat) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat), new(Rat)
		l.Sub(x, y)
		l.Bar(l)
		r.Sub(r.Bar(x), new(Rat).Bar(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubTildeDistributiveRat(t *testing.T) {
	f := func(x, y *Rat) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat), new(Rat)
		l.Sub(x, y)
		l.Tilde(l)
		r.Sub(r.Tilde(x), new(Rat).Tilde(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddScaleDistributiveRat(t *testing.T) {
	f := func(x, y *Rat) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(Rat), new(Rat)
		l.Scale(l.Add(x, y), a)
		r.Add(r.Scale(x, a), new(Rat).Scale(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubScaleDistributiveRat(t *testing.T) {
	f := func(x, y *Rat) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(Rat), new(Rat)
		l.Scale(l.Sub(x, y), a)
		r.Sub(r.Scale(x, a), new(Rat).Scale(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddMulDistributiveRat(t *testing.T) {
	f := func(x, y, z *Rat) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Rat), new(Rat)
		l.Mul(l.Add(x, y), z)
		r.Add(r.Mul(x, z), new(Rat).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubMulDistributiveRat(t *testing.T) {
	f := func(x, y, z *Rat) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Rat), new(Rat)
		l.Mul(l.Sub(x, y), z)
		r.Sub(r.Mul(x, z), new(Rat).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Positivity

func TestNormPositiveRat(t *testing.T) {
	f := func(x *Rat) bool {
		// t.Logf("x = %v", x)
		return x.Norm().Sign() > 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Composition

func TestCompositionRat(t *testing.T) {
	f := func(x, y *Rat) bool {
		// t.Logf("x = %v, y = %v", x, y)
		p := new(Rat)
		a, b := new(big.Rat), new(big.Rat)
		p.Mul(x, y)
		a.Set(p.Norm())
		b.Mul(x.Norm(), y.Norm())
		return a.Cmp(b) == 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
