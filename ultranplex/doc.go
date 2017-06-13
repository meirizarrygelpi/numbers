// Copyright (c) 2016-2017 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

/*
Package ultranplex implements arithmetic for ultra-nilplex numbers. Five types
are implemented:
    Int64   (int64 values)
    Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
An ultra-nilplex number has eight components and it is written in the form
    a+bW+cX+dWX+eY+fWY+gXY+h(WX)Y
The multiplication table is
    +-------+--------+--------+--------+--------+--------+--------+-------+
    | Mul   | W      | X      | WX     | Y      | WY     | XY     | (WX)Y |
    +-------+--------+--------+--------+--------+--------+--------+-------+
    | W     | 0      | +WX    | 0      | +WY    | 0      | -(WX)Y | 0     |
    +-------+--------+--------+--------+--------+--------+--------+-------+
    | X     | -WX    | 0      | 0      | +XY    | +(WX)Y | 0      | 0     |
    +-------+--------+--------+--------+--------+--------+--------+-------+
    | WX    | 0      | 0      | 0      | +(WX)Y | 0      | 0      | 0     |
    +-------+--------+--------+--------+--------+--------+--------+-------+
    | Y     | -WY    | -XY    | -(WX)Y | 0      | 0      | 0      | 0     |
    +-------+--------+--------+--------+--------+--------+--------+-------+
    | WY    | 0      | -(WX)Y | 0      | 0      | 0      | 0      | 0     |
    +-------+--------+--------+--------+--------+--------+--------+-------+
    | XY    | +(WX)Y | 0      | 0      | 0      | 0      | 0      | 0     |
    +-------+--------+--------+--------+--------+--------+--------+-------+
    | (WX)Y | 0      | 0      | 0      | 0      | 0      | 0      | 0     |
    +-------+--------+--------+--------+--------+--------+--------+-------+
The multiplcation operation for ultra-nilplex numbers is non-commutative and
non-associative.

Ultra-nilplex numbers are a parabolic Cayley-Dickson construct with
super-nilplex numbers.
*/
package ultranplex

const (
	leftBracket            = "⦗"
	rightBracket           = "⦘"
	zeroDivisorDenominator = "denominator is zero divisor"
	zeroDivisorInverse     = "inverse of zero divisor"
	unit1                  = "W"
	unit2                  = "X"
	unit3                  = "WX"
	unit4                  = "Y"
	unit5                  = "WY"
	unit6                  = "XY"
	unit7                  = "(WX)Y"
)

var (
	unitNames = [7]string{unit1, unit2, unit3, unit4, unit5, unit6, unit7}
)

// ResetUnitNames sets the names of the ultra-nilplex units equal to their
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

// SetUnitNames sets the names of the ultra-nilplex units.
func SetUnitNames(u1, u2, u3, u4, u5, u6, u7 string) {
	unitNames[0] = u1
	unitNames[1] = u2
	unitNames[2] = u3
	unitNames[3] = u4
	unitNames[4] = u5
	unitNames[5] = u6
	unitNames[6] = u7
}

// UnitNames returns the current names of the ultra-nilplex units.
func UnitNames() [7]string {
	return unitNames
}
