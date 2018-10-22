// Copyright (c) 2017 Melvin Eloy Irizarry-Gelpí
// Licensed under the MIT License.

package heegner

import (
	"math/big"
	"testing"
	"testing/quick"
)

// Commutativity

func TestAddCommutativeRat43(t *testing.T) {
	f := func(x, y *Rat43) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(Rat43).Add(x, y)
		r := new(Rat43).Add(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulCommutativeRat43(t *testing.T) {
	f := func(x, y *Rat43) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(Rat43).Mul(x, y)
		r := new(Rat43).Mul(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestNegConjCommutativeRat43(t *testing.T) {
	f := func(x *Rat43) bool {
		// t.Logf("x = %v", x)
		l, r := new(Rat43), new(Rat43)
		l.Neg(l.Conj(x))
		r.Conj(r.Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-commutativity

func TestSubAntiCommutativeRat43(t *testing.T) {
	f := func(x, y *Rat43) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat43), new(Rat43)
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

func TestAddAssociativeRat43(t *testing.T) {
	f := func(x, y, z *Rat43) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Rat43), new(Rat43)
		l.Add(l.Add(x, y), z)
		r.Add(x, r.Add(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulAssociativeRat43(t *testing.T) {
	f := func(x, y, z *Rat43) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Rat43), new(Rat43)
		l.Mul(l.Mul(x, y), z)
		r.Mul(x, r.Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Identity

func TestAddZeroRat43(t *testing.T) {
	zero := new(Rat43)
	f := func(x *Rat43) bool {
		// t.Logf("x = %v", x)
		l := new(Rat43).Add(x, zero)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulOneRat43(t *testing.T) {
	one := &Rat43{
		l: *big.NewRat(1, 1),
	}
	f := func(x *Rat43) bool {
		// t.Logf("x = %v", x)
		l := new(Rat43).Mul(x, one)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulInvOneRat43(t *testing.T) {
	one := &Rat43{
		l: *big.NewRat(1, 1),
	}
	f := func(x *Rat43) bool {
		// t.Logf("x = %v", x)
		l := new(Rat43)
		l.Mul(x, l.Inv(x))
		return l.Equals(one)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddNegSubRat43(t *testing.T) {
	f := func(x, y *Rat43) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat43), new(Rat43)
		l.Sub(x, y)
		r.Add(x, r.Neg(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddScaleDoubleRat43(t *testing.T) {
	f := func(x *Rat43) bool {
		// t.Logf("x = %v", x)
		l, r := new(Rat43), new(Rat43)
		l.Add(x, x)
		r.Scale(x, big.NewRat(2, 1))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Involutivity

func TestInvInvolutiveRat43(t *testing.T) {
	f := func(x *Rat43) bool {
		// t.Logf("x = %v", x)
		l := new(Rat43)
		l.Inv(l.Inv(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestNegInvolutiveRat43(t *testing.T) {
	f := func(x *Rat43) bool {
		// t.Logf("x = %v", x)
		l := new(Rat43)
		l.Neg(l.Neg(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestConjInvolutiveRat43(t *testing.T) {
	f := func(x *Rat43) bool {
		// t.Logf("x = %v", x)
		l := new(Rat43)
		l.Conj(l.Conj(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-distributivity

func TestMulConjAntiDistributiveRat43(t *testing.T) {
	f := func(x, y *Rat43) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat43), new(Rat43)
		l.Conj(l.Mul(x, y))
		r.Mul(r.Conj(y), new(Rat43).Conj(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulInvAntiDistributiveRat43(t *testing.T) {
	f := func(x, y *Rat43) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat43), new(Rat43)
		l.Inv(l.Mul(x, y))
		r.Mul(r.Inv(y), new(Rat43).Inv(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Distributivity

func TestAddConjDistributiveRat43(t *testing.T) {
	f := func(x, y *Rat43) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat43), new(Rat43)
		l.Add(x, y)
		l.Conj(l)
		r.Add(r.Conj(x), new(Rat43).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubConjDistributiveRat43(t *testing.T) {
	f := func(x, y *Rat43) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat43), new(Rat43)
		l.Sub(x, y)
		l.Conj(l)
		r.Sub(r.Conj(x), new(Rat43).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddScaleDistributiveRat43(t *testing.T) {
	f := func(x, y *Rat43) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(Rat43), new(Rat43)
		l.Scale(l.Add(x, y), a)
		r.Add(r.Scale(x, a), new(Rat43).Scale(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubScaleDistributiveRat43(t *testing.T) {
	f := func(x, y *Rat43) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(Rat43), new(Rat43)
		l.Scale(l.Sub(x, y), a)
		r.Sub(r.Scale(x, a), new(Rat43).Scale(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddMulDistributiveRat43(t *testing.T) {
	f := func(x, y, z *Rat43) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Rat43), new(Rat43)
		l.Mul(l.Add(x, y), z)
		r.Add(r.Mul(x, z), new(Rat43).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubMulDistributiveRat43(t *testing.T) {
	f := func(x, y, z *Rat43) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Rat43), new(Rat43)
		l.Mul(l.Sub(x, y), z)
		r.Sub(r.Mul(x, z), new(Rat43).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Positivity

func TestQuadPositiveRat43(t *testing.T) {
	f := func(x *Rat43) bool {
		// t.Logf("x = %v", x)
		return x.Quad().Sign() > 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Composition

func TestCompositionRat43(t *testing.T) {
	f := func(x, y *Rat43) bool {
		// t.Logf("x = %v, y = %v", x, y)
		p := new(Rat43)
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

func TestQuotientsRat43(t *testing.T) {
	f := func(x, y *Rat43) bool {
		// t.Logf("x = %v, y = %v", x, y)
		z, p, q := new(Rat43), new(Rat43), new(Rat43)
		z.Mul(x, y)
		p.Quo(z, y)
		q.Quo(z, x)
		return p.Equals(x) && q.Equals(y)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
