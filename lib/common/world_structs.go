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
	Aliens            map[int]*Alien         // Id => alien
	Cities            []string               // list of world cities
	Map               map[string]*Directions // direction => city
	Name              string
	NumAliensKilled   int
	NumAliensTrapped  int
}
