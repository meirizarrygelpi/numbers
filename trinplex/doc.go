// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

/*
Package trinplex implements arithmetic for tri-nilplex numbers. Five types are
implemented:
    Int64   (int64 values)
    Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
A tri-nilplex number has eight components and it is written in the form
    a+bα+cΓ+dαΓ+eΛ+fαΛ+gΓΛ+hαΓΛ
The multiplication table is:
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | Mul | α   | Γ   | αΓ  | Λ   | αΛ  | ΓΛ  | αΓΛ |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | α   | 0   | αΓ  | 0   | αΛ  | 0   | αΓΛ | 0   |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | Γ   | αΓ  | 0   | 0   | ΓΛ  | αΓΛ | 0   | 0   |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | αΓ  | 0   | 0   | 0   | αΓΛ | 0   | 0   | 0   |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | Λ   | αΛ  | ΓΛ  | αΓΛ | 0   | 0   | 0   | 0   |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | αΛ  | 0   | αΓΛ | 0   | 0   | 0   | 0   | 0   |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | ΓΛ  | αΓΛ | 0   | 0   | 0   | 0   | 0   | 0   |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | αΓΛ | 0   | 0   | 0   | 0   | 0   | 0   | 0   |
    +-----+-----+-----+-----+-----+-----+-----+-----+
The multiplcation operation for tri-nilplex numbers is commutative and
associative (for non-floats).

Tri-nilplex numbers are a nilplexification of hyper numbers.
*/
package trinplex

const (
	leftBracket            = "⦗"
	rightBracket           = "⦘"
	zeroDivisorDenominator = "denominator is zero divisor"
	zeroDivisorInverse     = "inverse of zero divisor"
	unit1                  = "α"
	unit2                  = "Γ"
	unit3                  = "αΓ"
	unit4                  = "Λ"
	unit5                  = "αΛ"
	unit6                  = "ΓΛ"
	unit7                  = "αΓΛ"
)
