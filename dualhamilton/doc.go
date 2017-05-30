// Copyright (c) 2017 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

/*
Package dualhamilton implements arithmetic for dual-Hamilton quaternions. Five types are
implemented:
    Int64   (int64 values)
    Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
An dual-Hamilton quaternion has eight components and it is written in the form
    a+bi+cj+dk+eΓ+fiΓ+gjΓ+hkΓ
The multiplication table is
	+-----+-----+-----+-----+-----+-----+-----+-----+
	| Mul | i   | j   | k   | Γ   | iΓ  | jΓ  | kΓ  |
	+-----+-----+-----+-----+-----+-----+-----+-----+
	| i   | -1  | +k  | -j  | +iΓ | -Γ  | +kΓ | -jΓ |
	+-----+-----+-----+-----+-----+-----+-----+-----+
	| j   | -k  | -1  | +i  | +jΓ | -kΓ | -Γ  | +iΓ |
	+-----+-----+-----+-----+-----+-----+-----+-----+
	| k   | +j  | -i  | -1  | +kΓ | +jΓ | -iΓ | -Γ  |
	+-----+-----+-----+-----+-----+-----+-----+-----+
	| Γ   | +iΓ | +jΓ | +kΓ | 0   | 0   | 0   | 0   |
	+-----+-----+-----+-----+-----+-----+-----+-----+
	| iΓ  | -Γ  | +kΓ | -jΓ | 0   | 0   | 0   | 0   |
	+-----+-----+-----+-----+-----+-----+-----+-----+
	| jΓ  | -kΓ | -Γ  | +iΓ | 0   | 0   | 0   | 0   |
	+-----+-----+-----+-----+-----+-----+-----+-----+
	| kΓ  | +jΓ | -iΓ | -Γ  | 0   | 0   | 0   | 0   |
	+-----+-----+-----+-----+-----+-----+-----+-----+
The multiplcation operation for dual-Hamilton quaternions is non-commutative but
associative (for non-floats).

Dual-Hamilton quaternions are a nilplexification of Hamilton quaternions.
*/
package dualhamilton

const (
	leftBracket            = "⦗"
	rightBracket           = "⦘"
	zeroDivisorDenominator = "denominator is zero divisor"
	zeroDivisorInverse     = "inverse of zero divisor"
	unit1                  = "i"
	unit2                  = "j"
	unit3                  = "k"
	unit4                  = "Γ"
	unit5                  = "iΓ"
	unit6                  = "jΓ"
	unit7                  = "kΓ"
)

var (
	unitNames = [7]string{unit1, unit2, unit3, unit4, unit5, unit6, unit7}
)

// ResetUnitNames sets the names of the dual-Hamilton quaternion units equal
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

// SetUnitNames sets the names of the dual-Hamilton quaternion units.
func SetUnitNames(u1, u2, u3, u4, u5, u6, u7 string) {
	unitNames[0] = u1
	unitNames[1] = u2
	unitNames[2] = u3
	unitNames[3] = u4
	unitNames[4] = u5
	unitNames[5] = u6
	unitNames[6] = u7
}

// UnitNames returns the current names of the dual-Hamilton quaternion units.
func UnitNames() [7]string {
	return unitNames
}
