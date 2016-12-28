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
    ...
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
