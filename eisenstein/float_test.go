// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licensed under the MIT License.

package eisenstein

import (
	"math/big"
	"testing"
	"testing/quick"
)

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

func TestAddZeroFloat(t *testing.T) {
	zero := new(Float).Zero()
	f := func(x *Float) bool {
		// t.Logf("x = %v", x)
		l := new(Float).Add(x, zero)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulCommutativeFloat(t *testing.T) {
	f := func(x, y *Float) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(Float).Mul(x, y)
		r := new(Float).Mul(y, x)
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

func XTestConjInvolutiveFloat(t *testing.T) {
	f := func(x *Float) bool {
		// t.Logf("x = %v", x)
		l := new(Float)
		l.Conj(l.Conj(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestNegConjCommutativeFloat(t *testing.T) {
	f := func(x *Float) bool {
		// t.Logf("x = %v", x)
		l, r := new(Float), new(Float)
		l.Neg(l.Conj(x))
		r.Conj(r.Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func XTestMulConjAntiDistributiveFloat(t *testing.T) {
	f := func(x, y *Float) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Float), new(Float)
		l.Conj(l.Mul(x, y))
		r.Mul(r.Conj(y), new(Float).Conj(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddScaleDoubleFloat(t *testing.T) {
	f := func(x *Float) bool {
		// t.Logf("x = %v", x)
		l, r := new(Float), new(Float)
		l.Add(x, x)
		r.Scale(x, big.NewFloat(2))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

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

func XTestAddConjDistributiveFloat(t *testing.T) {
	f := func(x, y *Float) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Float), new(Float)
		l.Add(x, y)
		l.Conj(l)
		r.Add(r.Conj(x), new(Float).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func XTestSubConjDistributiveFloat(t *testing.T) {
	f := func(x, y *Float) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Float), new(Float)
		l.Sub(x, y)
		l.Conj(l)
		r.Sub(r.Conj(x), new(Float).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddScaleDistributiveFloat(t *testing.T) {
	f := func(x, y *Float) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewFloat(2)
		l, r := new(Float), new(Float)
		l.Scale(l.Add(x, y), a)
		r.Add(r.Scale(x, a), new(Float).Scale(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubScaleDistributiveFloat(t *testing.T) {
	f := func(x, y *Float) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewFloat(2)
		l, r := new(Float), new(Float)
		l.Scale(l.Sub(x, y), a)
		r.Sub(r.Scale(x, a), new(Float).Scale(y, a))
		return l.Equals(r)
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

func TestQuadPositiveFloat(t *testing.T) {
	f := func(x *Float) bool {
		// t.Logf("x = %v", x)
		return x.Quad().Sign() > 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
