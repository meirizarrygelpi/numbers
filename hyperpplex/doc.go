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
A hyper-perplex number has eight components and it is written in the form
    a+bs+cΓ+dsΓ+eΛ+fsΛ+gΓΛ+hsΓΛ
The multiplication table is:
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | Mul | s   | Γ   | sΓ  | Λ   | sΛ  | ΓΛ  | sΓΛ |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | s   | 1   | sΓ  | Γ   | sΛ  | Λ   | sΓΛ | ΓΛ  |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | Γ   | sΓ  | 0   | 0   | ΓΛ  | sΓΛ | 0   | 0   |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | sΓ  | Γ   | 0   | 0   | sΓΛ | ΓΛ  | 0   | 0   |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | Λ   | sΛ  | ΓΛ  | sΓΛ | 0   | 0   | 0   | 0   |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | sΛ  | Λ   | sΓΛ | ΓΛ  | 0   | 0   | 0   | 0   |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | ΓΛ  | sΓΛ | 0   | 0   | 0   | 0   | 0   | 0   |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | sΓΛ | ΓΛ  | 0   | 0   | 0   | 0   | 0   | 0   |
    +-----+-----+-----+-----+-----+-----+-----+-----+
The multiplcation operation for hyper-perplex numbers is commutative and
associative (for non-floats).

Hyper-perplex numbers are a nilplexification of dual-perplex numbers.
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
