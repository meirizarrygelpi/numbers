// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

/*
Package trinplex implements arithmetic for tri-nilplex numbers. Five types are
implemented:
    Int64   (int64 values)
    Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
A tri-nilplex number has eight components and it is written in the form
    a+bα+cΓ+dαΓ+eΛ+fαΛ+gΓΛ+hαΓΛ
The multiplication table is:
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | Mul | α   | Γ   | αΓ  | Λ   | αΛ  | ΓΛ  | αΓΛ |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | α   | 0   | αΓ  | 0   | αΛ  | 0   | αΓΛ | 0   |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | Γ   | αΓ  | 0   | 0   | ΓΛ  | αΓΛ | 0   | 0   |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | αΓ  | 0   | 0   | 0   | αΓΛ | 0   | 0   | 0   |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | Λ   | αΛ  | ΓΛ  | αΓΛ | 0   | 0   | 0   | 0   |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | αΛ  | 0   | αΓΛ | 0   | 0   | 0   | 0   | 0   |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | ΓΛ  | αΓΛ | 0   | 0   | 0   | 0   | 0   | 0   |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | αΓΛ | 0   | 0   | 0   | 0   | 0   | 0   | 0   |
    +-----+-----+-----+-----+-----+-----+-----+-----+
The multiplcation operation for tri-nilplex numbers is commutative and
associative (for non-floats).

Tri-nilplex numbers are a nilplexification of hyper numbers.
*/
package trinplex

const (
	leftBracket            = "⦗"
	rightBracket           = "⦘"
	zeroDivisorDenominator = "denominator is zero divisor"
	zeroDivisorInverse     = "inverse of zero divisor"
	unit1                  = "α"
	unit2                  = "Γ"
	unit3                  = "αΓ"
	unit4                  = "Λ"
	unit5                  = "αΛ"
	unit6                  = "ΓΛ"
	unit7                  = "αΓΛ"
)

var (
	unitNames = [7]string{unit1, unit2, unit3, unit4, unit5, unit6, unit7}
)

// ResetUnitNames sets the names of the tri-nilplex units equal to their
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

// SetUnitNames sets the names of the tri-nilplex units.
func SetUnitNames(u1, u2, u3, u4, u5, u6, u7 string) {
	unitNames[0] = u1
	unitNames[1] = u2
	unitNames[2] = u3
	unitNames[3] = u4
	unitNames[4] = u5
	unitNames[5] = u6
	unitNames[6] = u7
}

// UnitNames returns the current names of the tri-nilplex units.
func UnitNames() [7]string {
	return unitNames
}
