// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

/*
Package infrahamilton implements arithmetic for infra-Hamilton quaternions. Five types are
implemented:
    Int64   (int64 values)
    Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
An infra-Hamilton quaternion has eight components and it is written in the form
    a+bi+cj+dk+eα+fβ+gγ+hδ
The multiplication table is
    ...
The multiplcation operation for infra-Hamilton quaternions is non-commutative and
non-associative.

Infra-Hamilton quaternions are a parabolic Cayley-Dickson construct with Hamilton
quaternions.
*/
package infrahamilton

const (
	leftBracket            = "⦗"
	rightBracket           = "⦘"
	zeroDivisorDenominator = "denominator is zero divisor"
	zeroDivisorInverse     = "inverse of zero divisor"
	unit1                  = "i"
	unit2                  = "j"
	unit3                  = "k"
	unit4                  = "α"
	unit5                  = "β"
	unit6                  = "γ"
	unit7                  = "δ"
)
