// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package tripplex

import (
	"math/big"
	"testing"
	"testing/quick"
)

// Commutativity

func TestAddCommutativeFloat(t *testing.T) {
	f := func(x, y *Float) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(Float).Add(x, y)
		r := new(Float).Add(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestNegStar1CommutativeFloat(t *testing.T) {
	f := func(x *Float) bool {
		// t.Logf("x = %v", x)
		l, r := new(Float), new(Float)
		l.Neg(l.Star1(x))
		r.Star1(r.Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestNegStar2CommutativeFloat(t *testing.T) {
	f := func(x *Float) bool {
		// t.Logf("x = %v", x)
		l, r := new(Float), new(Float)
		l.Neg(l.Star2(x))
		r.Star2(r.Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-commutativity

func TestSubAntiCommutativeFloat(t *testing.T) {
	f := func(x, y *Float) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Float), new(Float)
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

func XTestAddAssociativeFloat(t *testing.T) {
	f := func(x, y, z *Float) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Float), new(Float)
		l.Add(l.Add(x, y), z)
		r.Add(x, r.Add(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func XTestMulAssociativeFloat(t *testing.T) {
	f := func(x, y, z *Float) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Float), new(Float)
		l.Mul(l.Mul(x, y), z)
		r.Mul(x, r.Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Identity

func TestAddZeroFloat(t *testing.T) {
	zero := new(Float)
	f := func(x *Float) bool {
		// t.Logf("x = %v", x)
		l := new(Float).Add(x, zero)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulOneFloat(t *testing.T) {
	one := new(Float).One()
	f := func(x *Float) bool {
		// t.Logf("x = %v", x)
		l := new(Float).Mul(x, one)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func XTestMulInvOneFloat(t *testing.T) {
	one := new(Float).One()
	f := func(x *Float) bool {
		// t.Logf("x = %v", x)
		l := new(Float)
		l.Mul(x, l.Inv(x))
		return l.Equals(one)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func XTestAddNegSubFloat(t *testing.T) {
	f := func(x, y *Float) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Float), new(Float)
		l.Sub(x, y)
		r.Add(x, r.Neg(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddDilateDoubleFloat(t *testing.T) {
	f := func(x *Float) bool {
		// t.Logf("x = %v", x)
		l, r := new(Float), new(Float)
		l.Add(x, x)
		r.Dilate(x, big.NewFloat(2))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Involutivity

func XTestInvInvolutiveFloat(t *testing.T) {
	f := func(x *Float) bool {
		// t.Logf("x = %v", x)
		l := new(Float)
		l.Inv(l.Inv(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestNegInvolutiveFloat(t *testing.T) {
	f := func(x *Float) bool {
		// t.Logf("x = %v", x)
		l := new(Float)
		l.Neg(l.Neg(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestStar1InvolutiveFloat(t *testing.T) {
	f := func(x *Float) bool {
		// t.Logf("x = %v", x)
		l := new(Float)
		l.Star1(l.Star1(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestStar2InvolutiveFloat(t *testing.T) {
	f := func(x *Float) bool {
		// t.Logf("x = %v", x)
		l := new(Float)
		l.Star2(l.Star2(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-distributivity

func TestMulStar1AntiDistributiveFloat(t *testing.T) {
	f := func(x, y *Float) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Float), new(Float)
		l.Star1(l.Mul(x, y))
		r.Mul(r.Star1(y), new(Float).Star1(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulStar2AntiDistributiveFloat(t *testing.T) {
	f := func(x, y *Float) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Float), new(Float)
		l.Star2(l.Mul(x, y))
		r.Mul(r.Star2(y), new(Float).Star2(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func XTestMulInvAntiDistributiveFloat(t *testing.T) {
	f := func(x, y *Float) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Float), new(Float)
		l.Inv(l.Mul(x, y))
		r.Mul(r.Inv(y), new(Float).Inv(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Distributivity

func TestAddStar1DistributiveFloat(t *testing.T) {
	f := func(x, y *Float) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Float), new(Float)
		l.Add(x, y)
		l.Star1(l)
		r.Add(r.Star1(x), new(Float).Star1(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddStar2DistributiveFloat(t *testing.T) {
	f := func(x, y *Float) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Float), new(Float)
		l.Add(x, y)
		l.Star2(l)
		r.Add(r.Star2(x), new(Float).Star2(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubStar1DistributiveFloat(t *testing.T) {
	f := func(x, y *Float) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Float), new(Float)
		l.Sub(x, y)
		l.Star1(l)
		r.Sub(r.Star1(x), new(Float).Star1(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubStar2DistributiveFloat(t *testing.T) {
	f := func(x, y *Float) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Float), new(Float)
		l.Sub(x, y)
		l.Star2(l)
		r.Sub(r.Star2(x), new(Float).Star2(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddDilateDistributiveFloat(t *testing.T) {
	f := func(x, y *Float) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewFloat(2)
		l, r := new(Float), new(Float)
		l.Dilate(l.Add(x, y), a)
		r.Add(r.Dilate(x, a), new(Float).Dilate(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubDilateDistributiveFloat(t *testing.T) {
	f := func(x, y *Float) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewFloat(2)
		l, r := new(Float), new(Float)
		l.Dilate(l.Sub(x, y), a)
		r.Sub(r.Dilate(x, a), new(Float).Dilate(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func XTestAddMulDistributiveFloat(t *testing.T) {
	f := func(x, y, z *Float) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Float), new(Float)
		l.Mul(l.Add(x, y), z)
		r.Add(r.Mul(x, z), new(Float).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func XTestSubMulDistributiveFloat(t *testing.T) {
	f := func(x, y, z *Float) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Float), new(Float)
		l.Mul(l.Sub(x, y), z)
		r.Sub(r.Mul(x, z), new(Float).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Composition

func XTestCompositionFloat(t *testing.T) {
	f := func(x, y *Float) bool {
		// t.Logf("x = %v, y = %v", x, y)
		p := new(Float)
		a, b := new(big.Float), new(big.Float)
		p.Mul(x, y)
		a.Set(p.Norm())
		b.Mul(x.Norm(), y.Norm())
		return a.Cmp(b) == 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
