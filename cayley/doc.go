// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

/*
Package cayley implements arithmetic for Cayley octonions. Five types are
implemented:
    Int64   (int64 values)
    Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
The multiplcation operation for Cayley octonions is non-commutative and
non-associative.
*/
package cayley

const (
	leftBracket     = "⦗"
	rightBracket    = "⦘"
	zeroDenominator = "denominator is zero"
	zeroInverse     = "inverse of zero"
	unit1           = "i"
	unit2           = "j"
	unit3           = "k"
	unit4           = "m"
	unit5           = "n"
	unit6           = "p"
	unit7           = "q"
)
