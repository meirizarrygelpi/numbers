// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

/*
Package ultra implements arithmetic for ultra numbers. Five types are
implemented:
    Int64   (int64 values)
    Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
An ultra number has eight components and it is written in the form
    a+bα+cβ+dγ+eδ+fε+gζ+hη
The multiplication table is
    ...
The multiplcation operation for ultra numbers is non-commutative and
non-associative.

Ultra numbers are a parabolic Cayley-Dickson construct with supra numbers.
*/
package ultra

const (
	leftBracket            = "⦗"
	rightBracket           = "⦘"
	zeroDivisorDenominator = "denominator is zero divisor"
	zeroDivisorInverse     = "inverse of zero divisor"
	unit1                  = "α"
	unit2                  = "β"
	unit3                  = "γ"
	unit4                  = "δ"
	unit5                  = "ε"
	unit6                  = "ζ"
	unit7                  = "η"
)
