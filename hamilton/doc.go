// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licensed under the MIT License.

/*
Package hamilton implements arithmetic for Hamilton quaternions. Five types are
implemented:
    Int64   (int64 values)
    Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
A Hamilton quaternion has four components and it is written in the form
    a+bi+cj+dk
The multiplication table is:
    +-----+----+----+----+
    | Mul | i  | j  | k  |
    +-----+----+----+----+
    | i   | -1 | +k | -j |
    +-----+----+----+----+
    | j   | -k | -1 | +i |
    +-----+----+----+----+
    | k   | +j | -i | -1 |
    +-----+----+----+----+
The multiplication operation for Hamilton quaternions is non-commutative but
associative (for non-floats).

Hamilton quaternions are an elliptic Cayley-Dickson construct with complex
numbers.
*/
package hamilton

const (
	leftBracket     = "("
	rightBracket    = ")"
	zeroDenominator = "denominator is zero"
	zeroInverse     = "inverse of zero"
	unit1           = "i"
	unit2           = "j"
	unit3           = "k"
)

var (
	unitNames = [3]string{unit1, unit2, unit3}
)

// ResetUnitNames sets the names of the Hamilton quaternion units equal to their
// default values.
func ResetUnitNames() {
	unitNames[0] = unit1
	unitNames[1] = unit2
	unitNames[2] = unit3
}

// SetUnitNames sets the names of the Hamilton quaternion units.
func SetUnitNames(u1, u2, u3 string) {
	unitNames[0] = u1
	unitNames[1] = u2
	unitNames[2] = u3
}

// UnitNames returns the current names of the Hamilton quaternion units.
func UnitNames() [3]string {
	return unitNames
}
