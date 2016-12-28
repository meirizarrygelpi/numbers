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
A hyper-complex number has eight components and it is written in the form
    a+bi+cΓ+diΓ+eΛ+fiΛ+gΓΛ+hiΓΛ
The multiplication table is:
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | Mul | i   | Γ   | iΓ  | Λ   | iΛ  | ΓΛ  | iΓΛ |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | i   | -1  | iΓ  | -Γ  | iΛ  | -Λ  | iΓΛ | -ΓΛ |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | Γ   | iΓ  | 0   | 0   | ΓΛ  | iΓΛ | 0   | 0   |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | iΓ  | -Γ  | 0   | 0   | iΓΛ | -ΓΛ | 0   | 0   |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | Λ   | iΛ  | ΓΛ  | iΓΛ | 0   | 0   | 0   | 0   |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | iΛ  | -Λ  | iΓΛ | -ΓΛ | 0   | 0   | 0   | 0   |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | ΓΛ  | iΓΛ | 0   | 0   | 0   | 0   | 0   | 0   |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | iΓΛ | -ΓΛ | 0   | 0   | 0   | 0   | 0   | 0   |
    +-----+-----+-----+-----+-----+-----+-----+-----+
The multiplcation operation for hyper-complex numbers is commutative and
associative (for non-floats).

Hyper-complex numbers are a nilplexification of dual-complex numbers.
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
