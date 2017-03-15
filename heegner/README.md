# heegner

Package `heegner` implements arithmetic for imaginary quadratic fields with class number equal to one. There are nine types:

* `Rat1`
* `Rat2`
* `Rat3`
* `Rat7`
* `Rat11`
* `Rat19`
* `Rat43`
* `Rat67`
* `Rat163`

Each value is printed in the form "⦗a+bX⦘" where `X` is the square root of the negative of one of the nine Heegner numbers. This is similar to `complex128` values, but uses tortoise shell brackets to distinguish.

[![Go Report Card](https://goreportcard.com/badge/gojp/goreportcard)](https://goreportcard.com/report/github.com/meirizarrygelpi/numbers/heegner) [![GoDoc](https://godoc.org/github.com/meirizarrygelpi/numbers/heegner?status.svg)](https://godoc.org/github.com/meirizarrygelpi/numbers/heegner)