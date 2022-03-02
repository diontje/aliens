/*
 * alien.go
 *
 */
package common

import (
	"math/rand"
	"sync"

	"github.com/sirupsen/logrus"
)

// InitAliens() initializes world.Aliens i.e., assigns N aliens to their
// randomly assigned city, assigns an Id, a Trapped status (false), and a counter
// for the number of moves taken.
func InitAliens(world *World, numAliens int) (err error) {
	world.Aliens = make(map[int]*Alien)
	for i := 0; i < numAliens; i++ {
		n := rand.Int() % len(world.Cities)
		temp := &Alien{
			Id: i,
			Fighting: false,
			Location: world.Cities[n],
			NumMoves: 0,
			Trapped: false,
		}
		world.Aliens[i] = temp
	}
	return nil
}

// Move() moves randomly assigns alien to a new city
func Move(wg *sync.WaitGroup, Id int, world *World, log *logrus.Logger) error {
	defer wg.Done()

	log.Infof("Id=%d, currentCity=%s", Id, world.Aliens[Id].Location)

	// choose a random city, by selecting a random direction to follow
	// if the city does not gives directions, Trapped=true
	if contains(world.Cities, world.Aliens[Id].Location) == false {
		world.Aliens[Id].Trapped = true
		log.Infof("Id=%d, Trapped=%t", Id, world.Aliens[Id].Trapped)
		world.NumAliensTrapped++
		return nil
	}
	var newCity string
	for newCity == "" {
		switch rand.Int() % 4 {
		case 0:
			newCity = world.Map[world.Aliens[Id].Location].North
		case 1:
			newCity = world.Map[world.Aliens[Id].Location].South
		case 2:
			newCity = world.Map[world.Aliens[Id].Location].East
		case 3:
			newCity = world.Map[world.Aliens[Id].Location].West
		}
	}
	log.Infof("Id=%d, newCity=%s", Id, newCity)
	world.Aliens[Id].Location = newCity
	log.Infof("alien Id=%d moved", Id)
	return nil
}
