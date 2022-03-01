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
	log.Debug("calling LoadGameMap(), with args=%+v", cmdArgs)
	world, err := common.LoadGameMap(cmdArgs)
	if err != nil {
		return err
	}
	log.Debug("calling InitAliens(), with NumAliens=%d, world=%+v",
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
	log.Debug("alien stats at end of game")
	for i := 0; i < cmdArgs.NumAliens; i++ {
		log.Debugf("aliens id=%d data=%+v", i, world.Aliens[i])
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
	// now delete directions pointing to city
	for k, _ := range world.Map {
		if world.Map[k].North == city {
			log.Debugf("clearing direction North for city=%s", city)
			world.Map[k].East = ""
		}
		if world.Map[k].South == city {
			log.Debugf("clearing direction South for city=%s", city)
			world.Map[k].East = ""
		}
		if world.Map[k].East == city {
			log.Debug("clearing direction East for city=%s", city)
			world.Map[k].East = ""
		}
		if world.Map[k].West == city {
			log.Debug("clearing direction West for city=%s", city)
			world.Map[k].East = ""
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
			log.Debugf("working alien id=%d", world.Aliens[j].Id)
			if world.Aliens[j].Trapped {
				continue
			}
			wg.Add(1)
			go common.Move(&wg, world.Aliens[j].Id, world, cmdArgs.Log)
		}
		wg.Wait()
		log.Debugf("all aliens have moved, checking for fighting")
		for _, city := range world.Cities {
			fighting := updateCity(cmdArgs, world, city)
			if fighting {
				// destroy the city, update the game map
				err := deleteCity(cmdArgs, world, city)
				if err != nil {
					return err
				}
				log.Infof("city=%s, destroyed", city)
			}
		}
		//log.Debugf("waiting for interation to end")
		if world.NumAliensKilled == cmdArgs.NumAliens {
			log.Debug("there can only be 1, we have a winner")
			break
			//TODO: write func to get last man standing who is not stuck
			// assumption, stuck aliens lost
		} else if world.NumAliensTrapped == cmdArgs.NumAliens - world.NumAliensKilled {
			log.Debug("all aliens are stuck")
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
	log.Debugf("checking city=%s for fighting", city)
	for i := 0; i < len(world.Aliens); i++ {
		log.Debugf("working alien Id=%d", world.Aliens[i].Id)
		if world.Aliens[i].Location == city {
			ids = append(ids, world.Aliens[i].Id)
		}
	}
	if len(ids) > 1 {
		log.Infof("aliens=%v, are fighting in city=%s", ids, city)
		log.Debug("killing off aliens and setting city for destruction")
		world.NumAliensKilled = len(ids)
		for j := 0; j < len(ids); j++ {
			delete(world.Aliens, ids[j])
		}
		fighting = true
	}
	return
}
