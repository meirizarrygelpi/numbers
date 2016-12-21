# cockle

Package `cockle` implements arithmetic for Cockle quaternions. There are five types:

* `Int64`
* `Float64`
* `Int`
* `Float`
* `Rat`

Each Cockle quaternion value is printed in the form "⦗a+bi+ct+du⦘". This is similar to `complex128` values, but uses tortoise shell brackets to distinguish.

[![Go Report Card](https://goreportcard.com/badge/gojp/goreportcard)](https://goreportcard.com/report/github.com/meirizarrygelpi/numbers/cockle) [![GoDoc](https://godoc.org/github.com/meirizarrygelpi/numbers/cockle?status.svg)](https://godoc.org/github.com/meirizarrygelpi/numbers/cockle)

## To-Do

1. Inf and NaN for `cockle.Float` and `cockle.Float64` types.
2. Fix `String` method for float types.