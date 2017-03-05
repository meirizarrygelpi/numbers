// Copyright (c) 2017 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

/*
Package heegner implements arithmetic for imaginary quadratic fields with
class number equal to one. Nine types are implemented:
    Rat1        (big.Rat values with adjoined √−1)
    Rat2        (big.Rat values with adjoined √−2)
    Rat3        (big.Rat values with adjoined √−3)
    Rat7        (big.Rat values with adjoined √−7)
    Rat11       (big.Rat values with adjoined √−11)
    Rat19       (big.Rat values with adjoined √−19)
    Rat43       (big.Rat values with adjoined √−43)
    Rat67       (big.Rat values with adjoined √−67)
    Rat163      (big.Rat values with adjoined √−163)
An element of one of these imaginary quadratic fields has two components
and it is written in the form
    a+b√−d
where d is one of the nine Heegner number. The multiplcation operation for
these elements is commutative and associative, since they are also complex
numbers. Prime factorization is unique in each of these imaginary quadratic
fields.
*/
package heegner

import "math/big"

const (
	leftBracket     = "⦗"
	rightBracket    = "⦘"
	zeroDenominator = "denominator is zero"
	zeroInverse     = "inverse of zero"
	radical1        = "√₋₁"
	radical2        = "√₋₂"
	radical3        = "√₋₃"
	radical7        = "√₋₇"
	radical11       = "√₋₁₁"
	radical19       = "√₋₁₉"
	radical43       = "√₋₄₃"
	radical67       = "√₋₆₇"
	radical163      = "√₋₁₆₃"
)

var (
	h2   = big.NewRat(2, 1)
	h3   = big.NewRat(3, 1)
	h7   = big.NewRat(7, 1)
	h11  = big.NewRat(11, 1)
	h19  = big.NewRat(19, 1)
	h43  = big.NewRat(43, 1)
	h67  = big.NewRat(67, 1)
	h163 = big.NewRat(163, 1)
)
