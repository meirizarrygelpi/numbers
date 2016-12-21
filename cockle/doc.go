// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

/*
Package cockle implements arithmetic for Cockle quaternions. Five types are
implemented:
    Int64   (int64 values)
 	Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
The multiplcation operation for Cockle quaternions is non-commutative but
associative (for non-floats).
*/
package cockle

const (
	leftBracket     = "⦗"
	rightBracket    = "⦘"
	zeroDivisorDenominator = "denominator is zero divisor"
	zeroDivisorInverse     = "inverse of zero divisor"
	unit1           = "i"
	unit2           = "t"
	unit3           = "u"
)
