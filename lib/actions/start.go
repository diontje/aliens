/*
 * start.go
 *
 */
package actions

import (
	"fmt"

	"github.com/diontje/aliens/lib/common"
)

// func Start() initializes game play
func Start(cmdArgs common.CmdArgs) error {

	// initializeGame() loads the game map and performing error checking, as well as
	// processing any commandline args such as logging
	log := cmdArgs.Log
	log.Debug("loading and validating game file")
	world, err := common.LoadGameMap(cmdArgs)
	if err != nil {
		return err
	}
	log.Debugf("adding %d aliens in ready position", cmdArgs.NumAliens)
	err = common.InitAliens(world, cmdArgs.NumAliens)
	if err != nil {
		return err
	}
	fmt.Printf("world[Bar] cities=%+v\n", world.Map["Bar"])
	fmt.Printf("world[Foo] cities=%+v\n", world.Map["Foo"])
	fmt.Printf("aliens id=%d data=%+v\n", 0, world.Aliens[0])
	fmt.Printf("aliens id=%d data=%+v\n", 1, world.Aliens[1])
	fmt.Printf("world cities=%v\n", world.Cities)

	return nil
}
