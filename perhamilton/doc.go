// Copyright (c) 2017 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

/*
Package perhamilton implements arithmetic for perplex-Hamilton quaternions. Five types are
implemented:
    Int64   (int64 values)
    Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
An perplex-Hamilton quaternion has eight components and it is written in the form
    a+bi+cj+dk+eS+fiS+gjS+hkS
The multiplication table is
	+-----+-----+-----+-----+-----+-----+-----+-----+
	| Mul | i   | j   | k   | S   | iS  | jS  | kS  |
	+-----+-----+-----+-----+-----+-----+-----+-----+
	| i   | -1  | +k  | -j  | +iS | -S  | +kS | -jS |
	+-----+-----+-----+-----+-----+-----+-----+-----+
	| j   | -k  | -1  | +i  | +jS | -kS | -S  | +iS |
	+-----+-----+-----+-----+-----+-----+-----+-----+
	| k   | +j  | -i  | -1  | +kS | +jS | -iS | -S  |
	+-----+-----+-----+-----+-----+-----+-----+-----+
	| S   | +iS | +jS | +kS | +1  | +i  | +j  | +k  |
	+-----+-----+-----+-----+-----+-----+-----+-----+
	| iS  | -S  | +kS | -jS | +i  | -1  | +k  | -j  |
	+-----+-----+-----+-----+-----+-----+-----+-----+
	| jS  | -kS | -S  | +iS | +j  | -k  | -1  | +i  |
	+-----+-----+-----+-----+-----+-----+-----+-----+
	| kS  | +jS | -iS | -S  | +k  | +j  | -i  | -1  |
	+-----+-----+-----+-----+-----+-----+-----+-----+
The multiplcation operation for perplex-Hamilton quaternions is non-commutative but
associative (for non-floats).

Perplex-Hamilton quaternions are a perplexification of Hamilton quaternions.
*/
package perhamilton

const (
	leftBracket            = "⦗"
	rightBracket           = "⦘"
	zeroDivisorDenominator = "denominator is zero divisor"
	zeroDivisorInverse     = "inverse of zero divisor"
	unit1                  = "i"
	unit2                  = "j"
	unit3                  = "k"
	unit4                  = "H"
	unit5                  = "iH"
	unit6                  = "jH"
	unit7                  = "kH"
)

var (
	unitNames = [7]string{unit1, unit2, unit3, unit4, unit5, unit6, unit7}
)

// ResetUnitNames sets the names of the perplex-Hamilton quaternion units equal
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

// SetUnitNames sets the names of the perplex-Hamilton quaternion units.
func SetUnitNames(u1, u2, u3, u4, u5, u6, u7 string) {
	unitNames[0] = u1
	unitNames[1] = u2
	unitNames[2] = u3
	unitNames[3] = u4
	unitNames[4] = u5
	unitNames[5] = u6
	unitNames[6] = u7
}

// UnitNames returns the current names of the perplex-Hamilton quaternion units.
func UnitNames() [7]string {
	return unitNames
}
