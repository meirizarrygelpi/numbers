// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package hypercplex

import (
	"testing"
	"testing/quick"
)

// Commutativity

func TestAddCommutativeFloat64(t *testing.T) {
	f := func(x, y *Float64) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(Float64).Add(x, y)
		r := new(Float64).Add(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestNegStar1CommutativeFloat64(t *testing.T) {
	f := func(x *Float64) bool {
		// t.Logf("x = %v", x)
		l, r := new(Float64), new(Float64)
		l.Neg(l.Bar(x))
		r.Bar(r.Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestNegStar2CommutativeFloat64(t *testing.T) {
	f := func(x *Float64) bool {
		// t.Logf("x = %v", x)
		l, r := new(Float64), new(Float64)
		l.Neg(l.Tilde(x))
		r.Tilde(r.Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-commutativity

func TestSubAntiCommutativeFloat64(t *testing.T) {
	f := func(x, y *Float64) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Float64), new(Float64)
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

func XTestAddAssociativeFloat64(t *testing.T) {
	f := func(x, y, z *Float64) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Float64), new(Float64)
		l.Add(l.Add(x, y), z)
		r.Add(x, r.Add(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func XTestMulAssociativeFloat64(t *testing.T) {
	f := func(x, y, z *Float64) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Float64), new(Float64)
		l.Mul(l.Mul(x, y), z)
		r.Mul(x, r.Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Identity

func TestAddZeroFloat64(t *testing.T) {
	zero := new(Float64)
	f := func(x *Float64) bool {
		// t.Logf("x = %v", x)
		l := new(Float64).Add(x, zero)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulOneFloat64(t *testing.T) {
	one := new(Float64).One()
	f := func(x *Float64) bool {
		// t.Logf("x = %v", x)
		l := new(Float64).Mul(x, one)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func XTestMulInvOneFloat64(t *testing.T) {
	one := new(Float64).One()
	f := func(x *Float64) bool {
		// t.Logf("x = %v", x)
		l := new(Float64)
		l.Mul(x, l.Inv(x))
		return l.Equals(one)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func XTestAddNegSubFloat64(t *testing.T) {
	f := func(x, y *Float64) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Float64), new(Float64)
		l.Sub(x, y)
		r.Add(x, r.Neg(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddDilateDoubleFloat64(t *testing.T) {
	f := func(x *Float64) bool {
		// t.Logf("x = %v", x)
		l, r := new(Float64), new(Float64)
		l.Add(x, x)
		r.Dilate(x, 2)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Involutivity

func XTestInvInvolutiveFloat64(t *testing.T) {
	f := func(x *Float64) bool {
		// t.Logf("x = %v", x)
		l := new(Float64)
		l.Inv(l.Inv(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestNegInvolutiveFloat64(t *testing.T) {
	f := func(x *Float64) bool {
		// t.Logf("x = %v", x)
		l := new(Float64)
		l.Neg(l.Neg(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestStar1InvolutiveFloat64(t *testing.T) {
	f := func(x *Float64) bool {
		// t.Logf("x = %v", x)
		l := new(Float64)
		l.Bar(l.Bar(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestStar2InvolutiveFloat64(t *testing.T) {
	f := func(x *Float64) bool {
		// t.Logf("x = %v", x)
		l := new(Float64)
		l.Tilde(l.Tilde(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-distributivity

func TestMulStar1AntiDistributiveFloat64(t *testing.T) {
	f := func(x, y *Float64) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Float64), new(Float64)
		l.Bar(l.Mul(x, y))
		r.Mul(r.Bar(y), new(Float64).Bar(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulStar2AntiDistributiveFloat64(t *testing.T) {
	f := func(x, y *Float64) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Float64), new(Float64)
		l.Tilde(l.Mul(x, y))
		r.Mul(r.Tilde(y), new(Float64).Tilde(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func XTestMulInvAntiDistributiveFloat64(t *testing.T) {
	f := func(x, y *Float64) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Float64), new(Float64)
		l.Inv(l.Mul(x, y))
		r.Mul(r.Inv(y), new(Float64).Inv(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Distributivity

func TestAddStar1DistributiveFloat64(t *testing.T) {
	f := func(x, y *Float64) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Float64), new(Float64)
		l.Add(x, y)
		l.Bar(l)
		r.Add(r.Bar(x), new(Float64).Bar(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddStar2DistributiveFloat64(t *testing.T) {
	f := func(x, y *Float64) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Float64), new(Float64)
		l.Add(x, y)
		l.Tilde(l)
		r.Add(r.Tilde(x), new(Float64).Tilde(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubStar1DistributiveFloat64(t *testing.T) {
	f := func(x, y *Float64) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Float64), new(Float64)
		l.Sub(x, y)
		l.Bar(l)
		r.Sub(r.Bar(x), new(Float64).Bar(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubStar2DistributiveFloat64(t *testing.T) {
	f := func(x, y *Float64) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Float64), new(Float64)
		l.Sub(x, y)
		l.Tilde(l)
		r.Sub(r.Tilde(x), new(Float64).Tilde(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddDilateDistributiveFloat64(t *testing.T) {
	f := func(x, y *Float64) bool {
		// t.Logf("x = %v, y = %v", x, y)
		var a float64 = 2
		l, r := new(Float64), new(Float64)
		l.Dilate(l.Add(x, y), a)
		r.Add(r.Dilate(x, a), new(Float64).Dilate(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubDilateDistributiveFloat64(t *testing.T) {
	f := func(x, y *Float64) bool {
		// t.Logf("x = %v, y = %v", x, y)
		var a float64 = 2
		l, r := new(Float64), new(Float64)
		l.Dilate(l.Sub(x, y), a)
		r.Sub(r.Dilate(x, a), new(Float64).Dilate(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func XTestAddMulDistributiveFloat64(t *testing.T) {
	f := func(x, y, z *Float64) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Float64), new(Float64)
		l.Mul(l.Add(x, y), z)
		r.Add(r.Mul(x, z), new(Float64).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func XTestSubMulDistributiveFloat64(t *testing.T) {
	f := func(x, y, z *Float64) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Float64), new(Float64)
		l.Mul(l.Sub(x, y), z)
		r.Sub(r.Mul(x, z), new(Float64).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Positivity

func TestNormPositiveFloat64(t *testing.T) {
	f := func(x *Float64) bool {
		// t.Logf("x = %v", x)
		return x.Norm() > 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Composition

func XTestCompositionFloat64(t *testing.T) {
	f := func(x, y *Float64) bool {
		// t.Logf("x = %v, y = %v", x, y)
		p := new(Float64)
		p.Mul(x, y)
		a := p.Norm()
		b := x.Norm() * y.Norm()
		return a == b
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
