// Copyright (c) 2017 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

/*
Package dualcockle implements arithmetic for dual-Cockle quaternions. Five types are
implemented:
    Int64   (int64 values)
    Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
A dual-Cockle quaternion has eight components and it is written in the form
    a+bi+ct+du+eΓ+fiΓ+gtΓ+huΓ
The multiplication table is
	+-----+-----+-----+-----+-----+-----+-----+-----+
	| Mul | i   | t   | u   | Γ   | iΓ  | tΓ  | uΓ  |
	+-----+-----+-----+-----+-----+-----+-----+-----+
	| i   | -1  | +u  | -t  | +iΓ | -Γ  | +uΓ | -tΓ |
	+-----+-----+-----+-----+-----+-----+-----+-----+
	| t   | -u  | +1  | -i  | +tΓ | -uΓ | +Γ  | -iΓ |
	+-----+-----+-----+-----+-----+-----+-----+-----+
	| u   | +t  | +i  | +1  | +uΓ | +tΓ | +iΓ | +Γ  |
	+-----+-----+-----+-----+-----+-----+-----+-----+
	| Γ   | +iΓ | +tΓ | +uΓ | 0   | 0   | 0   | 0   |
	+-----+-----+-----+-----+-----+-----+-----+-----+
	| iΓ  | -Γ  | +uΓ | -tΓ | 0   | 0   | 0   | 0   |
	+-----+-----+-----+-----+-----+-----+-----+-----+
	| tΓ  | -uΓ | +Γ  | -iΓ | 0   | 0   | 0   | 0   |
	+-----+-----+-----+-----+-----+-----+-----+-----+
	| uΓ  | +tΓ | +iΓ | +Γ  | 0   | 0   | 0   | 0   |
	+-----+-----+-----+-----+-----+-----+-----+-----+
The multiplcation operation for dual-Cockle quaternions is non-commutative but
associative (for non-floats).

Dual-Cockle quaternions are a nilplexification of Cockle quaternions.
*/
package dualcockle

const (
	leftBracket            = "("
	rightBracket           = ")"
	zeroDivisorDenominator = "denominator is zero divisor"
	zeroDivisorInverse     = "inverse of zero divisor"
	unit1                  = "i"
	unit2                  = "t"
	unit3                  = "u"
	unit4                  = "Γ"
	unit5                  = "iΓ"
	unit6                  = "tΓ"
	unit7                  = "uΓ"
)

var (
	unitNames = [7]string{unit1, unit2, unit3, unit4, unit5, unit6, unit7}
)

// ResetUnitNames sets the names of the dual-Cockle quaternion units equal
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

// SetUnitNames sets the names of the dual-Cockle quaternion units.
func SetUnitNames(u1, u2, u3, u4, u5, u6, u7 string) {
	unitNames[0] = u1
	unitNames[1] = u2
	unitNames[2] = u3
	unitNames[3] = u4
	unitNames[4] = u5
	unitNames[5] = u6
	unitNames[6] = u7
}

// UnitNames returns the current names of the dual-Cockle quaternion units.
func UnitNames() [7]string {
	return unitNames
}
