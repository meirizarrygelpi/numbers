// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

/*
Package supracplex implements arithmetic for supra-complex numbers. Five types are
implemented:
    Int64   (int64 values)
    Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
A supra-complex number has eight components and it is written in the form
    a+bi+cα+dβ+eγ+fδ+gε+hζ
The multiplication table is
    ...
The multiplcation operation for supra-complex numbers is non-commutative and
non-associative.

Supra-complex numbers are a parabolic Cayley-Dickson construct with infra-complex
numbers.
*/
package supracplex

const (
	leftBracket            = "⦗"
	rightBracket           = "⦘"
	zeroDivisorDenominator = "denominator is zero divisor"
	zeroDivisorInverse     = "inverse of zero divisor"
	unit1                  = "i"
	unit2                  = "α"
	unit3                  = "β"
	unit4                  = "γ"
	unit5                  = "δ"
	unit6                  = "ε"
	unit7                  = "ζ"
)
