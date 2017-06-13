// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

/*
Package dualpplex implements arithmetic for dual-perplex numbers. Five types are
implemented:
    Int64   (int64 values)
    Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
A dual-perplex number has four components and it is written in the form
    a+bs+cΓ+dsΓ
The multiplication table is:
    +-----+----+----+----+
    | Mul | s  | Γ  | sΓ |
    +-----+----+----+----+
    | s   | +1 | sΓ | +Γ |
    +-----+----+----+----+
    | Γ   | sΓ | 0  | 0  |
    +-----+----+----+----+
    | sΓ  | +Γ | 0  | 0  |
    +-----+----+----+----+
The multiplcation operation for dual-perplex numbers is commutative and
associative (for non-floats).

Dual-perplex numbers are a nilplexification of perplex numbers.
*/
package dualpplex

const (
	leftBracket            = "("
	rightBracket           = ")"
	zeroDivisorDenominator = "denominator is zero divisor"
	zeroDivisorInverse     = "inverse of zero divisor"
	unit1                  = "s"
	unit2                  = "Γ"
	unit3                  = "sΓ"
)

var (
	unitNames = [3]string{unit1, unit2, unit3}
)

// ResetUnitNames sets the names of the dual-perplex units equal to their
// default values.
func ResetUnitNames() {
	unitNames[0] = unit1
	unitNames[1] = unit2
	unitNames[2] = unit3
}

// SetUnitNames sets the names of the dual-perplex units.
func SetUnitNames(u1, u2, u3 string) {
	unitNames[0] = u1
	unitNames[1] = u2
	unitNames[2] = u3
}

// UnitNames returns the current names of the dual-perplex units.
func UnitNames() [3]string {
	return unitNames
}
