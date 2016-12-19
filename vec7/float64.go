package vec7

import (
	"fmt"
	"strings"
)

// A Float64 is a seven-dimensional vector of float64 values.
type Float64 [7]float64

// NewFloat64 returns the pointer to an Float64 with given components.
func NewFloat64(a, b, c, d, e, f, g float64) *Float64 {
	v := new(Float64)
	v[0] = a
	v[1] = b
	v[2] = c
	v[3] = d
	v[4] = e
	v[5] = f
	v[6] = g
	return v
}

// String returns the string representation of a Float64 as a sequence of
// values separated by commas, and surrounded by double angle brackets.
func (v *Float64) String() string {
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
