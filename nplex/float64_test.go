// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package nplex

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

func TestMulCommutativeFloat64(t *testing.T) {
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

func TestConjInvolutiveFloat64(t *testing.T) {
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

// Anti-distributivity

func TestMulConjAntiDistributiveFloat64(t *testing.T) {
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

func TestAddConjDistributiveFloat64(t *testing.T) {
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

func TestSubConjDistributiveFloat64(t *testing.T) {
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
		var a float64 = 2.0
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
		var a float64 = 2.0
		l, r := new(Float64), new(Float64)
		l.Scale(l.Sub(x, y), a)
		r.Sub(r.Scale(x, a), new(Float64).Scale(y, a))
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

func TestQuadPositiveFloat64(t *testing.T) {
	f := func(x *Float64) bool {
		// t.Logf("x = %v", x)
		return x.Quad() > 0
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
		a := p.Quad()
		b := x.Quad() * y.Quad()
		return a == b
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
