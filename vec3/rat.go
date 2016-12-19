package vec3

import (
	"math/big"
	"strings"
)

// A Rat is a three-dimensional vector of big.Rat pointers.
type Rat [3]*big.Rat

// NewRat returns the pointer to a Rat with given components.
func NewRat(a, b, c *big.Rat) *Rat {
	x := new(Rat)
	x[0] = a
	x[1] = b
	x[2] = c
	return x
}

// String returns the string representation of a Rat as a sequence of
// values separated by commas, and surrounded by double angle brackets.
func (v *Rat) String() string {
	var x [7]string
	x[0] = leftBracket
	x[1] = v[0].RatString()
	x[2] = commaSpace
	x[3] = v[1].RatString()
	x[4] = commaSpace
	x[5] = v[2].RatString()
	x[6] = rightBracket
	return strings.Join(x[:], "")
}
