# cayley

Package `cayley` implements arithmetic for Cayley octonions. There are five types:

* `Int64`
* `Float64`
* `Int`
* `Float`
* `Rat`

Each Cayley octonion value is printed in the form "⦗a+bi+cj+dk+em+fn+gp+hq⦘". This is similar to `complex128` values, but uses tortoise shell brackets to distinguish.

[![Go Report Card](https://goreportcard.com/badge/gojp/goreportcard)](https://goreportcard.com/report/github.com/meirizarrygelpi/numbers/cayley) [![GoDoc](https://godoc.org/github.com/meirizarrygelpi/numbers/cayley?status.svg)](https://godoc.org/github.com/meirizarrygelpi/numbers/cayley)

## To-Do

1. Inf and NaN for `cayley.Float` and `cayley.Float64` types.
2. Fix `String` method for float types.