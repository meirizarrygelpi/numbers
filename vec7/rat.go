package vec7

import (
	"math/big"
	"strings"
)

// A Rat is a seven-dimensional vector of big.Rat pointers.
type Rat [7]*big.Rat

// NewRat returns the pointer to an Rat with given components.
func NewRat(a, b, c, d, e, f, g *big.Rat) *Rat {
	v := new(Rat)
	v[0] = a
	v[1] = b
	v[2] = c
	v[3] = d
	v[4] = e
	v[5] = f
	v[6] = g
	return v
}

// String returns the string representation of a Rat as a sequence of
// values separated by commas, and surrounded by double angle brackets.
func (v *Rat) String() string {
	var x [15]string
	x[0] = leftBracket
	x[1] = v[0].RatString()
	x[2] = commaSpace
	x[3] = v[1].RatString()
	x[4] = commaSpace
	x[5] = v[2].RatString()
	x[6] = commaSpace
	x[7] = v[3].RatString()
	x[8] = commaSpace
	x[9] = v[4].RatString()
	x[10] = commaSpace
	x[11] = v[5].RatString()
	x[12] = commaSpace
	x[13] = v[6].RatString()
	x[14] = rightBracket
	return strings.Join(x[:], "")
}
