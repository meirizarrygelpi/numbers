// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

/*
Package supra implements arithmetic for supra numbers. Five types are
implemented:
    Int64   (int64 values)
    Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
A supra number has four components and it is written in the form
    a+bα+cβ+dγ
The multiplication table is:
    +-----+----+----+---+
    | Mul | α  | β  | γ |
    +-----+----+----+---+
    | α   | 0  | +γ | 0 |
    +-----+----+----+---+
    | β   | -γ | 0  | 0 |
    +-----+----+----+---+
    | γ   | 0  | 0  | 0 |
    +-----+----+----+---+
The multiplcation operation for supra numbers is non-commutative but
associative (for non-floats).

Supra numbers are a parabolic Cayley-Dickson construct with nilplex numbers.
*/
package supra

const (
	leftBracket            = "⦗"
	rightBracket           = "⦘"
	zeroDivisorDenominator = "denominator is zero divisor"
	zeroDivisorInverse     = "inverse of zero divisor"
	unit1                  = "α"
	unit2                  = "β"
	unit3                  = "γ"
)
