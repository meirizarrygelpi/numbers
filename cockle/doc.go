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
A Cockle quaternion has four components and it is written in the form
    a+bi+ct+du
The multiplication table is:
    +-----+----+----+----+
    | Mul | i  | t  | u  |
    +-----+----+----+----+
    | i   | -1 | +u | -t |
    +-----+----+----+----+
    | t   | -u | +1 | -i |
    +-----+----+----+----+
    | u   | +t | +i | +1 |
    +-----+----+----+----+
The multiplcation operation for Cockle quaternions is non-commutative but
associative (for non-floats).

Cockle quaternions are a hyperbolic Cayley-Dickson construct with complex
numbers.
*/
package cockle

const (
	leftBracket            = "⦗"
	rightBracket           = "⦘"
	zeroDivisorDenominator = "denominator is zero divisor"
	zeroDivisorInverse     = "inverse of zero divisor"
	unit1                  = "i"
	unit2                  = "t"
	unit3                  = "u"
)
