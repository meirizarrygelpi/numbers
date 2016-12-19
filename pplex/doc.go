// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

/*
Package pplex implements arithmetic for perplex numbers. Five types are
implemented:
    Int64   (int64 values)
 	Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
The multiplcation operation for perplex numbers is commutative and associative
(for non-floats).
*/
package pplex

const (
	leftBracket            = "⦗"
	rightBracket           = "⦘"
	zeroDivisorDenominator = "denominator is zero"
	zeroDivisorInverse     = "inverse of zero"
	unit                   = "s"
)
