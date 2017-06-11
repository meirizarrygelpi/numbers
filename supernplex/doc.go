// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

/*
Package supernplex implements arithmetic for super-nilplex numbers. Five types
are implemented:
    Int64   (int64 values)
    Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
A super-nilplex number has four components and it is written in the form
    a+bα+cβ+dγ
The multiplication table is:
    +-----+----+----+---+
    | Mul | α  | β  | γ |
    +-----+----+----+---+
    | α   | 0  | +γ | 0 |
    +-----+----+----+---+
    | β   | -γ | 0  | 0 |
    +-----+----+----+---+
    | γ   | 0  | 0  | 0 |
    +-----+----+----+---+
The multiplcation operation for super-nilplex numbers is non-commutative but
associative (for non-floats).

Super-nilplex numbers are a parabolic Cayley-Dickson construct with nilplex
numbers. Another interpretation of super-nilplex numbers is as a two-dimensional
exterior algebra (a.k.a. Grassmann algebra) with the multiplication being the
wedge product.
*/
package supernplex

const (
	leftBracket            = "⦗"
	rightBracket           = "⦘"
	zeroDivisorDenominator = "denominator is zero divisor"
	zeroDivisorInverse     = "inverse of zero divisor"
	unit1                  = "α"
	unit2                  = "β"
	unit3                  = "γ"
)

var (
	unitNames = [3]string{unit1, unit2, unit3}
)

// ResetUnitNames sets the names of the super-nilplex units equal to their
// default values.
func ResetUnitNames() {
	unitNames[0] = unit1
	unitNames[1] = unit2
	unitNames[2] = unit3
}

// SetUnitNames sets the names of the super-nilplex units.
func SetUnitNames(u1, u2, u3 string) {
	unitNames[0] = u1
	unitNames[1] = u2
	unitNames[2] = u3
}

// UnitNames returns the current names of the super-nilplex units.
func UnitNames() [3]string {
	return unitNames
}
