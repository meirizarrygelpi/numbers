// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

/*
Package supercockle implements arithmetic for super-Cockle quaternions. Five types are
implemented:
    Int64   (int64 values)
    Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
An super-Cockle quaternion has eight components and it is written in the form
    a+bi+ct+du+eρ+fσ+gτ+hυ
The multiplication table is
	+-----+----+----+----+----+----+----+----+
	| Mul | i  | t  | u  | ρ  | σ  | τ  | υ  |
	+-----+----+----+----+----+----+----+----+
	| i   | -1 | +u | -t | +σ | -ρ | -υ | +τ |
	+-----+----+----+----+----+----+----+----+
	| t   | -u | +1 | -i | +τ | +υ | +ρ | +σ |
	+-----+----+----+----+----+----+----+----+
	| u   | +t | +1 | +1 | +υ | -τ | -σ | +ρ |
	+-----+----+----+----+----+----+----+----+
	| ρ   | -σ | -τ | -υ | 0  | 0  | 0  | 0  |
	+-----+----+----+----+----+----+----+----+
	| σ   | +ρ | -υ | +τ | 0  | 0  | 0  | 0  |
	+-----+----+----+----+----+----+----+----+
	| τ   | +υ | -ρ | +σ | 0  | 0  | 0  | 0  |
	+-----+----+----+----+----+----+----+----+
	| υ   | -τ | -σ | -ρ | 0  | 0  | 0  | 0  |
	+-----+----+----+----+----+----+----+----+
The multiplcation operation for super-Cockle quaternions is non-commutative and
non-associative.

Infra-Cockle quaternions are a parabolic Cayley-Dickson construct with Cockle
quaternions.
*/
package supercockle

const (
	leftBracket            = "⦗"
	rightBracket           = "⦘"
	zeroDivisorDenominator = "denominator is zero divisor"
	zeroDivisorInverse     = "inverse of zero divisor"
	unit1                  = "i"
	unit2                  = "t"
	unit3                  = "u"
	unit4                  = "ρ"
	unit5                  = "σ"
	unit6                  = "τ"
	unit7                  = "υ"
)

var (
	unitNames = [7]string{unit1, unit2, unit3, unit4, unit5, unit6, unit7}
)

// ResetUnitNames sets the names of the super-Cockle quaternion units equal to
// their default values.
func ResetUnitNames() {
	unitNames[0] = unit1
	unitNames[1] = unit2
	unitNames[2] = unit3
	unitNames[3] = unit4
	unitNames[4] = unit5
	unitNames[5] = unit6
	unitNames[6] = unit7
}

// SetUnitNames sets the names of the super-Cockle quaternion units.
func SetUnitNames(u1, u2, u3, u4, u5, u6, u7 string) {
	unitNames[0] = u1
	unitNames[1] = u2
	unitNames[2] = u3
	unitNames[3] = u4
	unitNames[4] = u5
	unitNames[5] = u6
	unitNames[6] = u7
}

// UnitNames returns the current names of the super-Cockle quaternion units.
func UnitNames() [7]string {
	return unitNames
}
