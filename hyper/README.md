# hyper

Package `hyper` implements arithmetic for hyper numbers. There are five types:

* `Int64`
* `Float64`
* `Int`
* `Float`
* `Rat`

Each hyper value is printed in the form "⦗a+bα+cΓ+dαΓ⦘". This is similar to `complex128` values, but uses tortoise shell brackets to distinguish.

[![Go Report Card](https://goreportcard.com/badge/gojp/goreportcard)](https://goreportcard.com/report/github.com/meirizarrygelpi/numbers/hyper) [![GoDoc](https://godoc.org/github.com/meirizarrygelpi/numbers/hyper?status.svg)](https://godoc.org/github.com/meirizarrygelpi/numbers/hyper)

## To-Do

1. Inf and NaN for `hyper.Float` and `hyper.Float64` types.
2. Fix `String` method for float types.