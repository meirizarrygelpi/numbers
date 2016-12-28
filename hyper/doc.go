// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

/*
Package hyper implements arithmetic for hyper numbers. Five types are
implemented:
    Int64   (int64 values)
    Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
A hyper number has four components and it is written in the form
    a+bα+cΓ+dαΓ
The multiplication table is:
    +-----+----+----+----+
    | Mul | α  | Γ  | αΓ |
    +-----+----+----+----+
    | α   | 0  | αΓ | 0  |
    +-----+----+----+----+
    | Γ   | αΓ | 0  | 0  |
    +-----+----+----+----+
    | αΓ  | 0  | 0  | 0  |
    +-----+----+----+----+
The multiplcation operation for hyper numbers is commutative and
associative (for non-floats).

Hyper numbers are a nilplexification of nilplex numbers.
*/
package hyper

const (
	leftBracket            = "⦗"
	rightBracket           = "⦘"
	zeroDivisorDenominator = "denominator is zero divisor"
	zeroDivisorInverse     = "inverse of zero divisor"
	unit1                  = "α"
	unit2                  = "Γ"
	unit3                  = "αΓ"
)
