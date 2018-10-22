// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licensed under the MIT License.

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
The multiplication operation for complex numbers is commutative and associative
(for non-floats).

Complex numbers are a complexification of real numbers.
*/
package cplex

const (
	leftBracket     = "("
	rightBracket    = ")"
	zeroDenominator = "denominator is zero"
	zeroInverse     = "inverse of zero"
	unit            = "i"
)

var (
	unitName = unit
)

// ResetUnitName sets the name of the complex unit to its default value.
func ResetUnitName() {
	unitName = unit
}

// SetUnitName sets the name of the complex unit.
func SetUnitName(s string) {
	unitName = s
}

// UnitName returns the current value of the name of the complex unit.
func UnitName() string {
	return unitName
}
