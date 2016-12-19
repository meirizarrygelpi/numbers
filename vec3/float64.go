package vec3

import (
	"fmt"
	"strings"
)

// A Float64 is a three-dimensional vector of float64 values.
type Float64 [3]float64

// NewFloat64 returns the pointer to a Float64 with given components.
func NewFloat64(a, b, c float64) *Float64 {
	x := new(Float64)
	x[0] = a
	x[1] = b
	x[2] = c
	return x
}

// String returns the string representation of a Float64 as a sequence of
// values separated by commas, and surrounded by double angle brackets.
func (v *Float64) String() string {
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
