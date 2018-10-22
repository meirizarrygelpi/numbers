// Copyright (c) 2017 Melvin Eloy Irizarry-Gelpí
// Licensed under the MIT License.

/*
Package maclaurin implements univariate Maclaurin polynomials. Five
types are implemented:
    Int64   (int64 values)
    Float64 (float64 values)
    Int     (big.Int values)
    Float   (big.Float values)
    Rat     (big.Rat values)
A Maclaurin polynomial is a sum of monomials with non-negative degree.
*/
package maclaurin

const zeroDenominator = "denominator is zero"
