# dualcplex

Package `dualcplex` implements arithmetic for dual-complex numbers. There are five types:

* `Int64`
* `Float64`
* `Int`
* `Float`
* `Rat`

Each dual-complex value is printed in the form "⦗a+bi+cΓ+diΓ⦘". This is similar to `complex128` values, but uses tortoise shell brackets to distinguish.

[![Go Report Card](https://goreportcard.com/badge/gojp/goreportcard)](https://goreportcard.com/report/github.com/meirizarrygelpi/numbers/dualcplex) [![GoDoc](https://godoc.org/github.com/meirizarrygelpi/numbers/dualcplex?status.svg)](https://godoc.org/github.com/meirizarrygelpi/numbers/dualcplex)

## To-Do

1. Inf and NaN for `dualcplex.Float` and `dualcplex.Float64` types.
2. Fix `String` method for float types.