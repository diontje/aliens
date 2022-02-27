/*
 * world.go
 *
 */
package common

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

var (
    _, r, _, _ = runtime.Caller(0)
    Root = filepath.Join(filepath.Dir(r), "../..")    // repository root
)

const defaultGame string = "worlds/earth.dat"

// IsValidWorld() verifies the syntax or correctness of game world map
// returns err with faulty entry or ptr to common.World with data contents if valid
func IsValidWorld(world *World, params ...string) (err error) {
	var file []string

	// read the default game file "earth.dat" or the file to use
	if len(params[0]) == 0 {
		file, err = readGameFile(defaultGame)
	} else {
		file, err = readGameFile(params[0])
	}
	if err != nil {
		return err
	} else if len(file) == 0 {
		err = errors.New("cannot load game file, empty map")
		return err
	}

	// check the correctness of each line of the game file map
	cities := make(map[string]*Directions)
	for i := 0; i < len(file) - 1; i++ {
		re := regexp.MustCompile(`(^[a-zA-Z]+) (((?i)(north|south|east|west)=[a-zA-Z\-]+(\s|$))+)`)
		vals := re.FindStringSubmatch(file[i])
		if len(vals) == 0 {
			return errors.New(fmt.Sprintf("invalid map entry found on line %d, %s", i, file[i]))
		}
		city := vals[1]
		var d = new(Directions)
		// load city directions
		directions := strings.Split(vals[2], " ")
		for _, e := range directions {
			parts := strings.Split(e, "=")
			switch strings.ToLower(parts[0]) {
			case "north":
				d.North = strings.TrimSpace(parts[1])
			case "south":
				d.South = strings.TrimSpace(parts[1])
			case "east":
				d.East = strings.TrimSpace(parts[1])
			case "west":
				d.West = strings.TrimSpace(parts[1])
			}
		}
		cities[city] = d
	}
	world.Cities = cities
	return
}

// LoadGameMap() accepts the filename of thhe selected world to play
// the default is earth.dat, returns a ptr to World
func LoadGameMap(cmdArgs CmdArgs) (world *World, err error) {
	world = &World{
		Name: cmdArgs.WorldName,
		Cities: nil,
	}
	err = IsValidWorld(world, cmdArgs.WorldFileLocation)
	return
}

// readGameFile() returns game file map contents into slice
func readGameFile(filename string) (file []string, err error) {
	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return file, err
	}
	file = strings.Split(string(fileBytes), "\n")
	return file, err
}
