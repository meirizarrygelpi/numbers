// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licensed under the MIT License.

package grassmann2

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

func TestNegConjCommutativeRat(t *testing.T) {
	f := func(x *Rat) bool {
		// t.Logf("x = %v", x)
		l, r := new(Rat), new(Rat)
		l.Neg(l.Conj(x))
		r.Conj(r.Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Non-commutativity

func TestMulNonCommutativeRat(t *testing.T) {
	f := func(x, y *Rat) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(Rat).Commutator(x, y)
		zero := new(Rat)
		return !l.Equals(zero)
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

func TestConjInvolutiveRat(t *testing.T) {
	f := func(x *Rat) bool {
		// t.Logf("x = %v", x)
		l := new(Rat)
		l.Conj(l.Conj(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-distributivity

func TestMulConjAntiDistributiveRat(t *testing.T) {
	f := func(x, y *Rat) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat), new(Rat)
		l.Conj(l.Mul(x, y))
		r.Mul(r.Conj(y), new(Rat).Conj(x))
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

func TestAddConjDistributiveRat(t *testing.T) {
	f := func(x, y *Rat) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat), new(Rat)
		l.Add(x, y)
		l.Conj(l)
		r.Add(r.Conj(x), new(Rat).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubConjDistributiveRat(t *testing.T) {
	f := func(x, y *Rat) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat), new(Rat)
		l.Sub(x, y)
		l.Conj(l)
		r.Sub(r.Conj(x), new(Rat).Conj(y))
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

// Composition

func TestCompositionRat(t *testing.T) {
	f := func(x, y *Rat) bool {
		// t.Logf("x = %v, y = %v", x, y)
		p := new(Rat)
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
