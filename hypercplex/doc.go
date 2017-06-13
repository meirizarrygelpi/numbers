// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

/*
Package hypercplex implements arithmetic for hyper-complex numbers. Five types are
implemented:
    Int64   (int64 values)
    Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
A hyper-complex number has eight components and it is written in the form
    a+bi+cΓ+diΓ+eΛ+fiΛ+gΓΛ+hiΓΛ
The multiplication table is:
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | Mul | i   | Γ   | iΓ  | Λ   | iΛ  | ΓΛ  | iΓΛ |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | i   | -1  | iΓ  | -Γ  | iΛ  | -Λ  | iΓΛ | -ΓΛ |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | Γ   | iΓ  | 0   | 0   | ΓΛ  | iΓΛ | 0   | 0   |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | iΓ  | -Γ  | 0   | 0   | iΓΛ | -ΓΛ | 0   | 0   |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | Λ   | iΛ  | ΓΛ  | iΓΛ | 0   | 0   | 0   | 0   |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | iΛ  | -Λ  | iΓΛ | -ΓΛ | 0   | 0   | 0   | 0   |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | ΓΛ  | iΓΛ | 0   | 0   | 0   | 0   | 0   | 0   |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | iΓΛ | -ΓΛ | 0   | 0   | 0   | 0   | 0   | 0   |
    +-----+-----+-----+-----+-----+-----+-----+-----+
The multiplcation operation for hyper-complex numbers is commutative and
associative (for non-floats).

Hyper-complex numbers are a nilplexification of dual-complex numbers.
*/
package hypercplex

const (
	leftBracket            = "("
	rightBracket           = ")"
	zeroDivisorDenominator = "denominator is zero divisor"
	zeroDivisorInverse     = "inverse of zero divisor"
	unit1                  = "i"
	unit2                  = "Γ"
	unit3                  = "iΓ"
	unit4                  = "Λ"
	unit5                  = "iΛ"
	unit6                  = "ΓΛ"
	unit7                  = "iΓΛ"
)

var (
	unitNames = [7]string{unit1, unit2, unit3, unit4, unit5, unit6, unit7}
)

// ResetUnitNames sets the names of the hyper-complex units equal to their
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

// SetUnitNames sets the names of the hyper-complex units.
func SetUnitNames(u1, u2, u3, u4, u5, u6, u7 string) {
	unitNames[0] = u1
	unitNames[1] = u2
	unitNames[2] = u3
	unitNames[3] = u4
	unitNames[4] = u5
	unitNames[5] = u6
	unitNames[6] = u7
}

// UnitNames returns the current names of the hyper-complex units.
func UnitNames() [7]string {
	return unitNames
}
