// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licensed under the MIT License.

package eisenstein

import (
	"testing"
	"testing/quick"
)

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

func XTestMulCommutativeFloat64(t *testing.T) {
	f := func(x, y *Float64) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(Float64).Mul(x, y)
		r := new(Float64).Mul(y, x)
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

func XTestConjInvolutiveFloat64(t *testing.T) {
	f := func(x *Float64) bool {
		// t.Logf("x = %v", x)
		l := new(Float64)
		l.Conj(l.Conj(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestNegConjCommutativeFloat64(t *testing.T) {
	f := func(x *Float64) bool {
		// t.Logf("x = %v", x)
		l, r := new(Float64), new(Float64)
		l.Neg(l.Conj(x))
		r.Conj(r.Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func XTestMulConjAntiDistributiveFloat64(t *testing.T) {
	f := func(x, y *Float64) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Float64), new(Float64)
		l.Conj(l.Mul(x, y))
		r.Mul(r.Conj(y), new(Float64).Conj(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddScaleDoubleFloat64(t *testing.T) {
	f := func(x *Float64) bool {
		// t.Logf("x = %v", x)
		l, r := new(Float64), new(Float64)
		l.Add(x, x)
		r.Scale(x, 2)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

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

func XTestAddConjDistributiveFloat64(t *testing.T) {
	f := func(x, y *Float64) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Float64), new(Float64)
		l.Add(x, y)
		l.Conj(l)
		r.Add(r.Conj(x), new(Float64).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func XTestSubConjDistributiveFloat64(t *testing.T) {
	f := func(x, y *Float64) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Float64), new(Float64)
		l.Sub(x, y)
		l.Conj(l)
		r.Sub(r.Conj(x), new(Float64).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddScaleDistributiveFloat64(t *testing.T) {
	f := func(x, y *Float64) bool {
		// t.Logf("x = %v, y = %v", x, y)
		var a float64 = 2
		l, r := new(Float64), new(Float64)
		l.Scale(l.Add(x, y), a)
		r.Add(r.Scale(x, a), new(Float64).Scale(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubScaleDistributiveFloat64(t *testing.T) {
	f := func(x, y *Float64) bool {
		// t.Logf("x = %v, y = %v", x, y)
		var a float64 = 2
		l, r := new(Float64), new(Float64)
		l.Scale(l.Sub(x, y), a)
		r.Sub(r.Scale(x, a), new(Float64).Scale(y, a))
		return l.Equals(r)
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

func TestQuadPositiveFloat64(t *testing.T) {
	f := func(x *Float64) bool {
		// t.Logf("x = %v", x)
		return x.Quad() > 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
