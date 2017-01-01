// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

/*
Package hyperpplex implements arithmetic for hyper-perplex numbers. Five types are
implemented:
    Int64   (int64 values)
    Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
A hyper-perplex number has eight components and it is written in the form
    a+bs+cΓ+dsΓ+eΛ+fsΛ+gΓΛ+hsΓΛ
The multiplication table is:
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | Mul | s   | Γ   | sΓ  | Λ   | sΛ  | ΓΛ  | sΓΛ |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | s   | 1   | sΓ  | Γ   | sΛ  | Λ   | sΓΛ | ΓΛ  |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | Γ   | sΓ  | 0   | 0   | ΓΛ  | sΓΛ | 0   | 0   |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | sΓ  | Γ   | 0   | 0   | sΓΛ | ΓΛ  | 0   | 0   |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | Λ   | sΛ  | ΓΛ  | sΓΛ | 0   | 0   | 0   | 0   |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | sΛ  | Λ   | sΓΛ | ΓΛ  | 0   | 0   | 0   | 0   |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | ΓΛ  | sΓΛ | 0   | 0   | 0   | 0   | 0   | 0   |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | sΓΛ | ΓΛ  | 0   | 0   | 0   | 0   | 0   | 0   |
    +-----+-----+-----+-----+-----+-----+-----+-----+
The multiplcation operation for hyper-perplex numbers is commutative and
associative (for non-floats).

Hyper-perplex numbers are a nilplexification of dual-perplex numbers.
*/
package hyperpplex

const (
	leftBracket            = "⦗"
	rightBracket           = "⦘"
	zeroDivisorDenominator = "denominator is zero divisor"
	zeroDivisorInverse     = "inverse of zero divisor"
	unit1                  = "s"
	unit2                  = "Γ"
	unit3                  = "sΓ"
	unit4                  = "Λ"
	unit5                  = "sΛ"
	unit6                  = "ΓΛ"
	unit7                  = "sΓΛ"
)

var (
	unitNames = [7]string{unit1, unit2, unit3, unit4, unit5, unit6, unit7}
)

// ResetUnitNames sets the names of the hyper-perplex units equal to their
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

// SetUnitNames sets the names of the hyper-perplex units.
func SetUnitNames(u1, u2, u3, u4, u5, u6, u7 string) {
	unitNames[0] = u1
	unitNames[1] = u2
	unitNames[2] = u3
	unitNames[3] = u4
	unitNames[4] = u5
	unitNames[5] = u6
	unitNames[6] = u7
}

// UnitNames returns the current names of the hyper-perplex units.
func UnitNames() [7]string {
	return unitNames
}
