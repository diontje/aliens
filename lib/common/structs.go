/*
 * structs.go
 *
 */
package common

import (
	"github.com/sirupsen/logrus"
)

// struct to store command line arguments
type CmdArgs struct {
	Debug             bool
	Log               *logrus.Logger
	NumAliens         int
	WorldName         string
	WorldFileLocation string // TODO: passing world from non default location
}
