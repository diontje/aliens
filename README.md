# Alien Invasion
A commandline Alien game simulation

## Installation

Building requires the installation of [Golang](https://golang.org/doc/install).

```bash
go install
```

## Game Assumptions

* Aliens do not fight if they are initialized (start) in the same city. Fighting commences only after iteration begins
* Aliens can remain in a city if their new randomly chosen city equals their current location

## Example Command Usage

```bash
aliens start
   # start the simulation using the default world, worlds/earth.dat
aliens start --debug
   # start the simulation with debug messaging
```

## Testing
Unit tests can be run by either changing into the directory with tests or by running the tests from the repository root. For example, to run all tests in the common package:

```bash
$ go test -v ./.../common
=== RUN   TestInvalidDirection
--- PASS: TestInvalidDirection (0.00s)
=== RUN   TestInvalidDirectionValue
--- PASS: TestInvalidDirectionValue (0.00s)
=== RUN   TestInvalidCityWithNoDirections
--- PASS: TestInvalidCityWithNoDirections (0.00s)
=== RUN   TestInvalidDuplicateCityEntry
--- PASS: TestInvalidDuplicateCityEntry (0.00s)
PASS
ok  	github.com/diontje/aliens/lib/common	0.832s
```
