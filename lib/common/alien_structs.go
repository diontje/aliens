/*
 * world_structs.go
 *
 */
package common

const MAX_NUM_ALIENS int = 100

// struct type defining an Alien
type Alien struct {
	Id       int
	Fighting bool
	Location string
	NumMoves int
	Stuck    bool
}
