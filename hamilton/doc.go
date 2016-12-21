// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

/*
Package hamilton implements arithmetic for Hamilton quaternions. Five types are
implemented:
    Int64   (int64 values)
 	Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
The multiplcation operation for Hamilton quaternions is non-commutative but
associative (for non-floats).
*/
package hamilton

const (
	leftBracket     = "⦗"
	rightBracket    = "⦘"
	zeroDenominator = "denominator is zero"
	zeroInverse     = "inverse of zero"
	unit1           = "i"
	unit2           = "j"
	unit3           = "k"
)
