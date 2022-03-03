# Alien Invasion
Alien Invasion is a command line Alien game simulation. A number of aliens (--numAliens, default == 2) start from random cities determined by supplied data file (default == worlds/earth.dat). Aliens move for maximum 10000 times until either all aliens are Trapped or they cannot move to another city. During an iteration if more than one Alien lands in the same city, they fight and the city is removed from the world map.

## Game Assumptions

* Aliens do not fight if they are initialized (start) in the same city. Fighting commences only after iteration begins
* Aliens can remain in a city if their new randomly chosen city equals their current location


## Installation

Building requires the installation of [Golang](https://golang.org/doc/install).

```bash
go install
```

## Example Command Usage

```bash
aliens start
   # start the simulation using N=2 aliens and the default world, worlds/earth.dat
aliens start --numAliens 100
   # start the simulation with N=100 aliens with the default world worlds/earth.dat
aliens start --debug
   # start the simulation with debug logging 
```

## Sample Data file format
```
Foo north=Bar west=Baz south=Qu-ux
Bar south=Foo west=Bee
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
