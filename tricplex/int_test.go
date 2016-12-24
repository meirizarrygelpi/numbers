// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package tricplex

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

func TestNegStar1CommutativeInt(t *testing.T) {
	f := func(x *Int) bool {
		// t.Logf("x = %v", x)
		l, r := new(Int), new(Int)
		l.Neg(l.Star1(x))
		r.Star1(r.Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestNegStar2CommutativeInt(t *testing.T) {
	f := func(x *Int) bool {
		// t.Logf("x = %v", x)
		l, r := new(Int), new(Int)
		l.Neg(l.Star2(x))
		r.Star2(r.Neg(x))
		return l.Equals(r)
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

func TestStar1InvolutiveInt(t *testing.T) {
	f := func(x *Int) bool {
		// t.Logf("x = %v", x)
		l := new(Int)
		l.Star1(l.Star1(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestStar2InvolutiveInt(t *testing.T) {
	f := func(x *Int) bool {
		// t.Logf("x = %v", x)
		l := new(Int)
		l.Star2(l.Star2(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Anti-distributivity

func TestMulStar1AntiDistributiveInt(t *testing.T) {
	f := func(x, y *Int) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Int), new(Int)
		l.Star1(l.Mul(x, y))
		r.Mul(r.Star1(y), new(Int).Star1(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulStar2AntiDistributiveInt(t *testing.T) {
	f := func(x, y *Int) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Int), new(Int)
		l.Star2(l.Mul(x, y))
		r.Mul(r.Star2(y), new(Int).Star2(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Distributivity

func TestAddStar1DistributiveInt(t *testing.T) {
	f := func(x, y *Int) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Int), new(Int)
		l.Add(x, y)
		l.Star1(l)
		r.Add(r.Star1(x), new(Int).Star1(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddStar2DistributiveInt(t *testing.T) {
	f := func(x, y *Int) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Int), new(Int)
		l.Add(x, y)
		l.Star2(l)
		r.Add(r.Star2(x), new(Int).Star2(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubStar1DistributiveInt(t *testing.T) {
	f := func(x, y *Int) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Int), new(Int)
		l.Sub(x, y)
		l.Star1(l)
		r.Sub(r.Star1(x), new(Int).Star1(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubStar2DistributiveInt(t *testing.T) {
	f := func(x, y *Int) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Int), new(Int)
		l.Sub(x, y)
		l.Star2(l)
		r.Sub(r.Star2(x), new(Int).Star2(y))
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

func TestNormPositiveInt(t *testing.T) {
	f := func(x *Int) bool {
		// t.Logf("x = %v", x)
		return x.Norm().Sign() > 0
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
		a.Set(p.Norm())
		b.Mul(x.Norm(), y.Norm())
		return a.Cmp(b) == 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}