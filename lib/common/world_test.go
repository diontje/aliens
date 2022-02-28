/*
 * world_test.go
 *
 */
package common_test

import (
	"testing"

	"github.com/diontje/aliens/lib/common"
)

var world1 string = common.Root + "/test/data/world1.dat"
var world2 string = common.Root + "/test/data/world2.dat"
var world3 string = common.Root + "/test/data/world3.dat"
var world4 string = common.Root + "/test/data/world4.dat"

// Test game map with invalid direction, should be valid
// Invalid directions will be ignored
func TestInvalidDirection(t *testing.T) {
	world := &common.World{Name: "", Map: nil}
	err := common.IsValidWorld(world, world1)
	if err != nil {
		t.Error(err.Error())
	}
}

// Test game map with empty direction value, should be valid
// City was not assigned to direction
func TestInvalidDirectionValue(t *testing.T) {
	world := &common.World{Name: "", Map: nil}
	err := common.IsValidWorld(world, world2)
	if err != nil {
		t.Error(err.Error())
	}
}

// Test game map with empty directions for city, should be invalid
// At least one city direction should be found
func TestInvalidCityWithNoDirections(t *testing.T) {
	world := &common.World{Name: "", Map: nil}
	err := common.IsValidWorld(world, world3)
	if err == nil {
		t.Error("failed to detect invalid city with no directions")
	}
}

// Test game map with duplicate cities, should be invalid
// Each map entry for a city must be unique
func TestInvalidDuplicateCityEntry(t *testing.T) {
	world := &common.World{Name: "", Map: nil}
	err := common.IsValidWorld(world, world4)
	if err == nil {
		t.Errorf("failed to detect duplicate city entry")
	}
}
