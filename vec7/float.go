package vec7

import (
	"math/big"
	"strings"
)

// A Float is a seven-dimensional vector of big.Float pointers.
type Float [7]*big.Float

// NewFloat returns the pointer to an Float with given components.
func NewFloat(a, b, c, d, e, f, g *big.Float) *Float {
	v := new(Float)
	v[0] = a
	v[1] = b
	v[2] = c
	v[3] = d
	v[4] = e
	v[5] = f
	v[6] = g
	return v
}

// String returns the string representation of a Float as a sequence of
// values separated by commas, and surrounded by double angle brackets.
func (v *Float) String() string {
	var x [15]string
	x[0] = leftBracket
	x[1] = v[0].String()
	x[2] = commaSpace
	x[3] = v[1].String()
	x[4] = commaSpace
	x[5] = v[2].String()
	x[6] = commaSpace
	x[7] = v[3].String()
	x[8] = commaSpace
	x[9] = v[4].String()
	x[10] = commaSpace
	x[11] = v[5].String()
	x[12] = commaSpace
	x[13] = v[6].String()
	x[14] = rightBracket
	return strings.Join(x[:], "")
}
