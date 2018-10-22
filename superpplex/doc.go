// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licensed under the MIT License.

/*
Package superpplex implements arithmetic for super-perplex numbers. Five types are
implemented:
    Int64   (int64 values)
    Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
A super-perplex number has four components and it is written in the form
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
The multiplication operation for super-perplex numbers is non-commutative but
associative (for non-floats).

Infra-perplex numbers are a parabolic Cayley-Dickson construct with perplex
numbers.
*/
package superpplex

const (
	leftBracket            = "("
	rightBracket           = ")"
	zeroDivisorDenominator = "denominator is zero divisor"
	zeroDivisorInverse     = "inverse of zero divisor"
	unit1                  = "s"
	unit2                  = "τ"
	unit3                  = "υ"
)

var (
	unitNames = [3]string{unit1, unit2, unit3}
)

// ResetUnitNames sets the names of the super-perplex units equal to their
// default values.
func ResetUnitNames() {
	unitNames[0] = unit1
	unitNames[1] = unit2
	unitNames[2] = unit3
}

// SetUnitNames sets the names of the super-perplex units.
func SetUnitNames(u1, u2, u3 string) {
	unitNames[0] = u1
	unitNames[1] = u2
	unitNames[2] = u3
}

// UnitNames returns the current names of the super-perplex units.
func UnitNames() [3]string {
	return unitNames
}
