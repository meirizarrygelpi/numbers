# bipplex

Package `bipplex` implements arithmetic for bi-perplex numbers. There are five types:

* `Int64`
* `Float64`
* `Int`
* `Float`
* `Rat`

Each bi-perplex value is printed in the form "⦗a+bs+cT+dsT⦘". This is similar to `complex128` values, but uses tortoise shell brackets to distinguish.

[![Go Report Card](https://goreportcard.com/badge/gojp/goreportcard)](https://goreportcard.com/report/github.com/meirizarrygelpi/numbers/bipplex) [![GoDoc](https://godoc.org/github.com/meirizarrygelpi/numbers/bipplex?status.svg)](https://godoc.org/github.com/meirizarrygelpi/numbers/bipplex)

## To-Do

1. Inf and NaN for `bipplex.Float` and `bipplex.Float64` types.
2. Fix `String` method for float types.