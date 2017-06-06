// Copyright (c) 2017 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

/*
Package comhamilton implements arithmetic for complex-Hamilton quaternions. Five types are
implemented:
    Int64   (int64 values)
    Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
An complex-Hamilton quaternion has eight components and it is written in the form
    a+bi+cj+dk+eH+fiH+gjH+hkH
The multiplication table is
	+-----+-----+-----+-----+-----+-----+-----+-----+
	| Mul | i   | j   | k   | H   | iH  | jH  | kH  |
	+-----+-----+-----+-----+-----+-----+-----+-----+
	| i   | -1  | +k  | -j  | +iH | -H  | +kH | -jH |
	+-----+-----+-----+-----+-----+-----+-----+-----+
	| j   | -k  | -1  | +i  | +jH | -kH | -H  | +iH |
	+-----+-----+-----+-----+-----+-----+-----+-----+
	| k   | +j  | -i  | -1  | +kH | +jH | -iH | -H  |
	+-----+-----+-----+-----+-----+-----+-----+-----+
	| H   | +iH | +jH | +kH | -1  | -i  | -j  | -k  |
	+-----+-----+-----+-----+-----+-----+-----+-----+
	| iH  | -H  | +kH | -jH | -i  | +1  | -k  | +j  |
	+-----+-----+-----+-----+-----+-----+-----+-----+
	| jH  | -kH | -H  | +iH | -j  | +k  | +1  | -i  |
	+-----+-----+-----+-----+-----+-----+-----+-----+
	| kH  | +jH | -iH | -H  | -k  | -j  | +i  | +1  |
	+-----+-----+-----+-----+-----+-----+-----+-----+
The multiplcation operation for complex-Hamilton quaternions is non-commutative but
associative (for non-floats).

Complex-Hamilton quaternions are a complexification of Hamilton quaternions.
*/
package comhamilton

const (
	leftBracket            = "⦗"
	rightBracket           = "⦘"
	zeroDivisorDenominator = "denominator is zero divisor"
	zeroDivisorInverse     = "inverse of zero divisor"
	unit1                  = "i"
	unit2                  = "j"
	unit3                  = "k"
	unit4                  = "H"
	unit5                  = "iH"
	unit6                  = "jH"
	unit7                  = "kH"
)

var (
	unitNames = [7]string{unit1, unit2, unit3, unit4, unit5, unit6, unit7}
)

// ResetUnitNames sets the names of the complex-Hamilton quaternion units equal
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

// SetUnitNames sets the names of the complex-Hamilton quaternion units.
func SetUnitNames(u1, u2, u3, u4, u5, u6, u7 string) {
	unitNames[0] = u1
	unitNames[1] = u2
	unitNames[2] = u3
	unitNames[3] = u4
	unitNames[4] = u5
	unitNames[5] = u6
	unitNames[6] = u7
}

// UnitNames returns the current names of the complex-Hamilton quaternion units.
func UnitNames() [7]string {
	return unitNames
}
