// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

/*
Package zorn implements arithmetic for Zorn octonions. Five types are
implemented:
    Int64   (int64 values)
    Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
The multiplcation operation for Zorn octonions is non-commutative and
non-associative.
*/
package zorn

const (
	leftBracket            = "⦗"
	rightBracket           = "⦘"
	zeroDivisorDenominator = "denominator is zero divisor"
	zeroDivisorInverse     = "inverse of zero divisor"
	unit1                  = "i"
	unit2                  = "j"
	unit3                  = "k"
	unit4                  = "r"
	unit5                  = "s"
	unit6                  = "t"
	unit7                  = "u"
)
