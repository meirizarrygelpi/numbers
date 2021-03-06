package vec7

import (
	"fmt"
	"strings"
)

// An Int64 is a seven-dimensional vector of int64 values.
type Int64 [7]int64

// NewInt64 returns the pointer to an Int64 with given components.
func NewInt64(a, b, c, d, e, f, g int64) *Int64 {
	v := new(Int64)
	v[0] = a
	v[1] = b
	v[2] = c
	v[3] = d
	v[4] = e
	v[5] = f
	v[6] = g
	return v
}

// String returns the string representation of an Int64 as a sequence of
// values separated by commas, and surrounded by double angle brackets.
func (v *Int64) String() string {
	var x [15]string
	x[0] = leftBracket
	x[1] = fmt.Sprint(v[0])
	x[2] = commaSpace
	x[3] = fmt.Sprint(v[1])
	x[4] = commaSpace
	x[5] = fmt.Sprint(v[2])
	x[6] = commaSpace
	x[7] = fmt.Sprint(v[3])
	x[8] = commaSpace
	x[9] = fmt.Sprint(v[4])
	x[10] = commaSpace
	x[11] = fmt.Sprint(v[5])
	x[12] = commaSpace
	x[13] = fmt.Sprint(v[6])
	x[14] = rightBracket
	return strings.Join(x[:], "")
}
