// Copyright (c) 2017 Melvin Eloy Irizarry-Gelpí
// Licensed under the MIT License.

/*
Package percockle implements arithmetic for perplex-Cockle quaternions. Five types are
implemented:
    Int64   (int64 values)
    Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
An perplex-Cockle quaternion has eight components and it is written in the form
    a+bi+ct+du+eS+fiS+gtS+huS
The multiplication table is
	+-----+-----+-----+-----+-----+-----+-----+-----+
	| Mul | i   | t   | u   | S   | iS  | tS  | uS  |
	+-----+-----+-----+-----+-----+-----+-----+-----+
	| i   | -1  | +u  | -t  | +iS | -S  | +uS | -tS |
	+-----+-----+-----+-----+-----+-----+-----+-----+
	| t   | -u  | +1  | -i  | +tS | -uS | +S  | -iS |
	+-----+-----+-----+-----+-----+-----+-----+-----+
	| u   | +t  | +i  | +1  | +uS | +tS | +iS | +S  |
	+-----+-----+-----+-----+-----+-----+-----+-----+
	| S   | +iS | +tS | +uS | +1  | +i  | +t  | +u  |
	+-----+-----+-----+-----+-----+-----+-----+-----+
	| iS  | -S  | +uS | -tS | +i  | -1  | +u  | -t  |
	+-----+-----+-----+-----+-----+-----+-----+-----+
	| tS  | -uS | +S  | -iS | +t  | -u  | +1  | -i  |
	+-----+-----+-----+-----+-----+-----+-----+-----+
	| uS  | +tS | +iS | +S  | +u  | +t  | +i  | +1  |
	+-----+-----+-----+-----+-----+-----+-----+-----+
The multiplication operation for perplex-Cockle quaternions is non-commutative but
associative (for non-floats).

Perplex-Cockle quaternions are a perplexification of Cockle quaternions.
*/
package percockle

const (
	leftBracket            = "("
	rightBracket           = ")"
	zeroDivisorDenominator = "denominator is zero divisor"
	zeroDivisorInverse     = "inverse of zero divisor"
	unit1                  = "i"
	unit2                  = "t"
	unit3                  = "u"
	unit4                  = "S"
	unit5                  = "iS"
	unit6                  = "tS"
	unit7                  = "uS"
)

var (
	unitNames = [7]string{unit1, unit2, unit3, unit4, unit5, unit6, unit7}
)

// ResetUnitNames sets the names of the perplex-Cockle quaternion units equal
// to their default values.
func ResetUnitNames() {
	unitNames[0] = unit1
	unitNames[1] = unit2
	unitNames[2] = unit3
	unitNames[3] = unit4
	unitNames[4] = unit5
	unitNames[5] = unit6
	unitNames[6] = unit7
}

// SetUnitNames sets the names of the perplex-Cockle quaternion units.
func SetUnitNames(u1, u2, u3, u4, u5, u6, u7 string) {
	unitNames[0] = u1
	unitNames[1] = u2
	unitNames[2] = u3
	unitNames[3] = u4
	unitNames[4] = u5
	unitNames[5] = u6
	unitNames[6] = u7
}

// UnitNames returns the current names of the perplex-Cockle quaternion units.
func UnitNames() [7]string {
	return unitNames
}
