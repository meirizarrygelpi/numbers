// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package trinplex

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

func TestNegStar1CommutativeRat(t *testing.T) {
	f := func(x *Rat) bool {
		// t.Logf("x = %v", x)
		l, r := new(Rat), new(Rat)
		l.Neg(l.Star1(x))
		r.Star1(r.Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestNegStar2CommutativeRat(t *testing.T) {
	f := func(x *Rat) bool {
		// t.Logf("x = %v", x)
		l, r := new(Rat), new(Rat)
		l.Neg(l.Star2(x))
		r.Star2(r.Neg(x))
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

func TestStar1InvolutiveRat(t *testing.T) {
	f := func(x *Rat) bool {
		// t.Logf("x = %v", x)
		l := new(Rat)
		l.Star1(l.Star1(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestStar2InvolutiveRat(t *testing.T) {
	f := func(x *Rat) bool {
		// t.Logf("x = %v", x)
		l := new(Rat)
		l.Star2(l.Star2(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-distributivity

func TestMulStar1AntiDistributiveRat(t *testing.T) {
	f := func(x, y *Rat) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat), new(Rat)
		l.Star1(l.Mul(x, y))
		r.Mul(r.Star1(y), new(Rat).Star1(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulStar2AntiDistributiveRat(t *testing.T) {
	f := func(x, y *Rat) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat), new(Rat)
		l.Star2(l.Mul(x, y))
		r.Mul(r.Star2(y), new(Rat).Star2(x))
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

func TestAddStar1DistributiveRat(t *testing.T) {
	f := func(x, y *Rat) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat), new(Rat)
		l.Add(x, y)
		l.Star1(l)
		r.Add(r.Star1(x), new(Rat).Star1(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddStar2DistributiveRat(t *testing.T) {
	f := func(x, y *Rat) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat), new(Rat)
		l.Add(x, y)
		l.Star2(l)
		r.Add(r.Star2(x), new(Rat).Star2(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubStar1DistributiveRat(t *testing.T) {
	f := func(x, y *Rat) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat), new(Rat)
		l.Sub(x, y)
		l.Star1(l)
		r.Sub(r.Star1(x), new(Rat).Star1(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubStar2DistributiveRat(t *testing.T) {
	f := func(x, y *Rat) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Rat), new(Rat)
		l.Sub(x, y)
		l.Star2(l)
		r.Sub(r.Star2(x), new(Rat).Star2(y))
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
