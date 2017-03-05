// Copyright (c) 2017 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package heegner

import (
	"math/big"
	"testing"
	"testing/quick"
)

// Commutativity

func TestAddCommutativeRat67(t *testing.T) {
	f := func(x, y *Rat67) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(Rat67).Add(x, y)
		r := new(Rat67).Add(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulCommutativeRat67(t *testing.T) {
	f := func(x, y *Rat67) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(Rat67).Mul(x, y)
		r := new(Rat67).Mul(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestNegConjCommutativeRat67(t *testing.T) {
	f := func(x *Rat67) bool {
		// t.Logf("x = %v", x)
		l, r := new(Rat67), new(Rat67)
		l.Neg(l.Conj(x))
		r.Conj(r.Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-commutativity

func TestSubAntiCommutativeRat67(t *testing.T) {
	f := func(x, y *Rat67) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat67), new(Rat67)
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

func TestAddAssociativeRat67(t *testing.T) {
	f := func(x, y, z *Rat67) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Rat67), new(Rat67)
		l.Add(l.Add(x, y), z)
		r.Add(x, r.Add(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulAssociativeRat67(t *testing.T) {
	f := func(x, y, z *Rat67) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Rat67), new(Rat67)
		l.Mul(l.Mul(x, y), z)
		r.Mul(x, r.Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Identity

func TestAddZeroRat67(t *testing.T) {
	zero := new(Rat67)
	f := func(x *Rat67) bool {
		// t.Logf("x = %v", x)
		l := new(Rat67).Add(x, zero)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulOneRat67(t *testing.T) {
	one := &Rat67{
		l: *big.NewRat(1, 1),
	}
	f := func(x *Rat67) bool {
		// t.Logf("x = %v", x)
		l := new(Rat67).Mul(x, one)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulInvOneRat67(t *testing.T) {
	one := &Rat67{
		l: *big.NewRat(1, 1),
	}
	f := func(x *Rat67) bool {
		// t.Logf("x = %v", x)
		l := new(Rat67)
		l.Mul(x, l.Inv(x))
		return l.Equals(one)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddNegSubRat67(t *testing.T) {
	f := func(x, y *Rat67) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat67), new(Rat67)
		l.Sub(x, y)
		r.Add(x, r.Neg(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddScaleDoubleRat67(t *testing.T) {
	f := func(x *Rat67) bool {
		// t.Logf("x = %v", x)
		l, r := new(Rat67), new(Rat67)
		l.Add(x, x)
		r.Scale(x, big.NewRat(2, 1))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Involutivity

func TestInvInvolutiveRat67(t *testing.T) {
	f := func(x *Rat67) bool {
		// t.Logf("x = %v", x)
		l := new(Rat67)
		l.Inv(l.Inv(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestNegInvolutiveRat67(t *testing.T) {
	f := func(x *Rat67) bool {
		// t.Logf("x = %v", x)
		l := new(Rat67)
		l.Neg(l.Neg(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestConjInvolutiveRat67(t *testing.T) {
	f := func(x *Rat67) bool {
		// t.Logf("x = %v", x)
		l := new(Rat67)
		l.Conj(l.Conj(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-distributivity

func TestMulConjAntiDistributiveRat67(t *testing.T) {
	f := func(x, y *Rat67) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat67), new(Rat67)
		l.Conj(l.Mul(x, y))
		r.Mul(r.Conj(y), new(Rat67).Conj(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulInvAntiDistributiveRat67(t *testing.T) {
	f := func(x, y *Rat67) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat67), new(Rat67)
		l.Inv(l.Mul(x, y))
		r.Mul(r.Inv(y), new(Rat67).Inv(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Distributivity

func TestAddConjDistributiveRat67(t *testing.T) {
	f := func(x, y *Rat67) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat67), new(Rat67)
		l.Add(x, y)
		l.Conj(l)
		r.Add(r.Conj(x), new(Rat67).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubConjDistributiveRat67(t *testing.T) {
	f := func(x, y *Rat67) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat67), new(Rat67)
		l.Sub(x, y)
		l.Conj(l)
		r.Sub(r.Conj(x), new(Rat67).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddScaleDistributiveRat67(t *testing.T) {
	f := func(x, y *Rat67) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(Rat67), new(Rat67)
		l.Scale(l.Add(x, y), a)
		r.Add(r.Scale(x, a), new(Rat67).Scale(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubScaleDistributiveRat67(t *testing.T) {
	f := func(x, y *Rat67) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(Rat67), new(Rat67)
		l.Scale(l.Sub(x, y), a)
		r.Sub(r.Scale(x, a), new(Rat67).Scale(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddMulDistributiveRat67(t *testing.T) {
	f := func(x, y, z *Rat67) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Rat67), new(Rat67)
		l.Mul(l.Add(x, y), z)
		r.Add(r.Mul(x, z), new(Rat67).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubMulDistributiveRat67(t *testing.T) {
	f := func(x, y, z *Rat67) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Rat67), new(Rat67)
		l.Mul(l.Sub(x, y), z)
		r.Sub(r.Mul(x, z), new(Rat67).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Positivity

func TestQuadPositiveRat67(t *testing.T) {
	f := func(x *Rat67) bool {
		// t.Logf("x = %v", x)
		return x.Quad().Sign() > 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Composition

func TestCompositionRat67(t *testing.T) {
	f := func(x, y *Rat67) bool {
		// t.Logf("x = %v, y = %v", x, y)
		p := new(Rat67)
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

func TestQuotientsRat67(t *testing.T) {
	f := func(x, y *Rat67) bool {
		// t.Logf("x = %v, y = %v", x, y)
		z, p, q := new(Rat67), new(Rat67), new(Rat67)
		z.Mul(x, y)
		p.Quo(z, y)
		q.Quo(z, x)
		return p.Equals(x) && q.Equals(y)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
