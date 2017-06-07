// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

/*
Package infrahamilton implements arithmetic for infra-Hamilton quaternions. Five types are
implemented:
    Int64   (int64 values)
    Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
An infra-Hamilton quaternion has eight components and it is written in the form
    a+bi+cj+dk+eα+fβ+gγ+hδ
The multiplication table is
	+-----+----+----+----+----+----+----+----+
	| Mul | i  | j  | k  | α  | β  | γ  | δ  |
	+-----+----+----+----+----+----+----+----+
	| i   | -1 | +k | -j | +β | -α | -δ | +γ |
	+-----+----+----+----+----+----+----+----+
	| j   | -k | -1 | +i | +γ | +δ | -α | -β |
	+-----+----+----+----+----+----+----+----+
	| k   | +j | -i | -1 | +δ | -γ | +β | -α |
	+-----+----+----+----+----+----+----+----+
	| α   | -β | -γ | -δ | 0  | 0  | 0  | 0  |
	+-----+----+----+----+----+----+----+----+
	| β   | +α | -δ | +γ | 0  | 0  | 0  | 0  |
	+-----+----+----+----+----+----+----+----+
	| γ   | +δ | +α | -β | 0  | 0  | 0  | 0  |
	+-----+----+----+----+----+----+----+----+
	| δ   | -γ | +β | +α | 0  | 0  | 0  | 0  |
	+-----+----+----+----+----+----+----+----+
The multiplcation operation for infra-Hamilton quaternions is non-commutative and
non-associative.

Infra-Hamilton quaternions are a parabolic Cayley-Dickson construct with Hamilton
quaternions.
*/
package infrahamilton

const (
	leftBracket            = "⦗"
	rightBracket           = "⦘"
	zeroDivisorDenominator = "denominator is zero divisor"
	zeroDivisorInverse     = "inverse of zero divisor"
	unit1                  = "i"
	unit2                  = "j"
	unit3                  = "k"
	unit4                  = "α"
	unit5                  = "β"
	unit6                  = "γ"
	unit7                  = "δ"
)

var (
	unitNames = [7]string{unit1, unit2, unit3, unit4, unit5, unit6, unit7}
)

// ResetUnitNames sets the names of the infra-Hamilton quaternion units equal
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

// SetUnitNames sets the names of the infra-Hamilton quaternion units.
func SetUnitNames(u1, u2, u3, u4, u5, u6, u7 string) {
	unitNames[0] = u1
	unitNames[1] = u2
	unitNames[2] = u3
	unitNames[3] = u4
	unitNames[4] = u5
	unitNames[5] = u6
	unitNames[6] = u7
}

// UnitNames returns the current names of the infra-Hamilton quaternion units.
func UnitNames() [7]string {
	return unitNames
}
