# pplex

Package `pplex` implements arithmetic for perplex numbers. There are five types:

* `Int64`
* `Float64`
* `Int`
* `Float`
* `Rat`

Each perplex value is printed in the form "⦗a+bs⦘". This is similar to `complex128` values, but uses tortoise shell brackets to distinguish.

[![Go Report Card](https://goreportcard.com/badge/gojp/goreportcard)](https://goreportcard.com/report/github.com/meirizarrygelpi/numbers/pplex) [![GoDoc](https://godoc.org/github.com/meirizarrygelpi/numbers/pplex?status.svg)](https://godoc.org/github.com/meirizarrygelpi/numbers/pplex)

## To-Do

1. Inf and NaN for `pplex.Float` and `pplex.Float64` types.
2. Fix `String` method for float types.