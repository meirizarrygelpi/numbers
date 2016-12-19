package vec3

import (
	"fmt"
	"strings"
)

// An Int64 is a three-dimensional vector of int64 values.
type Int64 [3]int64

// NewInt64 returns the pointer to an Int64 with given components.
func NewInt64(a, b, c int64) *Int64 {
	x := new(Int64)
	x[0] = a
	x[1] = b
	x[2] = c
	return x
}

// String returns the string representation of an Int64 as a sequence of
// values separated by commas, and surrounded by double angle brackets.
func (v *Int64) String() string {
	var x [7]string
	x[0] = leftBracket
	x[1] = fmt.Sprint(v[0])
	x[2] = commaSpace
	x[3] = fmt.Sprint(v[1])
	x[4] = commaSpace
	x[5] = fmt.Sprint(v[2])
	x[6] = rightBracket
	return strings.Join(x[:], "")
}
