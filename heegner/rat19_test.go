// Copyright (c) 2017 Melvin Eloy Irizarry-Gelpí
// Licensed under the MIT License.

package heegner

import (
	"math/big"
	"testing"
	"testing/quick"
)

// Commutativity

func TestAddCommutativeRat19(t *testing.T) {
	f := func(x, y *Rat19) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(Rat19).Add(x, y)
		r := new(Rat19).Add(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulCommutativeRat19(t *testing.T) {
	f := func(x, y *Rat19) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(Rat19).Mul(x, y)
		r := new(Rat19).Mul(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestNegConjCommutativeRat19(t *testing.T) {
	f := func(x *Rat19) bool {
		// t.Logf("x = %v", x)
		l, r := new(Rat19), new(Rat19)
		l.Neg(l.Conj(x))
		r.Conj(r.Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-commutativity

func TestSubAntiCommutativeRat19(t *testing.T) {
	f := func(x, y *Rat19) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat19), new(Rat19)
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

func TestAddAssociativeRat19(t *testing.T) {
	f := func(x, y, z *Rat19) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Rat19), new(Rat19)
		l.Add(l.Add(x, y), z)
		r.Add(x, r.Add(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulAssociativeRat19(t *testing.T) {
	f := func(x, y, z *Rat19) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Rat19), new(Rat19)
		l.Mul(l.Mul(x, y), z)
		r.Mul(x, r.Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Identity

func TestAddZeroRat19(t *testing.T) {
	zero := new(Rat19)
	f := func(x *Rat19) bool {
		// t.Logf("x = %v", x)
		l := new(Rat19).Add(x, zero)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulOneRat19(t *testing.T) {
	one := &Rat19{
		l: *big.NewRat(1, 1),
	}
	f := func(x *Rat19) bool {
		// t.Logf("x = %v", x)
		l := new(Rat19).Mul(x, one)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulInvOneRat19(t *testing.T) {
	one := &Rat19{
		l: *big.NewRat(1, 1),
	}
	f := func(x *Rat19) bool {
		// t.Logf("x = %v", x)
		l := new(Rat19)
		l.Mul(x, l.Inv(x))
		return l.Equals(one)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddNegSubRat19(t *testing.T) {
	f := func(x, y *Rat19) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat19), new(Rat19)
		l.Sub(x, y)
		r.Add(x, r.Neg(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddScaleDoubleRat19(t *testing.T) {
	f := func(x *Rat19) bool {
		// t.Logf("x = %v", x)
		l, r := new(Rat19), new(Rat19)
		l.Add(x, x)
		r.Scale(x, big.NewRat(2, 1))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Involutivity

func TestInvInvolutiveRat19(t *testing.T) {
	f := func(x *Rat19) bool {
		// t.Logf("x = %v", x)
		l := new(Rat19)
		l.Inv(l.Inv(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestNegInvolutiveRat19(t *testing.T) {
	f := func(x *Rat19) bool {
		// t.Logf("x = %v", x)
		l := new(Rat19)
		l.Neg(l.Neg(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestConjInvolutiveRat19(t *testing.T) {
	f := func(x *Rat19) bool {
		// t.Logf("x = %v", x)
		l := new(Rat19)
		l.Conj(l.Conj(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-distributivity

func TestMulConjAntiDistributiveRat19(t *testing.T) {
	f := func(x, y *Rat19) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat19), new(Rat19)
		l.Conj(l.Mul(x, y))
		r.Mul(r.Conj(y), new(Rat19).Conj(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulInvAntiDistributiveRat19(t *testing.T) {
	f := func(x, y *Rat19) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat19), new(Rat19)
		l.Inv(l.Mul(x, y))
		r.Mul(r.Inv(y), new(Rat19).Inv(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Distributivity

func TestAddConjDistributiveRat19(t *testing.T) {
	f := func(x, y *Rat19) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat19), new(Rat19)
		l.Add(x, y)
		l.Conj(l)
		r.Add(r.Conj(x), new(Rat19).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubConjDistributiveRat19(t *testing.T) {
	f := func(x, y *Rat19) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat19), new(Rat19)
		l.Sub(x, y)
		l.Conj(l)
		r.Sub(r.Conj(x), new(Rat19).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddScaleDistributiveRat19(t *testing.T) {
	f := func(x, y *Rat19) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(Rat19), new(Rat19)
		l.Scale(l.Add(x, y), a)
		r.Add(r.Scale(x, a), new(Rat19).Scale(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubScaleDistributiveRat19(t *testing.T) {
	f := func(x, y *Rat19) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(Rat19), new(Rat19)
		l.Scale(l.Sub(x, y), a)
		r.Sub(r.Scale(x, a), new(Rat19).Scale(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddMulDistributiveRat19(t *testing.T) {
	f := func(x, y, z *Rat19) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Rat19), new(Rat19)
		l.Mul(l.Add(x, y), z)
		r.Add(r.Mul(x, z), new(Rat19).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubMulDistributiveRat19(t *testing.T) {
	f := func(x, y, z *Rat19) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Rat19), new(Rat19)
		l.Mul(l.Sub(x, y), z)
		r.Sub(r.Mul(x, z), new(Rat19).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Positivity

func TestQuadPositiveRat19(t *testing.T) {
	f := func(x *Rat19) bool {
		// t.Logf("x = %v", x)
		return x.Quad().Sign() > 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Composition

func TestCompositionRat19(t *testing.T) {
	f := func(x, y *Rat19) bool {
		// t.Logf("x = %v, y = %v", x, y)
		p := new(Rat19)
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

func TestQuotientsRat19(t *testing.T) {
	f := func(x, y *Rat19) bool {
		// t.Logf("x = %v, y = %v", x, y)
		z, p, q := new(Rat19), new(Rat19), new(Rat19)
		z.Mul(x, y)
		p.Quo(z, y)
		q.Quo(z, x)
		return p.Equals(x) && q.Equals(y)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
