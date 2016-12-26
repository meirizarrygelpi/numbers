// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package tripplex

import (
	"testing"
	"testing/quick"
)

// Commutativity

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

func TestNegStar1CommutativeInt64(t *testing.T) {
	f := func(x *Int64) bool {
		// t.Logf("x = %v", x)
		l, r := new(Int64), new(Int64)
		l.Neg(l.Star1(x))
		r.Star1(r.Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestNegStar2CommutativeInt64(t *testing.T) {
	f := func(x *Int64) bool {
		// t.Logf("x = %v", x)
		l, r := new(Int64), new(Int64)
		l.Neg(l.Star2(x))
		r.Star2(r.Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-commutativity

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

// Associativity

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

// Identity

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

func TestAddDilateDoubleInt64(t *testing.T) {
	f := func(x *Int64) bool {
		// t.Logf("x = %v", x)
		l, r := new(Int64), new(Int64)
		l.Add(x, x)
		r.Dilate(x, 2)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Involutivity

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

func TestStar1InvolutiveInt64(t *testing.T) {
	f := func(x *Int64) bool {
		// t.Logf("x = %v", x)
		l := new(Int64)
		l.Star1(l.Star1(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestStar2InvolutiveInt64(t *testing.T) {
	f := func(x *Int64) bool {
		// t.Logf("x = %v", x)
		l := new(Int64)
		l.Star2(l.Star2(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-distributivity

func TestMulStar1AntiDistributiveInt64(t *testing.T) {
	f := func(x, y *Int64) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Int64), new(Int64)
		l.Star1(l.Mul(x, y))
		r.Mul(r.Star1(y), new(Int64).Star1(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulStar2AntiDistributiveInt64(t *testing.T) {
	f := func(x, y *Int64) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Int64), new(Int64)
		l.Star2(l.Mul(x, y))
		r.Mul(r.Star2(y), new(Int64).Star2(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Distributivity

func TestAddStar1DistributiveInt64(t *testing.T) {
	f := func(x, y *Int64) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Int64), new(Int64)
		l.Add(x, y)
		l.Star1(l)
		r.Add(r.Star1(x), new(Int64).Star1(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddStar2DistributiveInt64(t *testing.T) {
	f := func(x, y *Int64) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Int64), new(Int64)
		l.Add(x, y)
		l.Star2(l)
		r.Add(r.Star2(x), new(Int64).Star2(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubStar1DistributiveInt64(t *testing.T) {
	f := func(x, y *Int64) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Int64), new(Int64)
		l.Sub(x, y)
		l.Star1(l)
		r.Sub(r.Star1(x), new(Int64).Star1(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubStar2DistributiveInt64(t *testing.T) {
	f := func(x, y *Int64) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Int64), new(Int64)
		l.Sub(x, y)
		l.Star2(l)
		r.Sub(r.Star2(x), new(Int64).Star2(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddDilateDistributiveInt64(t *testing.T) {
	f := func(x, y *Int64) bool {
		// t.Logf("x = %v, y = %v", x, y)
		var a int64 = 2
		l, r := new(Int64), new(Int64)
		l.Dilate(l.Add(x, y), a)
		r.Add(r.Dilate(x, a), new(Int64).Dilate(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubDilateDistributiveInt64(t *testing.T) {
	f := func(x, y *Int64) bool {
		// t.Logf("x = %v, y = %v", x, y)
		var a int64 = 2
		l, r := new(Int64), new(Int64)
		l.Dilate(l.Sub(x, y), a)
		r.Sub(r.Dilate(x, a), new(Int64).Dilate(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddMulDistributiveInt64(t *testing.T) {
	f := func(x, y, z *Int64) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Int64), new(Int64)
		l.Mul(l.Add(x, y), z)
		r.Add(r.Mul(x, z), new(Int64).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubMulDistributiveInt64(t *testing.T) {
	f := func(x, y, z *Int64) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Int64), new(Int64)
		l.Mul(l.Sub(x, y), z)
		r.Sub(r.Mul(x, z), new(Int64).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Composition

func TestCompositionInt64(t *testing.T) {
	f := func(x, y *Int64) bool {
		// t.Logf("x = %v, y = %v", x, y)
		p := new(Int64)
		p.Mul(x, y)
		a := p.Norm()
		b := x.Norm() * y.Norm()
		return a == b
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
