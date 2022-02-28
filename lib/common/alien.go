/*
 * alien.go
 *
 */
package common

import (
	"math/rand"
)

// InitAliens() initializes world.Aliens i.e., places each N alien in their
// randomly assigned starting position, assigns an Id, Fighting bool to false,
// and current number of moves to 0
func InitAliens(world *World, numAliens int) (err error) {
	world.Aliens = make(map[int]*Alien)
	for i := 0; i < numAliens; i++ {
		n := rand.Int() % len(world.Cities)
		temp := &Alien{
			Id: i,
			Fighting: false,
			Location: world.Cities[n],
			NumMoves: 0,
		}
		world.Aliens[i] = temp
	}
	return nil
}
