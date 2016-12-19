// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package eisenstein

import (
	"testing"
	"testing/quick"
)

func TestAddCommutativeInt64(t *testing.T) {
	f := func(x, y *Int64) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(Int64).Add(x, y)
		r := new(Int64).Add(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddAssociativeInt64(t *testing.T) {
	f := func(x, y, z *Int64) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Int64), new(Int64)
		l.Add(l.Add(x, y), z)
		r.Add(x, r.Add(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddZeroInt64(t *testing.T) {
	zero := new(Int64)
	f := func(x *Int64) bool {
		// t.Logf("x = %v", x)
		l := new(Int64).Add(x, zero)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulCommutativeInt64(t *testing.T) {
	f := func(x, y *Int64) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(Int64).Mul(x, y)
		r := new(Int64).Mul(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulAssociativeInt64(t *testing.T) {
	f := func(x, y, z *Int64) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Int64), new(Int64)
		l.Mul(l.Mul(x, y), z)
		r.Mul(x, r.Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulOneInt64(t *testing.T) {
	one := new(Int64).One()
	f := func(x *Int64) bool {
		// t.Logf("x = %v", x)
		l := new(Int64).Mul(x, one)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestNegInvolutiveInt64(t *testing.T) {
	f := func(x *Int64) bool {
		// t.Logf("x = %v", x)
		l := new(Int64)
		l.Neg(l.Neg(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestConjInvolutiveInt64(t *testing.T) {
	f := func(x *Int64) bool {
		// t.Logf("x = %v", x)
		l := new(Int64)
		l.Conj(l.Conj(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestNegConjCommutativeInt64(t *testing.T) {
	f := func(x *Int64) bool {
		// t.Logf("x = %v", x)
		l, r := new(Int64), new(Int64)
		l.Neg(l.Conj(x))
		r.Conj(r.Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulConjAntiDistributiveInt64(t *testing.T) {
	f := func(x, y *Int64) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Int64), new(Int64)
		l.Conj(l.Mul(x, y))
		r.Mul(r.Conj(y), new(Int64).Conj(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddScaleDoubleInt64(t *testing.T) {
	f := func(x *Int64) bool {
		// t.Logf("x = %v", x)
		l, r := new(Int64), new(Int64)
		l.Add(x, x)
		r.Scale(x, 2)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubAntiCommutativeInt64(t *testing.T) {
	f := func(x, y *Int64) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Int64), new(Int64)
		l.Sub(x, y)
		r.Sub(y, x)
		r.Neg(r)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddConjDistributiveInt64(t *testing.T) {
	f := func(x, y *Int64) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Int64), new(Int64)
		l.Add(x, y)
		l.Conj(l)
		r.Add(r.Conj(x), new(Int64).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubConjDistributiveInt64(t *testing.T) {
	f := func(x, y *Int64) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Int64), new(Int64)
		l.Sub(x, y)
		l.Conj(l)
		r.Sub(r.Conj(x), new(Int64).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddScaleDistributiveInt64(t *testing.T) {
	f := func(x, y *Int64) bool {
		// t.Logf("x = %v, y = %v", x, y)
		var a int64 = 2
		l, r := new(Int64), new(Int64)
		l.Scale(l.Add(x, y), a)
		r.Add(r.Scale(x, a), new(Int64).Scale(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubScaleDistributiveInt64(t *testing.T) {
	f := func(x, y *Int64) bool {
		// t.Logf("x = %v, y = %v", x, y)
		var a int64 = 2
		l, r := new(Int64), new(Int64)
		l.Scale(l.Sub(x, y), a)
		r.Sub(r.Scale(x, a), new(Int64).Scale(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddNegSubInt64(t *testing.T) {
	f := func(x, y *Int64) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Int64), new(Int64)
		l.Sub(x, y)
		r.Add(x, r.Neg(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// This test can fail due to overflow.
func XTestQuadPositiveInt64(t *testing.T) {
	f := func(x *Int64) bool {
		// t.Logf("x = %v", x)
		return x.Quad() > 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
