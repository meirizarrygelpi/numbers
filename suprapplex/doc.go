// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

/*
Package suprapplex implements arithmetic for supra-perplex numbers. Five types are
implemented:
    Int64   (int64 values)
    Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
A supra-perplex number has eight components and it is written in the form
    a+bs+cρ+dσ+eτ+fυ+gφ+hψ
The multiplication table is
    ...
The multiplcation operation for supra-perplex numbers is non-commutative and
non-associative.

Supra-perplex numbers are a parabolic Cayley-Dickson construct with infra-perplex
numbers.
*/
package suprapplex

const (
	leftBracket            = "⦗"
	rightBracket           = "⦘"
	zeroDivisorDenominator = "denominator is zero divisor"
	zeroDivisorInverse     = "inverse of zero divisor"
	unit1                  = "s"
	unit2                  = "ρ"
	unit3                  = "σ"
	unit4                  = "τ"
	unit5                  = "υ"
	unit6                  = "φ"
	unit7                  = "ψ"
)
