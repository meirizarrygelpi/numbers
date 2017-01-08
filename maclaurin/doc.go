// Copyright (c) 2017 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

/*
Package maclaurin implements univariate Maclaurin polynomials. Five
types are implemented:
    Int64   (int64 values)
    Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
A Maclaurin polynomial is a sum of monomials with non-negative degree.
*/
package maclaurin

func merge(a, b []uint64) []uint64 {
	la, lb := len(a), len(b)
	i, j := 0, 0
	n := la + lb
	c := make([]uint64, n)
	for k := 0; k < n; k++ {
		if (i < la) && (j < lb) {
			if a[i] < b[j] {
				c[k] = a[i]
				i++
			} else {
				c[k] = b[j]
				j++
			}
		} else if i > la-1 {
			c[k] = b[j]
			j++
		} else {
			c[k] = a[i]
			i++
		}
	}
	return c
}

func sort(x []uint64) []uint64 {
	n := len(x)
	if (n == 0) || (n == 1) {
		return x
	}
	h := n / 2
	return merge(sort(x[:h]), sort(x[h:]))
}

func reverse(x []uint64) []uint64 {
	n := len(x)
	y := make([]uint64, n)
	for i := 0; i < n; i++ {
		y[i] = x[n-i-1]
	}
	return y
}
