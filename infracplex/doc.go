// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

/*
Package infracplex implements arithmetic for infra-complex numbers. Five types are
implemented:
    Int64   (int64 values)
    Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
An infra-complex number has four components and it is written in the form
    a+bi+cβ+dγ
The multiplication table is:
    +-----+----+----+----+
    | Mul | i  | β  | γ  |
    +-----+----+----+----+
    | i   | -1 | +γ | -β |
    +-----+----+----+----+
    | β   | -γ | 0  | 0  |
    +-----+----+----+----+
    | γ   | +β | 0  | 0  |
    +-----+----+----+----+
The multiplcation operation for infra-complex numbers is non-commutative but
associative (for non-floats).

Infra-complex numbers are a parabolic Cayley-Dickson construct with complex
numbers.
*/
package infracplex

const (
	leftBracket            = "⦗"
	rightBracket           = "⦘"
	zeroDivisorDenominator = "denominator is zero divisor"
	zeroDivisorInverse     = "inverse of zero divisor"
	unit1                  = "i"
	unit2                  = "β"
	unit3                  = "γ"
)
