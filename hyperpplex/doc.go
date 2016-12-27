// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

/*
Package hyperpplex implements arithmetic for hyper-perplex numbers. Five types are
implemented:
    Int64   (int64 values)
    Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
The multiplcation operation for hyper-perplex numbers is commutative and
associative (for non-floats).
*/
package hyperpplex

const (
	leftBracket            = "⦗"
	rightBracket           = "⦘"
	zeroDivisorDenominator = "denominator is zero divisor"
	zeroDivisorInverse     = "inverse of zero divisor"
	unit1                  = "s"
	unit2                  = "Γ"
	unit3                  = "sΓ"
	unit4                  = "Λ"
	unit5                  = "sΛ"
	unit6                  = "ΓΛ"
	unit7                  = "sΓΛ"
)
