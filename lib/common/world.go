/*
 * world.go
 *
 */
package common

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
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

// DumpGameResults() writes the end of game map to file
func DumpGameResults(cmdArgs CmdArgs, world *World) error {
    f, err := os.Create("out.dat")
    if err != nil {
		return err
    }
    defer f.Close()
	for k, v := range world.Map {
		var directions []string
		if v.North != "" {
			directions = append(directions, fmt.Sprintf("North=%s", v.North))
		}
		if v.South != "" {
			directions = append(directions, fmt.Sprintf("South=%s", v.South))
		}
		if v.East != "" {
			directions = append(directions, fmt.Sprintf("East=%s", v.East))
		}
		if v.West != "" {
			directions = append(directions, fmt.Sprintf("West=%s", v.West))
		}
		s := strings.TrimSpace(strings.Join(directions, " "))
		_, err = f.WriteString(fmt.Sprintf("%s %s\n", k, s))
		if err != nil {
			return err
		}
	}
	return nil
}

// IsValidWorld() verifies the syntax or correctness of the game world map
// returns err with faulty entry or ptr to common.World with data contents if valid
func IsValidWorld(world *World, params ...string) (err error) {
	var file []string

	// read the default game file "earth.dat" or the file to use
	if len(params[0]) == 0 {
		file, err = readGameFile(defaultGame)
	} else {
		//TODO: add support
		file, err = readGameFile(params[0])
	}
	if err != nil {
		return err
	} else if len(file) == 0 {
		err = errors.New("cannot load game file, empty map")
		return err
	}

	// check the correctness of each line of the game file map
	m := make(map[string]*Directions)
	for i := 0; i < len(file) - 1; i++ {
		re := regexp.MustCompile(`(^[a-zA-Z]+) (((?i)(north|south|east|west)=[a-zA-Z\-]+(\s|$))+)`)
		vals := re.FindStringSubmatch(file[i])
		if len(vals) == 0 {
			return errors.New(fmt.Sprintf("invalid map entry found on line %d, %s", i, file[i]))
		}
		city := vals[1]
		// check if city already exists, if so discard, duplicate cities are not allowed
		if len(m) != 0 {
			if _, ok := m[city]; ok {
				return errors.New(fmt.Sprintf("duplicate city data detected, city=%s\n", city))
			}
		}
		world.Cities = append(world.Cities, city)
		// load each city's map details
		var d = new(Directions)
		directions := strings.Split(vals[2], " ")
		for _, e := range directions {
			parts := strings.Split(e, "=")
			// TODO: check regexp, ensure that we're not picking up empty parts (spaces)
			if len(parts) != 2 {
				continue
			}
			direction := strings.ToLower(parts[0])
			toCity := strings.TrimSpace(parts[1])
			// ignore directions with no city value
			if toCity == "" {
				continue
			}
			switch direction {
			case "north":
				d.North = toCity
			case "south":
				d.South = toCity
			case "east":
				d.East = toCity
			case "west":
				d.West = toCity
			}
		}
		m[city] = d
	}
	world.Map = m
	return
}

// LoadGameMap() accepts the filename of the selected world to play
// (the default is earth.dat), and returns a ptr to the intialized game World
func LoadGameMap(cmdArgs CmdArgs) (world *World, err error) {
	world = &World{
		Map: nil,
		Name: cmdArgs.WorldName,
		NumAliensKilled: 0,
		NumAliensTrapped: 0,
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


