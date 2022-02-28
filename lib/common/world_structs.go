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
	Name   string
	Cities []string               // list of world cities
	Map    map[string]*Directions // direction => city
	Aliens map[int]*Alien         // Id => alien
}
