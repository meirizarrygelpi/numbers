// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

/*
Package infracockle implements arithmetic for infra-Cockle quaternions. Five types are
implemented:
    Int64   (int64 values)
    Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
An infra-Cockle quaternion has eight components and it is written in the form
    a+bi+ct+du+eρ+fσ+gτ+hυ
The multiplication table is
    ...
The multiplcation operation for infra-Cockle quaternions is non-commutative and
non-associative.

Infra-Cockle quaternions are a parabolic Cayley-Dickson construct with Cockle
quaternions.
*/
package infracockle

const (
	leftBracket            = "⦗"
	rightBracket           = "⦘"
	zeroDivisorDenominator = "denominator is zero divisor"
	zeroDivisorInverse     = "inverse of zero divisor"
	unit1                  = "i"
	unit2                  = "t"
	unit3                  = "u"
	unit4                  = "ρ"
	unit5                  = "σ"
	unit6                  = "τ"
	unit7                  = "υ"
)
