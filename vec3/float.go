package vec3

import (
	"math/big"
	"strings"
)

// A Float is a three-dimensional vector of big.Float pointers.
type Float [3]*big.Float

// NewFloat returns the pointer to a Float with given components.
func NewFloat(a, b, c *big.Float) *Float {
	x := new(Float)
	x[0] = a
	x[1] = b
	x[2] = c
	return x
}

// String returns the string representation of a Float as a sequence of
// values separated by commas, and surrounded by double angle brackets.
func (v *Float) String() string {
	var x [7]string
	x[0] = leftBracket
	x[1] = v[0].String()
	x[2] = commaSpace
	x[3] = v[1].String()
	x[4] = commaSpace
	x[5] = v[2].String()
	x[6] = rightBracket
	return strings.Join(x[:], "")
}
