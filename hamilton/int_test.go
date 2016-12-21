// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package hamilton

import (
	"math/big"
	"testing"
	"testing/quick"
)

// Commutativity

func TestAddCommutativeInt(t *testing.T) {
	f := func(x, y *Int) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(Int).Add(x, y)
		r := new(Int).Add(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestNegConjCommutativeInt(t *testing.T) {
	f := func(x *Int) bool {
		// t.Logf("x = %v", x)
		l, r := new(Int), new(Int)
		l.Neg(l.Conj(x))
		r.Conj(r.Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Non-commutativity

func TestMulNonCommutativeInt(t *testing.T) {
	f := func(x, y *Int) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(Int).Commutator(x, y)
		zero := new(Int)
		return !l.Equals(zero)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-commutativity

func TestSubAntiCommutativeInt(t *testing.T) {
	f := func(x, y *Int) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Int), new(Int)
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

func TestAddAssociativeInt(t *testing.T) {
	f := func(x, y, z *Int) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Int), new(Int)
		l.Add(l.Add(x, y), z)
		r.Add(x, r.Add(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulAssociativeInt(t *testing.T) {
	f := func(x, y, z *Int) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Int), new(Int)
		l.Mul(l.Mul(x, y), z)
		r.Mul(x, r.Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Identity

func TestAddZeroInt(t *testing.T) {
	zero := new(Int)
	f := func(x *Int) bool {
		// t.Logf("x = %v", x)
		l := new(Int).Add(x, zero)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulOneInt(t *testing.T) {
	one := new(Int).One()
	f := func(x *Int) bool {
		// t.Logf("x = %v", x)
		l := new(Int).Mul(x, one)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddNegSubInt(t *testing.T) {
	f := func(x, y *Int) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Int), new(Int)
		l.Sub(x, y)
		r.Add(x, r.Neg(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddDilateDoubleInt(t *testing.T) {
	f := func(x *Int) bool {
		// t.Logf("x = %v", x)
		l, r := new(Int), new(Int)
		l.Add(x, x)
		r.Dilate(x, big.NewInt(2))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Involutivity

func TestNegInvolutiveInt(t *testing.T) {
	f := func(x *Int) bool {
		// t.Logf("x = %v", x)
		l := new(Int)
		l.Neg(l.Neg(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestConjInvolutiveInt(t *testing.T) {
	f := func(x *Int) bool {
		// t.Logf("x = %v", x)
		l := new(Int)
		l.Conj(l.Conj(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-distributivity

func TestMulConjAntiDistributiveInt(t *testing.T) {
	f := func(x, y *Int) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Int), new(Int)
		l.Conj(l.Mul(x, y))
		r.Mul(r.Conj(y), new(Int).Conj(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Distributivity

func TestAddConjDistributiveInt(t *testing.T) {
	f := func(x, y *Int) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Int), new(Int)
		l.Add(x, y)
		l.Conj(l)
		r.Add(r.Conj(x), new(Int).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubConjDistributiveInt(t *testing.T) {
	f := func(x, y *Int) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Int), new(Int)
		l.Sub(x, y)
		l.Conj(l)
		r.Sub(r.Conj(x), new(Int).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddDilateDistributiveInt(t *testing.T) {
	f := func(x, y *Int) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewInt(2)
		l, r := new(Int), new(Int)
		l.Dilate(l.Add(x, y), a)
		r.Add(r.Dilate(x, a), new(Int).Dilate(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubDilateDistributiveInt(t *testing.T) {
	f := func(x, y *Int) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewInt(2)
		l, r := new(Int), new(Int)
		l.Dilate(l.Sub(x, y), a)
		r.Sub(r.Dilate(x, a), new(Int).Dilate(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddMulDistributiveInt(t *testing.T) {
	f := func(x, y, z *Int) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Int), new(Int)
		l.Mul(l.Add(x, y), z)
		r.Add(r.Mul(x, z), new(Int).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubMulDistributiveInt(t *testing.T) {
	f := func(x, y, z *Int) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Int), new(Int)
		l.Mul(l.Sub(x, y), z)
		r.Sub(r.Mul(x, z), new(Int).Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Positivity

func TestQuadPositiveInt(t *testing.T) {
	f := func(x *Int) bool {
		// t.Logf("x = %v", x)
		return x.Quad().Sign() > 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Composition

func TestCompositionInt(t *testing.T) {
	f := func(x, y *Int) bool {
		// t.Logf("x = %v, y = %v", x, y)
		p := new(Int)
		a, b := new(big.Int), new(big.Int)
		p.Mul(x, y)
		a.Set(p.Quad())
		b.Mul(x.Quad(), y.Quad())
		return a.Cmp(b) == 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
