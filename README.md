# numbers

[![Go Report Card](https://goreportcard.com/badge/gojp/goreportcard)](https://goreportcard.com/report/github.com/meirizarrygelpi/numbers) [![GoDoc](https://godoc.org/github.com/meirizarrygelpi/numbers?status.svg)](https://godoc.org/github.com/meirizarrygelpi/numbers)

Metapackage `numbers` is a collection of packages that implement arithmetic over many number systems, including dual numbers, quaternions, octonions, and their parabolic and hyperbolic cousins. In each package five types are implemented:

* `Int64`
* `Float64`
* `Int`
* `Float`
* `Rat`

Each value is printed in the form "⦗...⦘". This is similar to `complex128` values, but uses tortoise shell brackets to distinguish.

Here is a list of available packages:

1. `vec3`: three-dimensional vectors
1. `vec7`: seven-dimensional vectors
1. `eisenstein`: [Eisenstein numbers](https://en.wikipedia.org/wiki/Eisenstein_integer)
1. `heegner`: imaginary quadratic fields with class number 1. See [Heegner numbers](https://en.wikipedia.org/wiki/Heegner_number)
1. `maclaurin`: [Maclaurin polynomials](https://en.wikipedia.org/wiki/Polynomial)
1. `pade`: [Padé approximants](https://en.wikipedia.org/wiki/Pad%C3%A9_approximant)
1. `cplex`: [complex numbers](https://en.wikipedia.org/wiki/Complex_number)
1. `nplex`: nilplex numbers (more commonly known as [dual numbers](https://en.wikipedia.org/wiki/Dual_number))
1. `pplex`: perplex numbers (more commonly known as [split-complex numbers](https://en.wikipedia.org/wiki/Split-complex_number))
1. `hamilton`: Hamilton quaternions (i.e. traditional [quaternions](https://en.wikipedia.org/wiki/Quaternion); can also be referred to as elliptic quaternions; four-dimensional)
1. `cockle`: Cockle quaternions (more commonly known as [split-quaternions](https://en.wikipedia.org/wiki/Split-quaternion); can also be referred to as hyperbolic quaternions; four-dimensional)
1. `supernplex`: super-nilplex numbers (different from bi-nilplex numbers; can also be referred to as parabolic quaternions; four-dimensional)
1. `supercplex`: super-complex numbers (different from dual-complex numbers; four-dimensional)
1. `superpplex`: super-perplex numbers (different from dual-perplex numbers; four-dimensional)
1. `bicplex`: [bi-complex numbers](https://en.wikipedia.org/wiki/Bicomplex_number) (complexification of the complex numbers; four-dimensional)
1. `bipplex`: bi-perplex numbers (perplexification of the perplex numbers; four-dimensional)
1. `binplex`: bi-nilplex numbers (nilplexification of the nilplex numbers; four-dimensional)
1. `dualcplex`: dual-complex numbers (nilplexification of the complex numbers; four-dimensional)
1. `dualpplex`: dual-perplex numbers (nilplexification of the perplex numbers; four-dimensional)
1. `cayley`: Cayley octonions (i.e. traditional [octonions](https://en.wikipedia.org/wiki/Octonion); can also be referred to as elliptic octonions; eight-dimensional)
1. `zorn`: Zorn octonions (more commonly known as [split-octonions](https://en.wikipedia.org/wiki/Split-octonion); can also be referred to as hyperbolic octonions; eight-dimensional)
1. `ultranplex`: ultra-nilplex numbers (different from tri-nilplex numbers; can also be referred to as parabolic octonions; eight-dimensional)
1. `superhamilton`: super-Hamilton quaternions (different from the dual-Hamilton quaternions; eight-dimensional)
1. `supercockle`: super-Cockle quaternions (different from the dual-Cockle quaternions; eight-dimensional)
1. `ultracplex`: ultra-complex numbers (different from the hyper-complex numbers; eight-dimensional)
1. `ultrapplex`: ultra-perplex numbers (different from the hyper-perplex numbers; eight-dimensional)
1. `tricplex`: tri-complex numbers (complexification of the bi-complex numbers; eight-dimensional)
1. `trinplex`: tri-nilplex numbers (nilplexification of the bi-nilplex numbers; eight-dimensional)
1. `tripplex`: tri-perplex numbers (perplexification of the di-perplex numbers; eight-dimensional)
1. `hypercplex`: hyper-complex numbers (nilplexification of dual-complex numbers; eight-dimensional)
1. `hyperpplex`: hyper-perplex numbers (nilplexification of dual-perplex numbers; eight-dimensional)
1. `dualhamilton`: dual-Hamilton quaternions (nilplexification of Hamilton quaternions; eight-dimensional)
1. `dualcockle`: dual-Cockle quaternions (nilplexification of Cockle quaternions; eight-dimensional)
1. `comhamilton`: complex-Hamilton quaternions (complexification of Hamilton quaternions; eight-dimensional)
1. `perhamilton`: perplex-Hamilton quaternions (perplexification of Hamilton quaternions; eight-dimensional)
1. `percockle`: perplex-Cockle quaternions (perplexification of Cockle quaternions; eight-dimensional)

Here is a list of future packages:

1. `laurent`: [Laurent polynomials](https://en.wikipedia.org/wiki/Laurent_polynomial)

To-Do:

1. `SetReal` and `SetUnreal` methods
1. `Plus` and `Minus` methods
1. `Maclaurin` methods
1. `Padé` methods
1. `Inf` and `NaN` methods
1. `IsInf` and `IsNaN` methods
1. Rename `Star1` methods `Bar`
1. Rename `Star2` methods `Tilde`
1. Rename `Star3` methods `Star`
1. `Dot` and `Cross` methods
1. Rename `supracplex` package `ultracplex`
1. Rename `suprapplex` package `ultrapplex`
1. Rename `infracockle` package `supercockle`
1. Rename `infracplex` package `supercplex`
1. Rename `infrapplex` package `superpplex`