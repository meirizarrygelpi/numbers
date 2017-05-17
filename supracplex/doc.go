// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

/*
Package supracplex implements arithmetic for supra-complex numbers. Five types are
implemented:
    Int64   (int64 values)
    Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
A supra-complex number has eight components and it is written in the form
    a+bi+cα+dβ+eγ+fδ+gε+hζ
The multiplication table is
    +-----+----+----+----+----+----+----+----+
	| Mul | i  | α  | β  | γ  | δ  | ε  | ζ  |
	+-----+----+----+----+----+----+----+----+
	| i   | -1 | +β | -α | +δ | -γ | -ζ | +ε |
	+-----+----+----+----+----+----+----+----+
	| α   | -β | 0  | 0  | +ε | +ζ | 0  | 0  |
	+-----+----+----+----+----+----+----+----+
	| β   | +α | 0  | 0  | +ζ | -ε | 0  | 0  |
	+-----+----+----+----+----+----+----+----+
	| γ   | -δ | -ε | -ζ | 0  | 0  | 0  | 0  |
	+-----+----+----+----+----+----+----+----+
	| δ   | +γ | -ζ | +ε | 0  | 0  | 0  | 0  |
	+-----+----+----+----+----+----+----+----+
	| ε   | +ζ | 0  | 0  | 0  | 0  | 0  | 0  |
	+-----+----+----+----+----+----+----+----+
	| ζ   | -ε | 0  | 0  | 0  | 0  | 0  | 0  |
	+-----+----+----+----+----+----+----+----+
The multiplcation operation for supra-complex numbers is non-commutative and
non-associative.

Supra-complex numbers are a parabolic Cayley-Dickson construct with infra-complex
numbers.
*/
package supracplex

const (
	leftBracket            = "⦗"
	rightBracket           = "⦘"
	zeroDivisorDenominator = "denominator is zero divisor"
	zeroDivisorInverse     = "inverse of zero divisor"
	unit1                  = "i"
	unit2                  = "α"
	unit3                  = "β"
	unit4                  = "γ"
	unit5                  = "δ"
	unit6                  = "ε"
	unit7                  = "ζ"
)

var (
	unitNames = [7]string{unit1, unit2, unit3, unit4, unit5, unit6, unit7}
)

// ResetUnitNames sets the names of the supra-complex units equal to their
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

// SetUnitNames sets the names of the supra-complex units.
func SetUnitNames(u1, u2, u3, u4, u5, u6, u7 string) {
	unitNames[0] = u1
	unitNames[1] = u2
	unitNames[2] = u3
	unitNames[3] = u4
	unitNames[4] = u5
	unitNames[5] = u6
	unitNames[6] = u7
}

// UnitNames returns the current names of the supra-complex units.
func UnitNames() [7]string {
	return unitNames
}
