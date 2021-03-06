// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licensed under the MIT License.

/*
Package binplex implements arithmetic for bi-nilplex numbers. Five types are
implemented:
    Int64   (int64 values)
    Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
A bi-nilplex number has four components and it is written in the form
    a+bα+cΓ+dαΓ
The multiplication table is:
    +-----+----+----+----+
    | Mul | α  | Γ  | αΓ |
    +-----+----+----+----+
    | α   | 0  | αΓ | 0  |
    +-----+----+----+----+
    | Γ   | αΓ | 0  | 0  |
    +-----+----+----+----+
    | αΓ  | 0  | 0  | 0  |
    +-----+----+----+----+
The multiplication operation for bi-nilplex numbers is commutative and
associative (for non-floats).

Hyper numbers are a nilplexification of nilplex numbers.
*/
package binplex

const (
	leftBracket            = "("
	rightBracket           = ")"
	zeroDivisorDenominator = "denominator is zero divisor"
	zeroDivisorInverse     = "inverse of zero divisor"
	unit1                  = "α"
	unit2                  = "Γ"
	unit3                  = "αΓ"
)

var (
	unitNames = [3]string{unit1, unit2, unit3}
)

// ResetUnitNames sets the names of the bi-nilplex units equal to their
// default values.
func ResetUnitNames() {
	unitNames[0] = unit1
	unitNames[1] = unit2
	unitNames[2] = unit3
}

// SetUnitNames sets the names of the bi-nilplex units.
func SetUnitNames(u1, u2, u3 string) {
	unitNames[0] = u1
	unitNames[1] = u2
	unitNames[2] = u3
}

// UnitNames returns the current names of the bi-nilplex units.
func UnitNames() [3]string {
	return unitNames
}
