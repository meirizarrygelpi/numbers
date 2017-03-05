// Copyright (c) 2017 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package heegner

import (
	"math/big"
	"testing"
	"testing/quick"
)

// Commutativity

func TestAddCommutativeRat1(t *testing.T) {
	f := func(x, y *Rat1) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(Rat1).Add(x, y)
		r := new(Rat1).Add(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulCommutativeRat1(t *testing.T) {
	f := func(x, y *Rat1) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(Rat1).Mul(x, y)
		r := new(Rat1).Mul(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestNegConjCommutativeRat1(t *testing.T) {
	f := func(x *Rat1) bool {
		// t.Logf("x = %v", x)
		l, r := new(Rat1), new(Rat1)
		l.Neg(l.Conj(x))
		r.Conj(r.Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-commutativity

func TestSubAntiCommutativeRat1(t *testing.T) {
	f := func(x, y *Rat1) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat1), new(Rat1)
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

func TestAddAssociativeRat1(t *testing.T) {
	f := func(x, y, z *Rat1) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Rat1), new(Rat1)
		l.Add(l.Add(x, y), z)
		r.Add(x, r.Add(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulAssociativeRat1(t *testing.T) {
	f := func(x, y, z *Rat1) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Rat1), new(Rat1)
		l.Mul(l.Mul(x, y), z)
		r.Mul(x, r.Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Identity

func TestAddZeroRat1(t *testing.T) {
	zero := new(Rat1)
	f := func(x *Rat1) bool {
		// t.Logf("x = %v", x)
		l := new(Rat1).Add(x, zero)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulOneRat1(t *testing.T) {
	one := &Rat1{
		l: *big.NewRat(1, 1),
	}
	f := func(x *Rat1) bool {
		// t.Logf("x = %v", x)
		l := new(Rat1).Mul(x, one)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulInvOneRat1(t *testing.T) {
	one := &Rat1{
		l: *big.NewRat(1, 1),
	}
	f := func(x *Rat1) bool {
		// t.Logf("x = %v", x)
		l := new(Rat1)
		l.Mul(x, l.Inv(x))
		return l.Equals(one)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddNegSubRat1(t *testing.T) {
	f := func(x, y *Rat1) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat1), new(Rat1)
		l.Sub(x, y)
		r.Add(x, r.Neg(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddScaleDoubleRat1(t *testing.T) {
	f := func(x *Rat1) bool {
		// t.Logf("x = %v", x)
		l, r := new(Rat1), new(Rat1)
		l.Add(x, x)
		r.Scale(x, big.NewRat(2, 1))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Involutivity

func TestInvInvolutiveRat1(t *testing.T) {
	f := func(x *Rat1) bool {
		// t.Logf("x = %v", x)
		l := new(Rat1)
		l.Inv(l.Inv(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestNegInvolutiveRat1(t *testing.T) {
	f := func(x *Rat1) bool {
		// t.Logf("x = %v", x)
		l := new(Rat1)
		l.Neg(l.Neg(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestConjInvolutiveRat1(t *testing.T) {
	f := func(x *Rat1) bool {
		// t.Logf("x = %v", x)
		l := new(Rat1)
		l.Conj(l.Conj(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-distributivity

func TestMulConjAntiDistributiveRat1(t *testing.T) {
	f := func(x, y *Rat1) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat1), new(Rat1)
		l.Conj(l.Mul(x, y))
		r.Mul(r.Conj(y), new(Rat1).Conj(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulInvAntiDistributiveRat1(t *testing.T) {
	f := func(x, y *Rat1) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat1), new(Rat1)
		l.Inv(l.Mul(x, y))
		r.Mul(r.Inv(y), new(Rat1).Inv(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Distributivity

func TestAddConjDistributiveRat1(t *testing.T) {
	f := func(x, y *Rat1) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat1), new(Rat1)
		l.Add(x, y)
		l.Conj(l)
		r.Add(r.Conj(x), new(Rat1).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubConjDistributiveRat1(t *testing.T) {
	f := func(x, y *Rat1) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat1), new(Rat1)
		l.Sub(x, y)
		l.Conj(l)
		r.Sub(r.Conj(x), new(Rat1).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddScaleDistributiveRat1(t *testing.T) {
	f := func(x, y *Rat1) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(Rat1), new(Rat1)
		l.Scale(l.Add(x, y), a)
		r.Add(r.Scale(x, a), new(Rat1).Scale(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubScaleDistributiveRat1(t *testing.T) {
	f := func(x, y *Rat1) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewRat(2, 1)
		l, r := new(Rat1), new(Rat1)
		l.Scale(l.Sub(x, y), a)
		r.Sub(r.Scale(x, a), new(Rat1).Scale(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddMulDistributiveRat1(t *testing.T) {
	f := func(x, y, z *Rat1) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Rat1), new(Rat1)
		l.Mul(l.Add(x, y), z)
		r.Add(r.Mul(x, z), new(Rat1).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubMulDistributiveRat1(t *testing.T) {
	f := func(x, y, z *Rat1) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Rat1), new(Rat1)
		l.Mul(l.Sub(x, y), z)
		r.Sub(r.Mul(x, z), new(Rat1).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Positivity

func TestQuadPositiveRat1(t *testing.T) {
	f := func(x *Rat1) bool {
		// t.Logf("x = %v", x)
		return x.Quad().Sign() > 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Composition

func TestCompositionRat1(t *testing.T) {
	f := func(x, y *Rat1) bool {
		// t.Logf("x = %v, y = %v", x, y)
		p := new(Rat1)
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

func TestQuotientsRat1(t *testing.T) {
	f := func(x, y *Rat1) bool {
		// t.Logf("x = %v, y = %v", x, y)
		z, p, q := new(Rat1), new(Rat1), new(Rat1)
		z.Mul(x, y)
		p.Quo(z, y)
		q.Quo(z, x)
		return p.Equals(x) && q.Equals(y)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
