/*
 * common_world_data_structs.go
 *
 */
package common

type Direction int

const (
	North Direction = iota
	South
	East
	West
)

// defines a game's world
type World struct {
	Name   string							// name of the world == filename
	Cities map[string]map[Direction] string	// map of directions to other cities 
}
