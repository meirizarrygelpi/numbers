// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

/*
Package cayley implements arithmetic for Cayley octonions. Five types are
implemented:
    Int64   (int64 values)
    Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
A Cayley octonion has eight components and it is written in the form
    a+bi+cj+dk+em+fn+gp+hq
The multiplication table is:
    +-----+----+----+----+----+----+----+----+
    | Mul | i  | j  | k  | m  | n  | p  | q  |
    +-----+----+----+----+----+----+----+----+
    | i   | -1 | +k | -j | +n | -m | -q | +p |
    +-----+----+----+----+----+----+----+----+
    | j   | -k | -1 | +i | +p | +q | -m | -n |
    +-----+----+----+----+----+----+----+----+
    | k   | +j | -i | -1 | +q | -p | +n | -m |
    +-----+----+----+----+----+----+----+----+
    | m   | -n | -p | -q | -1 | +i | +j | +k |
    +-----+----+----+----+----+----+----+----+
    | n   | +m | -q | +p | -i | -1 | -k | +j |
    +-----+----+----+----+----+----+----+----+
    | p   | +q | +m | -n | -j | +k | -1 | -i |
    +-----+----+----+----+----+----+----+----+
    | q   | -p | +n | +m | -k | -j | +i | -1 |
    +-----+----+----+----+----+----+----+----+
The multiplcation operation for Cayley octonions is non-commutative and
non-associative.

Cayley octonions are an elliptic Cayley-Dickson construct with Hamilton
quaternions.
*/
package cayley

const (
	leftBracket     = "⦗"
	rightBracket    = "⦘"
	zeroDenominator = "denominator is zero"
	zeroInverse     = "inverse of zero"
	unit1           = "i"
	unit2           = "j"
	unit3           = "k"
	unit4           = "m"
	unit5           = "n"
	unit6           = "p"
	unit7           = "q"
)
