# nplex

Package `nplex` implements arithmetic for nilplex numbers. There are five types:

* `Int64`
* `Float64`
* `Int`
* `Float`
* `Rat`

Each nilplex value is printed in the form "⦗a+bα⦘". This is similar to `complex128` values, but uses tortoise shell brackets to distinguish. The nilplex numbers are more commonly known as dual numbers.

[![Go Report Card](https://goreportcard.com/badge/gojp/goreportcard)](https://goreportcard.com/report/github.com/meirizarrygelpi/numbers/nplex) [![GoDoc](https://godoc.org/github.com/meirizarrygelpi/numbers/nplex?status.svg)](https://godoc.org/github.com/meirizarrygelpi/numbers/nplex)

## To-Do

1. Inf and NaN for `nplex.Float` and `nplex.Float64` types.
2. Fix `String` method for float types.