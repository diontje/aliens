/*
 * start.go
 *
 */
package actions

import (
	"errors"
	"fmt"
	"sync"

	"github.com/diontje/aliens/lib/common"
)

const MAXITERATIONS int = 10000

// func Start() initializes and starts game
func Start(cmdArgs common.CmdArgs) error {
	log := cmdArgs.Log
	log.Debug("loading and validating game file")
	log.Debugf("calling LoadGameMap(), with args=%+v", cmdArgs)
	world, err := common.LoadGameMap(cmdArgs)
	if err != nil {
		return err
	}
	log.Debugf("calling InitAliens(), with NumAliens=%d, world=%+v",
		cmdArgs.NumAliens, world)
	err = common.InitAliens(world, cmdArgs.NumAliens)
	if err != nil {
		return err
	}
	log.Debugf("%d aliens landed into ready position", cmdArgs.NumAliens)
	for i := 0; i < cmdArgs.NumAliens; i++ {
		log.Debugf("aliens id=%d data=%+v", i, world.Aliens[i])
	}
	log.Info("game started")
	err = play(cmdArgs, world)
	if err != nil {
		return err
	}
	log.Debug("end of game alien stats")
	for _, alien := range world.Aliens {
		log.Debugf("aliens Id=%d, data=%+v", alien.Id, alien)
	}
	log.Debugf("dumping end of game map to out.dat")
	err = common.DumpGameResults(cmdArgs, world)
	if err != nil {
		return err
	}
	log.Info("game over")
	return nil
}

// deleteCity() takes a city out of game play (i.e. gets destroyed)
// aliens have fought and the city can no longer be used
// return error if unsuccessful
func deleteCity(cmdArgs common.CmdArgs, world *common.World, city string) error {
	log := cmdArgs.Log
	if len(world.Map) == 0 {
		return errors.New("game map empty, city cannot be deleted")
	}
	if _, ok := world.Map[city]; ok {
		if !ok {
			return errors.New(fmt.Sprintf("city=%s, not found in game map", city))
		}
	}
	delete(world.Map, city)
	world.Cities = common.Remove(world.Cities, city)
	if len(world.Map) == 0 { // last line in map deleted
		return nil
	}
	// now delete all directions pointing to deleted city
	for k := range world.Map {
		if world.Map[k].North == city {
			log.Debugf("clearing direction North for city=%s", k)
			world.Map[k].North = ""
		}
		if world.Map[k].South == city {
			log.Debugf("clearing direction South for city=%s", k)
			world.Map[k].South = ""
		}
		if world.Map[k].East == city {
			log.Debug("clearing direction East for city=%s", k)
			world.Map[k].East = ""
		}
		if world.Map[k].West == city {
			log.Debug("clearing direction West for city=%s", k)
			world.Map[k].West = ""
		}
	}
	return nil
}

func play(cmdArgs common.CmdArgs, world *common.World) error {
	var wg sync.WaitGroup

	log := cmdArgs.Log
	log.Debugf("playing with N=%d aliens", len(world.Aliens))
	for i := 1; i <= MAXITERATIONS; i++ {
		for j := 0; j < len(world.Aliens); j++ {
			log.Debugf("manipulating alien id=%d", world.Aliens[j].Id)
			if world.Aliens[j].Trapped {
				continue
			}
			wg.Add(1)
			go common.Move(&wg, world.Aliens[j].Id, world, cmdArgs.Log)
		}
		wg.Wait()
		log.Debugf("all aliens have moved, checking for fighting")

		// loop through each city
		k := 0
		for len(world.Map) != 0 && len(world.Cities) != 0 && k < len(world.Cities) {
			city := world.Cities[k]
			fighting := updateCity(cmdArgs, world, city)
			if fighting {
				// destroy the city, update the game map
				err := deleteCity(cmdArgs, world, city)
				if err != nil {
					return err
				}
				log.Infof("city=%s, destroyed", city)
				k = 0
			} else {
				log.Infof("no fighting in city=%s", city)
				k++
			}
		}
		if len(world.Aliens) == 1 {
			log.Debug("there can only be 1, we have a winner")
			break
		}
		// check for empty map and trapped aliens two conditions
		if (world.NumAliensTrapped == cmdArgs.NumAliens - world.NumAliensKilled) ||
				(len(world.Map) == 0 && len(world.Aliens) != 0) {
			log.Info("all aliens are Trapped")
			break
		}
	}
	return nil
}

// updateCity() updates the "game board" for city
// determines which aliens are fighting in that city and sets the city for destruction
// if there is fighting, kill off aliens and return true to blow up city
func updateCity(cmdArgs common.CmdArgs, world *common.World, city string) (fighting bool) {
	var ids []int

	log := cmdArgs.Log
	fighting = false
	log.Debugf("updating city=%s", city)
	for _, alien := range world.Aliens {
		log.Debugf("manipulating alien Id=%d", alien.Id)
		if alien.Location == city {
			ids = append(ids, alien.Id)
		}
	}
	if len(ids) > 1 {
		log.Infof("aliens Ids=%v, are fighting in city=%s", ids, city)
		log.Info("killing off aliens and setting city for destruction")
		world.NumAliensKilled = len(ids)
		for j := 0; j < len(ids); j++ {
			delete(world.Aliens, ids[j])
		}
		fighting = true
	} else {
		log.Debugf("fighting=%t, city not destroyed", fighting)
	}
	return
}
