// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

/*
Package ultrapplex implements arithmetic for ultra-perplex numbers. Five types are
implemented:
    Int64   (int64 values)
    Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
A ultra-perplex number has eight components and it is written in the form
    a+bs+cρ+dσ+eτ+fυ+gφ+hψ
The multiplication table is
	+-----+----+----+----+----+----+----+----+
	| Mul | s  | ρ  | σ  | τ  | υ  | φ  | ψ  |
	+-----+----+----+----+----+----+----+----+
	| s   | +1 | +σ | +ρ | +υ | +τ | -ψ | -φ |
	+-----+----+----+----+----+----+----+----+
	| ρ   | -σ | 0  | 0  | +φ | +ψ | 0  | 0  |
	+-----+----+----+----+----+----+----+----+
	| σ   | -ρ | 0  | 0  | +ψ | +φ | 0  | 0  |
	+-----+----+----+----+----+----+----+----+
	| τ   | -υ | -φ | -ψ | 0  | 0  | 0  | 0  |
	+-----+----+----+----+----+----+----+----+
	| υ   | -τ | -ψ | -φ | 0  | 0  | 0  | 0  |
	+-----+----+----+----+----+----+----+----+
	| φ   | +ψ | 0  | 0  | 0  | 0  | 0  | 0  |
	+-----+----+----+----+----+----+----+----+
	| ψ   | +φ | 0  | 0  | 0  | 0  | 0  | 0  |
	+-----+----+----+----+----+----+----+----+
The multiplcation operation for ultra-perplex numbers is non-commutative and
non-associative.

Ultra-perplex numbers are a parabolic Cayley-Dickson construct with infra-perplex
numbers.
*/
package ultrapplex

const (
	leftBracket            = "⦗"
	rightBracket           = "⦘"
	zeroDivisorDenominator = "denominator is zero divisor"
	zeroDivisorInverse     = "inverse of zero divisor"
	unit1                  = "s"
	unit2                  = "ρ"
	unit3                  = "σ"
	unit4                  = "τ"
	unit5                  = "υ"
	unit6                  = "φ"
	unit7                  = "ψ"
)

var (
	unitNames = [7]string{unit1, unit2, unit3, unit4, unit5, unit6, unit7}
)

// ResetUnitNames sets the names of the ultra-perplex units equal to their
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

// SetUnitNames sets the names of the ultra-perplex units.
func SetUnitNames(u1, u2, u3, u4, u5, u6, u7 string) {
	unitNames[0] = u1
	unitNames[1] = u2
	unitNames[2] = u3
	unitNames[3] = u4
	unitNames[4] = u5
	unitNames[5] = u6
	unitNames[6] = u7
}

// UnitNames returns the current names of the ultra-perplex units.
func UnitNames() [7]string {
	return unitNames
}
