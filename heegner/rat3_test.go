// Copyright (c) 2017 Melvin Eloy Irizarry-Gelpí
// Licensed under the MIT License.

package heegner

import (
	"math/big"
	"testing"
	"testing/quick"
)

// Commutativity

func TestAddCommutativeRat3(t *testing.T) {
	f := func(x, y *Rat3) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(Rat3).Add(x, y)
		r := new(Rat3).Add(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulCommutativeRat3(t *testing.T) {
	f := func(x, y *Rat3) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(Rat3).Mul(x, y)
		r := new(Rat3).Mul(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestNegConjCommutativeRat3(t *testing.T) {
	f := func(x *Rat3) bool {
		// t.Logf("x = %v", x)
		l, r := new(Rat3), new(Rat3)
		l.Neg(l.Conj(x))
		r.Conj(r.Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-commutativity

func TestSubAntiCommutativeRat3(t *testing.T) {
	f := func(x, y *Rat3) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat3), new(Rat3)
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

func TestAddAssociativeRat3(t *testing.T) {
	f := func(x, y, z *Rat3) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Rat3), new(Rat3)
		l.Add(l.Add(x, y), z)
		r.Add(x, r.Add(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulAssociativeRat3(t *testing.T) {
	f := func(x, y, z *Rat3) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Rat3), new(Rat3)
		l.Mul(l.Mul(x, y), z)
		r.Mul(x, r.Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Identity

func TestAddZeroRat3(t *testing.T) {
	zero := new(Rat3)
	f := func(x *Rat3) bool {
		// t.Logf("x = %v", x)
		l := new(Rat3).Add(x, zero)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulOneRat3(t *testing.T) {
	one := &Rat3{
		l: *big.NewRat(1, 1),
	}
	f := func(x *Rat3) bool {
		// t.Logf("x = %v", x)
		l := new(Rat3).Mul(x, one)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulInvOneRat3(t *testing.T) {
	one := &Rat3{
		l: *big.NewRat(1, 1),
	}
	f := func(x *Rat3) bool {
		// t.Logf("x = %v", x)
		l := new(Rat3)
		l.Mul(x, l.Inv(x))
		return l.Equals(one)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddNegSubRat3(t *testing.T) {
	f := func(x, y *Rat3) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat3), new(Rat3)
		l.Sub(x, y)
		r.Add(x, r.Neg(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddScaleDoubleRat3(t *testing.T) {
	f := func(x *Rat3) bool {
		// t.Logf("x = %v", x)
		l, r := new(Rat3), new(Rat3)
		l.Add(x, x)
		r.Scale(x, big.NewRat(2, 1))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Involutivity

func TestInvInvolutiveRat3(t *testing.T) {
	f := func(x *Rat3) bool {
		// t.Logf("x = %v", x)
		l := new(Rat3)
		l.Inv(l.Inv(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestNegInvolutiveRat3(t *testing.T) {
	f := func(x *Rat3) bool {
		// t.Logf("x = %v", x)
		l := new(Rat3)
		l.Neg(l.Neg(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestConjInvolutiveRat3(t *testing.T) {
	f := func(x *Rat3) bool {
		// t.Logf("x = %v", x)
		l := new(Rat3)
		l.Conj(l.Conj(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-distributivity

func TestMulConjAntiDistributiveRat3(t *testing.T) {
	f := func(x, y *Rat3) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat3), new(Rat3)
		l.Conj(l.Mul(x, y))
		r.Mul(r.Conj(y), new(Rat3).Conj(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulInvAntiDistributiveRat3(t *testing.T) {
	f := func(x, y *Rat3) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat3), new(Rat3)
		l.Inv(l.Mul(x, y))
		r.Mul(r.Inv(y), new(Rat3).Inv(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Distributivity

func TestAddConjDistributiveRat3(t *testing.T) {
	f := func(x, y *Rat3) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat3), new(Rat3)
		l.Add(x, y)
		l.Conj(l)
		r.Add(r.Conj(x), new(Rat3).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubConjDistributiveRat3(t *testing.T) {
	f := func(x, y *Rat3) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat3), new(Rat3)
		l.Sub(x, y)
		l.Conj(l)
		r.Sub(r.Conj(x), new(Rat3).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddScaleDistributiveRat3(t *testing.T) {
	f := func(x, y *Rat3) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(Rat3), new(Rat3)
		l.Scale(l.Add(x, y), a)
		r.Add(r.Scale(x, a), new(Rat3).Scale(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubScaleDistributiveRat3(t *testing.T) {
	f := func(x, y *Rat3) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(Rat3), new(Rat3)
		l.Scale(l.Sub(x, y), a)
		r.Sub(r.Scale(x, a), new(Rat3).Scale(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddMulDistributiveRat3(t *testing.T) {
	f := func(x, y, z *Rat3) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Rat3), new(Rat3)
		l.Mul(l.Add(x, y), z)
		r.Add(r.Mul(x, z), new(Rat3).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubMulDistributiveRat3(t *testing.T) {
	f := func(x, y, z *Rat3) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Rat3), new(Rat3)
		l.Mul(l.Sub(x, y), z)
		r.Sub(r.Mul(x, z), new(Rat3).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Positivity

func TestQuadPositiveRat3(t *testing.T) {
	f := func(x *Rat3) bool {
		// t.Logf("x = %v", x)
		return x.Quad().Sign() > 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Composition

func TestCompositionRat3(t *testing.T) {
	f := func(x, y *Rat3) bool {
		// t.Logf("x = %v, y = %v", x, y)
		p := new(Rat3)
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

func TestQuotientsRat3(t *testing.T) {
	f := func(x, y *Rat3) bool {
		// t.Logf("x = %v, y = %v", x, y)
		z, p, q := new(Rat3), new(Rat3), new(Rat3)
		z.Mul(x, y)
		p.Quo(z, y)
		q.Quo(z, x)
		return p.Equals(x) && q.Equals(y)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
