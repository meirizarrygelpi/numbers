// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package grassmann2

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

// Non-commutativity

func TestMulNonCommutativeInt64(t *testing.T) {
	f := func(x, y *Int64) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(Int64).Commutator(x, y)
		zero := new(Int64)
		return !l.Equals(zero)
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

// Anti-distributivity

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

// Distributivity

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
		a := p.Quad()
		b := x.Quad() * y.Quad()
		return a == b
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
