// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

/*
Package tripplex implements arithmetic for tri-perplex numbers. Five types are
implemented:
    Int64   (int64 values)
    Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
The multiplcation operation for tri-perplex numbers is commutative and
associative (for non-floats).
*/
package tripplex

const (
	leftBracket            = "⦗"
	rightBracket           = "⦘"
	zeroDivisorDenominator = "denominator is zero divisor"
	zeroDivisorInverse     = "inverse of zero divisor"
	unit1                  = "s"
	unit2                  = "T"
	unit3                  = "sT"
	unit4                  = "U"
	unit5                  = "sU"
	unit6                  = "TU"
	unit7                  = "sTU"
)
