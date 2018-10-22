// Copyright (c) 2017 Melvin Eloy Irizarry-Gelpí
// Licensed under the MIT License.

package heegner

import (
	"math/big"
	"testing"
	"testing/quick"
)

// Commutativity

func TestAddCommutativeRat2(t *testing.T) {
	f := func(x, y *Rat2) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(Rat2).Add(x, y)
		r := new(Rat2).Add(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulCommutativeRat2(t *testing.T) {
	f := func(x, y *Rat2) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(Rat2).Mul(x, y)
		r := new(Rat2).Mul(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestNegConjCommutativeRat2(t *testing.T) {
	f := func(x *Rat2) bool {
		// t.Logf("x = %v", x)
		l, r := new(Rat2), new(Rat2)
		l.Neg(l.Conj(x))
		r.Conj(r.Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-commutativity

func TestSubAntiCommutativeRat2(t *testing.T) {
	f := func(x, y *Rat2) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat2), new(Rat2)
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

func TestAddAssociativeRat2(t *testing.T) {
	f := func(x, y, z *Rat2) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Rat2), new(Rat2)
		l.Add(l.Add(x, y), z)
		r.Add(x, r.Add(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulAssociativeRat2(t *testing.T) {
	f := func(x, y, z *Rat2) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Rat2), new(Rat2)
		l.Mul(l.Mul(x, y), z)
		r.Mul(x, r.Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Identity

func TestAddZeroRat2(t *testing.T) {
	zero := new(Rat2)
	f := func(x *Rat2) bool {
		// t.Logf("x = %v", x)
		l := new(Rat2).Add(x, zero)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulOneRat2(t *testing.T) {
	one := &Rat2{
		l: *big.NewRat(1, 1),
	}
	f := func(x *Rat2) bool {
		// t.Logf("x = %v", x)
		l := new(Rat2).Mul(x, one)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulInvOneRat2(t *testing.T) {
	one := &Rat2{
		l: *big.NewRat(1, 1),
	}
	f := func(x *Rat2) bool {
		// t.Logf("x = %v", x)
		l := new(Rat2)
		l.Mul(x, l.Inv(x))
		return l.Equals(one)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddNegSubRat2(t *testing.T) {
	f := func(x, y *Rat2) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat2), new(Rat2)
		l.Sub(x, y)
		r.Add(x, r.Neg(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddScaleDoubleRat2(t *testing.T) {
	f := func(x *Rat2) bool {
		// t.Logf("x = %v", x)
		l, r := new(Rat2), new(Rat2)
		l.Add(x, x)
		r.Scale(x, big.NewRat(2, 1))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Involutivity

func TestInvInvolutiveRat2(t *testing.T) {
	f := func(x *Rat2) bool {
		// t.Logf("x = %v", x)
		l := new(Rat2)
		l.Inv(l.Inv(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestNegInvolutiveRat2(t *testing.T) {
	f := func(x *Rat2) bool {
		// t.Logf("x = %v", x)
		l := new(Rat2)
		l.Neg(l.Neg(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestConjInvolutiveRat2(t *testing.T) {
	f := func(x *Rat2) bool {
		// t.Logf("x = %v", x)
		l := new(Rat2)
		l.Conj(l.Conj(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-distributivity

func TestMulConjAntiDistributiveRat2(t *testing.T) {
	f := func(x, y *Rat2) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat2), new(Rat2)
		l.Conj(l.Mul(x, y))
		r.Mul(r.Conj(y), new(Rat2).Conj(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulInvAntiDistributiveRat2(t *testing.T) {
	f := func(x, y *Rat2) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat2), new(Rat2)
		l.Inv(l.Mul(x, y))
		r.Mul(r.Inv(y), new(Rat2).Inv(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Distributivity

func TestAddConjDistributiveRat2(t *testing.T) {
	f := func(x, y *Rat2) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat2), new(Rat2)
		l.Add(x, y)
		l.Conj(l)
		r.Add(r.Conj(x), new(Rat2).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubConjDistributiveRat2(t *testing.T) {
	f := func(x, y *Rat2) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat2), new(Rat2)
		l.Sub(x, y)
		l.Conj(l)
		r.Sub(r.Conj(x), new(Rat2).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddScaleDistributiveRat2(t *testing.T) {
	f := func(x, y *Rat2) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(Rat2), new(Rat2)
		l.Scale(l.Add(x, y), a)
		r.Add(r.Scale(x, a), new(Rat2).Scale(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubScaleDistributiveRat2(t *testing.T) {
	f := func(x, y *Rat2) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(Rat2), new(Rat2)
		l.Scale(l.Sub(x, y), a)
		r.Sub(r.Scale(x, a), new(Rat2).Scale(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddMulDistributiveRat2(t *testing.T) {
	f := func(x, y, z *Rat2) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Rat2), new(Rat2)
		l.Mul(l.Add(x, y), z)
		r.Add(r.Mul(x, z), new(Rat2).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubMulDistributiveRat2(t *testing.T) {
	f := func(x, y, z *Rat2) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Rat2), new(Rat2)
		l.Mul(l.Sub(x, y), z)
		r.Sub(r.Mul(x, z), new(Rat2).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Positivity

func TestQuadPositiveRat2(t *testing.T) {
	f := func(x *Rat2) bool {
		// t.Logf("x = %v", x)
		return x.Quad().Sign() > 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Composition

func TestCompositionRat2(t *testing.T) {
	f := func(x, y *Rat2) bool {
		// t.Logf("x = %v, y = %v", x, y)
		p := new(Rat2)
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

func TestQuotientsRat2(t *testing.T) {
	f := func(x, y *Rat2) bool {
		// t.Logf("x = %v, y = %v", x, y)
		z, p, q := new(Rat2), new(Rat2), new(Rat2)
		z.Mul(x, y)
		p.Quo(z, y)
		q.Quo(z, x)
		return p.Equals(x) && q.Equals(y)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
