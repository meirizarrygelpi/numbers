// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

/*
Package cplex implements arithmetic for complex numbers. Five types are
implemented:
    Int64   (int64 values)
    Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
A complex number has two components and it is written in the form
    a+bi
The multiplication table is:
    +-----+----+
    | Mul | i  |
    +-----+----+
    | i   | -1 |
    +-----+----+
The multiplcation operation for complex numbers is commutative and associative
(for non-floats).

Complex numbers are a complexification of real numbers.
*/
package cplex

const (
	leftBracket     = "⦗"
	rightBracket    = "⦘"
	zeroDenominator = "denominator is zero"
	zeroInverse     = "inverse of zero"
	unit            = "i"
)
