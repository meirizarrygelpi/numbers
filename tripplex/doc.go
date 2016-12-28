// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

/*
Package tripplex implements arithmetic for tri-perplex numbers. Five types are
implemented:
    Int64   (int64 values)
    Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
A tri-perplex number has eight components and it is written in the form
    a+bs+cT+dsT+eU+fsU+gTU+hsTU
The multiplication table is:
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | Mul | s   | T   | sT  | U   | sU  | TU  | sTU |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | s   | 1   | sT  | T   | sU  | U   | sTU | TU  |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | T   | sT  | 1   | s   | TU  | sTU | U   | sU  |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | sT  | T   | s   | 1   | sTU | TU  | sU  | U   |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | U   | sU  | TU  | sTU | 1   | s   | T   | sT  |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | sU  | U   | sTU | TU  | s   | 1   | sT  | T   |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | TU  | sTU | U   | sU  | T   | sT  | 1   | s   |
    +-----+-----+-----+-----+-----+-----+-----+-----+
    | sTU | TU  | sU  | U   | sT  | T   | s   | 1   |
    +-----+-----+-----+-----+-----+-----+-----+-----+
The multiplcation operation for tri-perplex numbers is commutative and
associative (for non-floats).

Tri-perplex numbers are a perplexification of the bi-perplex numbers.
*/
package tripplex

const (
	leftBracket            = "⦗"
	rightBracket           = "⦘"
	zeroDivisorDenominator = "denominator is zero divisor"
	zeroDivisorInverse     = "inverse of zero divisor"
	unit1                  = "s"
	unit2                  = "T"
	unit3                  = "sT"
	unit4                  = "U"
	unit5                  = "sU"
	unit6                  = "TU"
	unit7                  = "sTU"
)
