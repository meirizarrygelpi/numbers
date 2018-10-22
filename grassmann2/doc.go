// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licensed under the MIT License.

/*
Package grassmann2 implements arithmetic for two-dimensional Grassmann numbers.
Five types are implemented:
    Int64   (int64 values)
    Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
A two-dimensional Grassmann number has four components and it is written in the
form
    a+bW+cX+dWX
The multiplication table is:
    +-----+-----+-----+----+
    | Mul | W   | X   | WX |
    +-----+-----+-----+----+
    | W   | 0   | +WX | 0  |
    +-----+-----+-----+----+
    | X   | -WX | 0   | 0  |
    +-----+-----+-----+----+
    | WX  | 0   | 0   | 0  |
    +-----+-----+-----+----+
The multiplication operation for two-dimensional Grassmann numbers is
non-commutative but associative (for non-floats).

Two-dimensional Grassmann numbers are a parabolic Cayley-Dickson construct with
nilplex numbers.
*/
package grassmann2

const (
	leftBracket            = "("
	rightBracket           = ")"
	zeroDivisorDenominator = "denominator is zero divisor"
	zeroDivisorInverse     = "inverse of zero divisor"
	unit1                  = "W"
	unit2                  = "X"
	unit3                  = "WX"
)

var (
	unitNames = [3]string{unit1, unit2, unit3}
)

// ResetUnitNames sets the names of the two-dimensional Grassmann units equal
// to their default values.
func ResetUnitNames() {
	unitNames[0] = unit1
	unitNames[1] = unit2
	unitNames[2] = unit3
}

// SetUnitNames sets the names of the two-dimensional Grassmann units.
func SetUnitNames(u1, u2, u3 string) {
	unitNames[0] = u1
	unitNames[1] = u2
	unitNames[2] = u3
}

// UnitNames returns the current names of the two-dimensional Grassmann units.
func UnitNames() [3]string {
	return unitNames
}
