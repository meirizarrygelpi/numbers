// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licensed under the MIT License.

package eisenstein

import (
	"math/big"
	"testing"
	"testing/quick"
)

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

func TestMulCommutativeInt(t *testing.T) {
	f := func(x, y *Int) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(Int).Mul(x, y)
		r := new(Int).Mul(y, x)
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

func TestAddScaleDoubleInt(t *testing.T) {
	f := func(x *Int) bool {
		// t.Logf("x = %v", x)
		l, r := new(Int), new(Int)
		l.Add(x, x)
		r.Scale(x, big.NewInt(2))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

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

func TestAddScaleDistributiveInt(t *testing.T) {
	f := func(x, y *Int) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewInt(2)
		l, r := new(Int), new(Int)
		l.Scale(l.Add(x, y), a)
		r.Add(r.Scale(x, a), new(Int).Scale(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubScaleDistributiveInt(t *testing.T) {
	f := func(x, y *Int) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewInt(2)
		l, r := new(Int), new(Int)
		l.Scale(l.Sub(x, y), a)
		r.Sub(r.Scale(x, a), new(Int).Scale(y, a))
		return l.Equals(r)
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

func TestQuadPositiveInt(t *testing.T) {
	f := func(x *Int) bool {
		// t.Logf("x = %v", x)
		return x.Quad().Sign() > 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// Quotient

func TestQuotientsInt(t *testing.T) {
	f := func(x, y *Int) bool {
		// t.Logf("x = %v, y = %v", x, y)
		z, p, q := new(Int), new(Int), new(Int)
		z.Mul(x, y)
		p.Quo(z, y)
		q.Quo(z, x)
		return p.Equals(x) && q.Equals(y)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
