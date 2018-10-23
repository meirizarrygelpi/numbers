# pade

Package `pade` implements univariate [Padé approximants](https://en.wikipedia.org/wiki/Pad%C3%A9_approximant). There are five types:

* `Int64`
* `Float64`
* `Int`
* `Float`
* `Rat`

[![Go Report Card](https://goreportcard.com/badge/gojp/goreportcard)](https://goreportcard.com/report/github.com/meirizarrygelpi/numbers/pade) [![GoDoc](https://godoc.org/github.com/meirizarrygelpi/numbers/pade?status.svg)](https://godoc.org/github.com/meirizarrygelpi/numbers/pade)

## Examples

Here are some examples taken from Wikipedia.

### sin(z)

```go
package main

import (
    "fmt"
    "math/big"
    "math/cmplx"

    "github.com/meirizarrygelpi/numbers/cplex"
    "github.com/meirizarrygelpi/numbers/pade"
)

func main() {
    z := cplex.NewRat(
        big.NewRat(1, 2),
        big.NewRat(3, 4),
    )

    rSin := pade.NewRat()

    rSin.P.SetCoeff(1, big.NewRat(1, 1))
    rSin.P.SetCoeff(3, big.NewRat(-2363, 18183))
    rSin.P.SetCoeff(5, big.NewRat(12671, 4363920))

    rSin.Q.SetCoeff(0, big.NewRat(1, 1))
    rSin.Q.SetCoeff(2, big.NewRat(445, 12122))
    rSin.Q.SetCoeff(4, big.NewRat(601, 872784))
    rSin.Q.SetCoeff(6, big.NewRat(121, 16662240))

    y := new(cplex.Rat).Pade(z, rSin)

    fmt.Println("Padé approximant [5/6] of sine function:")
    fmt.Println(y.Real().FloatString(16), y.Unreal().FloatString(16))
    // Output:
    // 0.6207042311569567 0.7216508243105289

    fmt.Println("cmplx.Sin:")
    x := cmplx.Sin(complex(0.5, 0.75))
    fmt.Println(real(x), imag(x))
    // Output:
    // 0.6207042310780551 0.7216508242975646
}
```

### exp(z)

```go
package main

import (
    "fmt"
    "math/big"
    "math/cmplx"

    "github.com/meirizarrygelpi/numbers/cplex"
    "github.com/meirizarrygelpi/numbers/pade"
)

func main() {
    z := cplex.NewRat(
    big.NewRat(1, 2),
    big.NewRat(3, 4),
    )

    rExp := pade.NewRat()

    rExp.P.SetCoeff(0, big.NewRat(1, 1))
    rExp.P.SetCoeff(1, big.NewRat(1, 2))
    rExp.P.SetCoeff(2, big.NewRat(1, 9))
    rExp.P.SetCoeff(3, big.NewRat(1, 72))
    rExp.P.SetCoeff(4, big.NewRat(1, 1008))
    rExp.P.SetCoeff(5, big.NewRat(1, 30240))

    rExp.Q.SetCoeff(0, big.NewRat(1, 1))
    rExp.Q.SetCoeff(1, big.NewRat(-1, 2))
    rExp.Q.SetCoeff(2, big.NewRat(1, 9))
    rExp.Q.SetCoeff(3, big.NewRat(-1, 72))
    rExp.Q.SetCoeff(4, big.NewRat(1, 1008))
    rExp.Q.SetCoeff(5, big.NewRat(-1, 30240))

    y := new(cplex.Rat).Pade(z, rExp)

    fmt.Println("Padé approximant [5/5] of exponential function:")
    fmt.Println(y.Real().FloatString(16), y.Unreal().FloatString(16))
    // Output:
    // 1.2063510016753670 1.1238323225407718

    fmt.Println("cmplx.Exp:")
    x := cmplx.Exp(complex(0.5, 0.75))
    fmt.Println(real(x), imag(x))
    // Output:
    // 1.2063510016467855 1.1238323225841311
}
```

### Jacobi sn(z, 3)

```go
package main

import (
    "fmt"
    "math/big"

    "github.com/meirizarrygelpi/numbers/cplex"
    "github.com/meirizarrygelpi/numbers/pade"
)

func main() {
    z := cplex.NewRat(
    big.NewRat(1, 2),
    big.NewRat(3, 4),
    )

    rSN3 := pade.NewRat()

    rSN3.P.SetCoeff(1, big.NewRat(1, 1))
    rSN3.P.SetCoeff(3, big.NewRat(-572744, 4726821))
    rSN3.P.SetCoeff(5, big.NewRat(-9851629, 283609260))

    rSN3.Q.SetCoeff(0, big.NewRat(1, 1))
    rSN3.Q.SetCoeff(2, big.NewRat(859490, 1575607))
    rSN3.Q.SetCoeff(4, big.NewRat(-5922035, 56721852))
    rSN3.Q.SetCoeff(6, big.NewRat(62531591, 2977897230))

    y := new(cplex.Rat).Pade(z, rSN3)

    fmt.Println("Padé approximant [5/6] of Jacobi sn(z, 3):")
    fmt.Println(y.Real().FloatString(16), y.Unreal().FloatString(16))
    // Output:
    // 0.8637651266236634 0.4065519445560218

    // WolframAlpha Output:
    // 0.8637789783432315 0.4065354881306624
}
```

### Bessel J(5, z)

```go
package main

import (
    "fmt"
    "math/big"

    "github.com/meirizarrygelpi/numbers/cplex"
    "github.com/meirizarrygelpi/numbers/pade"
)

func main() {
    z := cplex.NewRat(
    big.NewRat(1, 2),
    big.NewRat(3, 4),
    )

    rJ5 := pade.NewRat()

    rJ5.P.SetCoeff(5, big.NewRat(1, 3840))
    rJ5.P.SetCoeff(7, big.NewRat(-107, 28416000))

    rJ5.Q.SetCoeff(0, big.NewRat(1, 1))
    rJ5.Q.SetCoeff(2, big.NewRat(151, 5550))
    rJ5.Q.SetCoeff(4, big.NewRat(1453, 3729600))
    rJ5.Q.SetCoeff(6, big.NewRat(1339, 358041600))
    rJ5.Q.SetCoeff(8, big.NewRat(2767, 120301977600))

    y := new(cplex.Rat).Pade(z, rJ5)

    fmt.Println("Padé approximant [7/8] of Bessel J(5, z):")
    fmt.Println(y.Real().FloatString(16), y.Unreal().FloatString(16))
    // Output:
    // 0.0000266221004577 -0.0001547290834750

    // WolframAlpha Output:
    // 0.0000266221 -0.000154729
}

```

### Bessel J(5, x)

```go
package main

import (
    "fmt"
    "math"
    "math/big"

    "github.com/meirizarrygelpi/numbers/cplex"
    "github.com/meirizarrygelpi/numbers/pade"
)

func main() {
    z := cplex.NewRat(
    big.NewRat(1, 2),
    big.NewRat(0, 1),
    )

    rJ5 := pade.NewRat()

    rJ5.P.SetCoeff(5, big.NewRat(1, 3840))
    rJ5.P.SetCoeff(7, big.NewRat(-107, 28416000))

    rJ5.Q.SetCoeff(0, big.NewRat(1, 1))
    rJ5.Q.SetCoeff(2, big.NewRat(151, 5550))
    rJ5.Q.SetCoeff(4, big.NewRat(1453, 3729600))
    rJ5.Q.SetCoeff(6, big.NewRat(1339, 358041600))
    rJ5.Q.SetCoeff(8, big.NewRat(2767, 120301977600))

    y := new(cplex.Rat).Pade(z, rJ5)

    fmt.Println("Padé approximant [7/8] of Bessel J(5, z):")
    fmt.Println(y.Real().FloatString(16), y.Unreal().FloatString(16))
    // Output:
    // 0.0000080536272414 0.0000000000000000

    // WolframAlpha Output:
    // 8.05363e-06

    fmt.Println("math.Jn:")
    x := math.Jn(5, 0.5)
    fmt.Println(x)
    // Output:
    // 8.053627241357474e-06
}

```