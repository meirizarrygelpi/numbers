package vec3

import (
	"math/big"
	"strings"
)

// An Int is a three-dimensional vector of big.Int pointers.
type Int [3]*big.Int

// NewInt returns the pointer to an Int with given components.
func NewInt(a, b, c *big.Int) *Int {
	x := new(Int)
	x[0] = a
	x[1] = b
	x[2] = c
	return x
}

// String returns the string representation of an Int as a sequence of
// values separated by commas, and surrounded by double angle brackets.
func (v *Int) String() string {
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
