// Copyright (c) 2017 Melvin Eloy Irizarry-Gelpí
// Licensed under the MIT License.

package maclaurin

import "sort"

// A Degrees is a slice of uint64 representing the degrees of a Maclaurin
// polynomial.
type Degrees []uint64

// Len is part of sort.Interface.
func (d Degrees) Len() int {
	return len(d)
}

// Swap is part of sort.Interface.
func (d Degrees) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

// Less is part of sort.Interface.
func (d Degrees) Less(i, j int) bool {
	return d[i] < d[j]
}

// Sort sorts the slice of degrees.
func (d Degrees) Sort() {
	sort.Sort(d)
}

// ReverseSort reverse-sorts the slice of degrees.
func (d Degrees) ReverseSort() {
	sort.Sort(sort.Reverse(d))
}
