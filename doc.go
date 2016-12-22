// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

/*
Package numbers is a collection of packages that implement arithmetic
for many number systems. Each individual package implements five numeric types:
    Int64   (int64 components)
    Float64 (float64 components)
    Int     (big.Int components)
    Float   (big.Float components)
    Rat     (big.Rat components)
The type of the components in these five types can be understood as a number
system in 2⁰ = 1 dimension. There are packages that implement number systems in
2¹ = 2, 2² = 4, and 2³ = 8 dimensions.

The two-dimensional packages are:
    cplex (complex numbers; complexification of the reals)
    nplex (nilplex numbers; dual numbers; nilplexification of the reals)
    pplex (perplex numpers; split-complex numbers; perplexification of the reals)
The four-dimensional packages are:
    hamilton   (quaternions; elliptic quaternions)
    cockle     (Cockle quaternions; split-quaternions; hyperbolic quaternions)
    supra      (supra numbers; parabolic quaternions)
    infracplex (infra-complex numbers)
    infrapplex (infra-perplex numbers)
    bicplex    (bi-complex numbers; tessarines; complexification of the complex numbers)
    bipplex    (bi-perplex numbers; perplexification of the perplex numbers)
    hyper      (hyper-numbers; nilplexification of the nilplex numbers)
    dualcplex  (dual-complex numbers; nilplexification of the complex numbers)
    dualpplex  (dual-perplex numbers; nilplexification of the perplex numbers)
The eight-dimensional packages are:
    cayley        (octonions; elliptic octonions)
    zorn          (Zorn octonions; split-octonions; hyperbolic octonions)
    ultra         (ultra numbers; parabolic octonions)
    infrahamilton (infra-Hamilton quaternions)
    infracockle   (infra-Cockle quaternions)
    supracplex    (supra-complex numbers)
    suprapplex    (supra-perplex numbers)
    tricplex      (tri-complex numbers; complexification of the bi-complex numbers)
    tripplex      (tri-perplex numbers; perplexification of the bi-perplex numbers)
    trinplex      (tri-nilplex numbers; nilplexification of the hyper-numbers)
    dualhamilton  (dual-Hamilton quaternions; nilplexification of the quaternions)
    dualcockle    (dual-Cockle quaternions; nilplexification of the Cockle quaternions)
    hypercplex    (hyper-complex numbers; nilplexification of the dual-complex numbers)
    hyperpplex    (hyper-perplex numbers; nilplexification of the dual-perplex numbers)
All three two-dimensional systems include a binary multiplication operation that
is commutative and associative (for non-floats). Both the nilplex and perplex numbers include
non-trivial zero divisors.

Five of the ten four-dimensional systems include a binary multiplication operation
that is non-commutative but associative (for non-floats). These are the Hamilton quaternions, Cockle
quaternions, supra numbers, infra-complex numbers, and infra-perplex numbers.

The other five four-dimensional numbers systems are plexifications of the three
two-dimensional number systems. Here the binary multiplication operation is
commutative and associative (for non-floats).

Except for the Hamilton quaternions, all other four-dimensional number systems
include non-trivial zero divisors.

Seven of the fourteen eight-dimensional systems include a binary multiplication
operation that is non-commutative and non-associative. These are the Cayley octonions,
Zorn octonions, ultra numbers, infra-Hamilton quaternions, infra-Cockle quaternions,
supra-complex numbers, and supra-perplex numbers.

The other seven eight-dimensional number systems are plexifications of
four-dimensional number systems. For dual-Hamilton quaternions, and dual-Cockle
quaternions, the multiplication operation is non-commutative but associative (for
non-floats). For tri-complex numbers, tri-perplex numbers, tri-nilplex numbers,
hyper-complex numbers, and hyper-perplex numbers the multiplication operation is
commutative and associative (for non-floats).
*/
package numbers
