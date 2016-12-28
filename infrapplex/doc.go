// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

/*
Package infrapplex implements arithmetic for infra-perplex numbers. Five types are
implemented:
    Int64   (int64 values)
    Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
An infra-perplex number has four components and it is written in the form
    a+bs+cτ+dυ
The multiplication table is:
    +-----+----+----+----+
    | Mul | s  | τ  | υ  |
    +-----+----+----+----+
    | s   | +1 | +υ | +τ |
    +-----+----+----+----+
    | τ   | -υ | 0  | 0  |
    +-----+----+----+----+
    | υ   | -τ | 0  | 0  |
    +-----+----+----+----+
The multiplcation operation for infra-perplex numbers is non-commutative but
associative (for non-floats).

Infra-perplex numbers are a parabolic Cayley-Dickson construct with perplex
numbers.
*/
package infrapplex

const (
	leftBracket            = "⦗"
	rightBracket           = "⦘"
	zeroDivisorDenominator = "denominator is zero divisor"
	zeroDivisorInverse     = "inverse of zero divisor"
	unit1                  = "s"
	unit2                  = "τ"
	unit3                  = "υ"
)
