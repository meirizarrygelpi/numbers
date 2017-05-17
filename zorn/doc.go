// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

/*
Package zorn implements arithmetic for Zorn octonions. Five types are
implemented:
    Int64   (int64 values)
    Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
A Zorn octonion has eight components and it is written in the form
    a+bi+cj+dk+er+fs+gt+hu
The multiplication table is
	+-----+----+----+----+----+----+----+----+
	| Mul | i  | j  | k  | r  | s  | t  | u  |
	+-----+----+----+----+----+----+----+----+
	| i   | -1 | +k | -j | +s | -r | -u | +t |
	+-----+----+----+----+----+----+----+----+
	| j   | -k | -1 | +i | +t | +u | -r | -s |
	+-----+----+----+----+----+----+----+----+
	| k   | +j | -i | -1 | +u | -t | +s | -r |
	+-----+----+----+----+----+----+----+----+
	| r   | -s | -t | -u | +1 | -i | -j | -k |
	+-----+----+----+----+----+----+----+----+
	| s   | +r | -u | +t | +i | +1 | +k | -j |
	+-----+----+----+----+----+----+----+----+
	| t   | +u | +r | -s | +j | -k | +1 | +i |
	+-----+----+----+----+----+----+----+----+
	| u   | -t | +s | +r | +k | +j | -i | +1 |
	+-----+----+----+----+----+----+----+----+
The multiplcation operation for Zorn octonions is non-commutative and
non-associative.

Zorn octonions are a hyperbolic Cayley-Dickson construct with Hamilton
quaternions.
*/
package zorn

const (
	leftBracket            = "⦗"
	rightBracket           = "⦘"
	zeroDivisorDenominator = "denominator is zero divisor"
	zeroDivisorInverse     = "inverse of zero divisor"
	unit1                  = "i"
	unit2                  = "j"
	unit3                  = "k"
	unit4                  = "r"
	unit5                  = "s"
	unit6                  = "t"
	unit7                  = "u"
)

var (
	unitNames = [7]string{unit1, unit2, unit3, unit4, unit5, unit6, unit7}
)

// ResetUnitNames sets the names of the Zorn octonion units equal to their
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

// SetUnitNames sets the names of the Zorn octonion units.
func SetUnitNames(u1, u2, u3, u4, u5, u6, u7 string) {
	unitNames[0] = u1
	unitNames[1] = u2
	unitNames[2] = u3
	unitNames[3] = u4
	unitNames[4] = u5
	unitNames[5] = u6
	unitNames[6] = u7
}

// UnitNames returns the current names of the Zorn octonion units.
func UnitNames() [7]string {
	return unitNames
}
