# supra

Package `supra` implements arithmetic for supra numbers. There are five types:

* `Int64`
* `Float64`
* `Int`
* `Float`
* `Rat`

Each supra number value is printed in the form "⦗a+bα+cβ+dγ⦘". This is similar to `complex128` values, but uses tortoise shell brackets to distinguish.

[![Go Report Card](https://goreportcard.com/badge/gojp/goreportcard)](https://goreportcard.com/report/github.com/meirizarrygelpi/numbers/supra) [![GoDoc](https://godoc.org/github.com/meirizarrygelpi/numbers/supra?status.svg)](https://godoc.org/github.com/meirizarrygelpi/numbers/supra)

## To-Do

1. Inf and NaN for `supra.Float` and `supra.Float64` types.
2. Fix `String` method for float types.