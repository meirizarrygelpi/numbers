// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

/*
Package bipplex implements arithmetic for bi-perplex numbers. Five types are
implemented:
    Int64   (int64 values)
    Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
A bi-perplex number has four components and is written in the form
    a+bs+cT+dsT
The multiplication table is:
    +-----+----+----+----+
    | Mul | s  | T  | sT |
    +-----+----+----+----+
    | s   | +1 | sT | T  |
    +-----+----+----+----+
    | T   | sT | +1 | s  |
    +-----+----+----+----+
    | sT  | T  | s  | +1 |
    +-----+----+----+----+
The multiplcation operation for bi-perplex numbers is commutative and
associative (for non-floats).

Bi-perplex numbers are a perplexification of perplex numbers.
*/
package bipplex

const (
	leftBracket            = "⦗"
	rightBracket           = "⦘"
	zeroDivisorDenominator = "denominator is zero divisor"
	zeroDivisorInverse     = "inverse of zero divisor"
	unit1                  = "s"
	unit2                  = "T"
	unit3                  = "sT"
)

var (
	unitNames = [3]string{unit1, unit2, unit3}
)

// ResetUnitNames sets the names of the bi-perplex units equal to their
// default values.
func ResetUnitNames() {
	unitNames[0] = unit1
	unitNames[1] = unit2
	unitNames[2] = unit3
}

// SetUnitNames sets the names of the bi-perplex units.
func SetUnitNames(u1, u2, u3 string) {
	unitNames[0] = u1
	unitNames[1] = u2
	unitNames[2] = u3
}

// UnitNames returns the current names of the bi-perplex units.
func UnitNames() [3]string {
	return unitNames
}
