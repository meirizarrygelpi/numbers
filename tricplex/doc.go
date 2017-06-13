// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

/*
Package tricplex implements arithmetic for tri-complex numbers. Five types are
implemented:
    Int64   (int64 values)
    Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
A tri-complex number has eight components and it is written in the form
    a+bi+cJ+diJ+eK+fiK+gJK+hiJK
The multiplication table is:
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | Mul | i   | J   | iJ  | K   | iK  | JK  | iJK |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | i   | -1  | iJ  | -J  | iK  | -K  | iJK | -JK |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | J   | iJ  | -1  | -i  | JK  | iJK | -K  | -iK |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | iJ  | -J  | -i  | +1  | iJK | -JK | -iK | +K  |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | K   | iK  | JK  | iJK | -1  | -i  | -J  | -iJ |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | iK  | -K  | iJK | -JK | -i  | +1  | -iJ | +J  |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | JK  | iJK | -K  | -iK | -J  | -iK | +1  | +i  |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | iJK | -JK | -iK | +K  | -iJ | +J  | +i  | -1  |
    +-----+-----+-----+-----+-----+-----+-----+-----+
The multiplcation operation for tri-complex numbers is commutative and
associative (for non-floats).

Tri-complex numbers are a complexification of bi-complex numbers.
*/
package tricplex

const (
	leftBracket            = "("
	rightBracket           = ")"
	zeroDivisorDenominator = "denominator is zero divisor"
	zeroDivisorInverse     = "inverse of zero divisor"
	unit1                  = "i"
	unit2                  = "J"
	unit3                  = "iJ"
	unit4                  = "K"
	unit5                  = "iK"
	unit6                  = "JK"
	unit7                  = "iJK"
)

var (
	unitNames = [7]string{unit1, unit2, unit3, unit4, unit5, unit6, unit7}
)

// ResetUnitNames sets the names of the tri-complex units equal to their
// default values.
func ResetUnitNames() {
	unitNames[0] = unit1
	unitNames[1] = unit2
	unitNames[2] = unit3
	unitNames[3] = unit4
	unitNames[4] = unit5
	unitNames[5] = unit6
	unitNames[6] = unit7
}

// SetUnitNames sets the names of the tri-complex units.
func SetUnitNames(u1, u2, u3, u4, u5, u6, u7 string) {
	unitNames[0] = u1
	unitNames[1] = u2
	unitNames[2] = u3
	unitNames[3] = u4
	unitNames[4] = u5
	unitNames[5] = u6
	unitNames[6] = u7
}

// UnitNames returns the current names of the tri-complex units.
func UnitNames() [7]string {
	return unitNames
}
