// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

/*
Package dualcplex implements arithmetic for dual-complex numbers. Five types are
implemented:
    Int64   (int64 values)
    Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
A dual-complex number has four components and it is written in the form
    a+bi+cΓ+diΓ
The multiplication table is:
    +-----+----+----+----+
    | Mul | i  | Γ  | iΓ |
    +-----+----+----+----+
    | i   | -1 | iΓ | -Γ |
    +-----+----+----+----+
    | Γ   | iΓ | 0  | 0  |
    +-----+----+----+----+
    | iΓ  | -Γ | 0  | 0  |
    +-----+----+----+----+
The multiplcation operation for dual-complex numbers is commutative and
associative (for non-floats).

Dual-complex numbers are a nilplexification of complex numbers.
*/
package dualcplex

const (
	leftBracket            = "⦗"
	rightBracket           = "⦘"
	zeroDivisorDenominator = "denominator is zero divisor"
	zeroDivisorInverse     = "inverse of zero divisor"
	unit1                  = "i"
	unit2                  = "Γ"
	unit3                  = "iΓ"
)

var (
	unitNames = [3]string{unit1, unit2, unit3}
)

// ResetUnitNames sets the names of the dual-complex units equal to their
// default values.
func ResetUnitNames() {
	unitNames[0] = unit1
	unitNames[1] = unit2
	unitNames[2] = unit3
}

// SetUnitNames sets the names of the dual-complex units.
func SetUnitNames(u1, u2, u3 string) {
	unitNames[0] = u1
	unitNames[1] = u2
	unitNames[2] = u3
}

// UnitNames returns the current names of the dual-complex units.
func UnitNames() [3]string {
	return unitNames
}
