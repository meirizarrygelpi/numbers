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
A Hamilton quaternion has four components and it is written in the form
    a+bi+cj+dk
The multiplication table is:
    +-----+----+----+----+
    | Mul | i  | j  | k  |
    +-----+----+----+----+
    | i   | -1 | +k | -j |
    +-----+----+----+----+
    | j   | -k | -1 | +i |
    +-----+----+----+----+
    | k   | +j | -i | -1 |
    +-----+----+----+----+
The multiplcation operation for Hamilton quaternions is non-commutative but
associative (for non-floats).

Hamilton quaternions are an elliptic Cayley-Dickson construct with complex
numbers.
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
