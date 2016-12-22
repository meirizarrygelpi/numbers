# infrapplex

Package `infrapplex` implements arithmetic for infra-perplex numbers. There are five types:

* `Int64`
* `Float64`
* `Int`
* `Float`
* `Rat`

Each infra-perplex number value is printed in the form "⦗a+bs+cτ+dυ⦘". This is similar to `complex128` values, but uses tortoise shell brackets to distinguish.

[![Go Report Card](https://goreportcard.com/badge/gojp/goreportcard)](https://goreportcard.com/report/github.com/meirizarrygelpi/numbers/infrapplex) [![GoDoc](https://godoc.org/github.com/meirizarrygelpi/numbers/infrapplex?status.svg)](https://godoc.org/github.com/meirizarrygelpi/numbers/infrapplex)

## To-Do

1. Inf and NaN for `infrapplex.Float` and `infrapplex.Float64` types.
2. Fix `String` method for float types.