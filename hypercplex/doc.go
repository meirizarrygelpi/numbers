// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

/*
Package hypercplex implements arithmetic for hyper-complex numbers. Five types are
implemented:
    Int64   (int64 values)
    Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
The multiplcation operation for hyper-complex numbers is commutative and
associative (for non-floats).
*/
package hypercplex

const (
	leftBracket            = "⦗"
	rightBracket           = "⦘"
	zeroDivisorDenominator = "denominator is zero divisor"
	zeroDivisorInverse     = "inverse of zero divisor"
	unit1                  = "i"
	unit2                  = "Γ"
	unit3                  = "iΓ"
	unit4                  = "Λ"
	unit5                  = "iΛ"
	unit6                  = "ΓΛ"
	unit7                  = "iΓΛ"
)
