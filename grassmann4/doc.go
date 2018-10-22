// Copyright (c) 2016-2017-2017 Melvin Eloy Irizarry-Gelpí
// Licensed under the MIT License.

/*
Package grassmann4 implements arithmetic for four-dimensional Grassmann numbers.
Five types are implemented:
    Int64   (int64 values)
    Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
A four-dimensional Grassmann number has sixteen components and it is written in
the form
    a+bW+cX+dWX+eY+fWY+gXY+h(WX)Y+AZ+BWZ+CXZ+D(WX)Z+EYZ+F(WY)Z+G(XY)Z+H((WX)Y)Z
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
The multiplication operation for four-dimensional Grassmann numbers is
non-commutative and non-associative.

Four-dimensional Grassmann numbers are a parabolic Cayley-Dickson construct with
three-dimensional Grassmann numbers.
*/
package grassmann4

const (
	leftBracket            = "("
	rightBracket           = ")"
	zeroDivisorDenominator = "denominator is zero divisor"
	zeroDivisorInverse     = "inverse of zero divisor"
	unit1                  = "W"
	unit2                  = "X"
	unit3                  = "WX"
	unit4                  = "Y"
	unit5                  = "WY"
	unit6                  = "XY"
	unit7                  = "(WX)Y"
	unit8                  = "Z"
	unit9                  = "WZ"
	unit10                 = "XZ"
	unit11                 = "(WX)Z"
	unit12                 = "YZ"
	unit13                 = "(WY)Z"
	unit14                 = "(XY)Z"
	unit15                 = "((WX)Y)Z"
)

var (
	unitNames = [15]string{unit1, unit2, unit3, unit4, unit5, unit6, unit7,
		unit8, unit9, unit10, unit11, unit12, unit13, unit14, unit15}
)

// ResetUnitNames sets the names of the four-dimensional Grassmann units equal to their
// default values.
func ResetUnitNames() {
	unitNames[0] = unit1
	unitNames[1] = unit2
	unitNames[2] = unit3
	unitNames[3] = unit4
	unitNames[4] = unit5
	unitNames[5] = unit6
	unitNames[6] = unit7
	unitNames[7] = unit8
	unitNames[8] = unit9
	unitNames[9] = unit10
	unitNames[10] = unit11
	unitNames[11] = unit12
	unitNames[12] = unit13
	unitNames[13] = unit14
	unitNames[14] = unit15
}

// SetUnitNames sets the names of the four-dimensional Grassmann units.
func SetUnitNames(u1, u2, u3, u4, u5, u6, u7, u8,
	u9, u10, u11, u12, u13, u14, u15 string) {
	unitNames[0] = u1
	unitNames[1] = u2
	unitNames[2] = u3
	unitNames[3] = u4
	unitNames[4] = u5
	unitNames[5] = u6
	unitNames[6] = u7
	unitNames[7] = u8
	unitNames[8] = u9
	unitNames[9] = u10
	unitNames[10] = u11
	unitNames[11] = u12
	unitNames[12] = u13
	unitNames[13] = u14
	unitNames[14] = u15
}

// UnitNames returns the current names of the four-dimensional Grassmann units.
func UnitNames() [15]string {
	return unitNames
}
