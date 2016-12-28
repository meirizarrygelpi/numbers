// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

/*
Package nplex implements arithmetic for nilplex numbers.

Five types are
implemented:
    Int64   (int64 values)
    Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
A nilplex number has two components and it is written in the form
    a+bα
The multiplication table is:
    +-----+---+
    | Mul | α |
    +-----+---+
    | α   | 0 |
    +-----+---+
The multiplcation operation for nilplex numbers is commutative and associative
(for non-floats). There are non-trivial zero divisors.

Nilplex numbers are a nilplexification of real numbers.
*/
package nplex

const (
	leftBracket            = "⦗"
	rightBracket           = "⦘"
	zeroDivisorDenominator = "denominator is zero divisor"
	zeroDivisorInverse     = "inverse of zero divisor"
	unit                   = "α"
)
