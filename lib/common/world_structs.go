/*
 * world_structs.go
 *
 */
package common

// struct type to contain mapping of direction to city
type Directions struct {
	North string
	South string
	East  string
	West  string
}

// struct type defining a game's world or map dat file
type World struct {
	Name  string
	Cities map[string]*Directions
}
