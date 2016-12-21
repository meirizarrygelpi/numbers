# infracplex

Package `infracplex` implements arithmetic for infra-complex numbers. There are five types:

* `Int64`
* `Float64`
* `Int`
* `Float`
* `Rat`

Each infra-complex number value is printed in the form "⦗a+bi+cβ+dγ⦘". This is similar to `complex128` values, but uses tortoise shell brackets to distinguish.

[![Go Report Card](https://goreportcard.com/badge/gojp/goreportcard)](https://goreportcard.com/report/github.com/meirizarrygelpi/numbers/infracplex) [![GoDoc](https://godoc.org/github.com/meirizarrygelpi/numbers/infracplex?status.svg)](https://godoc.org/github.com/meirizarrygelpi/numbers/infracplex)

## To-Do

1. Inf and NaN for `infracplex.Float` and `infracplex.Float64` types.
2. Fix `String` method for float types.