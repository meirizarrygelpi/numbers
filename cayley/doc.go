// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licensed under the MIT License.

/*
Package cayley implements arithmetic for Cayley octonions. Five types are
implemented:
    Int64   (int64 values)
    Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
A Cayley octonion has eight components and it is written in the form
    a+bi+cj+dk+em+fn+gp+hq
The multiplication table is:
    +-----+----+----+----+----+----+----+----+
    | Mul | i  | j  | k  | m  | n  | p  | q  |
    +-----+----+----+----+----+----+----+----+
    | i   | -1 | +k | -j | +n | -m | -q | +p |
    +-----+----+----+----+----+----+----+----+
    | j   | -k | -1 | +i | +p | +q | -m | -n |
    +-----+----+----+----+----+----+----+----+
    | k   | +j | -i | -1 | +q | -p | +n | -m |
    +-----+----+----+----+----+----+----+----+
    | m   | -n | -p | -q | -1 | +i | +j | +k |
    +-----+----+----+----+----+----+----+----+
    | n   | +m | -q | +p | -i | -1 | -k | +j |
    +-----+----+----+----+----+----+----+----+
    | p   | +q | +m | -n | -j | +k | -1 | -i |
    +-----+----+----+----+----+----+----+----+
    | q   | -p | +n | +m | -k | -j | +i | -1 |
    +-----+----+----+----+----+----+----+----+
The multiplication operation for Cayley octonions is non-commutative and
non-associative.

Cayley octonions are an elliptic Cayley-Dickson construct with Hamilton
quaternions.
*/
package cayley

const (
	leftBracket     = "("
	rightBracket    = ")"
	zeroDenominator = "denominator is zero"
	zeroInverse     = "inverse of zero"
	unit1           = "i"
	unit2           = "j"
	unit3           = "k"
	unit4           = "m"
	unit5           = "n"
	unit6           = "p"
	unit7           = "q"
)

var (
	unitNames = [7]string{unit1, unit2, unit3, unit4, unit5, unit6, unit7}
)

// ResetUnitNames sets the names of the Cayley octonion units equal to their
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

// SetUnitNames sets the names of the Cayley octonion units.
func SetUnitNames(u1, u2, u3, u4, u5, u6, u7 string) {
	unitNames[0] = u1
	unitNames[1] = u2
	unitNames[2] = u3
	unitNames[3] = u4
	unitNames[4] = u5
	unitNames[5] = u6
	unitNames[6] = u7
}

// UnitNames returns the current names of the Cayley octonion units.
func UnitNames() [7]string {
	return unitNames
}
