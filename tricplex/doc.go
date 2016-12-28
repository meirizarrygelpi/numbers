// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

/*
Package tricplex implements arithmetic for tri-complex numbers. Five types are
implemented:
    Int64   (int64 values)
    Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
A tri-complex number has eight components and it is written in the form
    a+bi+cJ+diJ+eK+fiK+gJK+hiJK
The multiplication table is:
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | Mul | i   | J   | iJ  | K   | iK  | JK  | iJK |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | i   | -1  | iJ  | -J  | iK  | -K  | iJK | -JK |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | J   | iJ  | -1  | -i  | JK  | iJK | -K  | -iK |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | iJ  | -J  | -i  | +1  | iJK | -JK | -iK | +K  |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | K   | iK  | JK  | iJK | -1  | -i  | -J  | -iJ |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | iK  | -K  | iJK | -JK | -i  | +1  | -iJ | +J  |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | JK  | iJK | -K  | -iK | -J  | -iK | +1  | +i  |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | iJK | -JK | -iK | +K  | -iJ | +J  | +i  | -1  |
    +-----+-----+-----+-----+-----+-----+-----+-----+
The multiplcation operation for tri-complex numbers is commutative and
associative (for non-floats).

Tri-complex numbers are a complexification of bi-complex numbers.
*/
package tricplex

const (
	leftBracket            = "⦗"
	rightBracket           = "⦘"
	zeroDivisorDenominator = "denominator is zero divisor"
	zeroDivisorInverse     = "inverse of zero divisor"
	unit1                  = "i"
	unit2                  = "J"
	unit3                  = "iJ"
	unit4                  = "K"
	unit5                  = "iK"
	unit6                  = "JK"
	unit7                  = "iJK"
)
