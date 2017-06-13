// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

/*
Package pplex implements arithmetic for perplex numbers. Five types are
implemented:
    Int64   (int64 values)
    Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
A perplex number has two components and it is written in the form
    a+bs
The multiplication table is:
    +-----+----+
    | Mul | s  |
    +-----+----+
    | s   | +1 |
    +-----+----+
The multiplcation operation for perplex numbers is commutative and associative
(for non-floats).

Perplex numbers are a perplexification of real numbers.
*/
package pplex

const (
	leftBracket            = "("
	rightBracket           = ")"
	zeroDivisorDenominator = "denominator is zero divisor"
	zeroDivisorInverse     = "inverse of zero divisor"
	unit                   = "s"
)

var (
	unitName = unit
)

// ResetUnitName sets the name of the perplex unit to its default value.
func ResetUnitName() {
	unitName = unit
}

// SetUnitName sets the name of the perplex unit.
func SetUnitName(s string) {
	unitName = s
}

// UnitName returns the current value of the name of the perplex unit.
func UnitName() string {
	return unitName
}
