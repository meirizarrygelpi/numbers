// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

/*
Package bicplex implements arithmetic for bi-complex numbers. Five types are
implemented:
    Int64   (int64 values)
    Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
A bi-complex number has four components and is written in the form
    a+bi+cJ+diJ
The multiplication table is:
    +-----+----+----+----+
    | Mul | i  | J  | iJ |
    +-----+----+----+----+
    | i   | -1 | iJ | -J |
    +-----+----+----+----+
    | J   | iJ | -1 | -i |
    +-----+----+----+----+
    | iJ  | -J | -i | +1 |
    +-----+----+----+----+
The multiplcation operation for bi-complex numbers is commutative and
associative (for non-floats).

Bi-complex numbers are a complexification of complex numbers.
*/
package bicplex

const (
	leftBracket            = "("
	rightBracket           = ")"
	zeroDivisorDenominator = "denominator is zero divisor"
	zeroDivisorInverse     = "inverse of zero divisor"
	unit1                  = "i"
	unit2                  = "J"
	unit3                  = "iJ"
)

var (
	unitNames = [3]string{unit1, unit2, unit3}
)

// ResetUnitNames sets the names of the bi-complex units equal to their
// default values.
func ResetUnitNames() {
	unitNames[0] = unit1
	unitNames[1] = unit2
	unitNames[2] = unit3
}

// SetUnitNames sets the names of the bi-complex units.
func SetUnitNames(u1, u2, u3 string) {
	unitNames[0] = u1
	unitNames[1] = u2
	unitNames[2] = u3
}

// UnitNames returns the current names of the bi-complex units.
func UnitNames() [3]string {
	return unitNames
}
