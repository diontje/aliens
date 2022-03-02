/*
 * common.go
 *
 */
package common

import (
	"io/ioutil"
	"strings"

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

// returns slice with element r removed
func Remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
// returns true if slice contains a string, false otherwise
func contains(elems []string, v string) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

// readGameFile() returns game file map contents into slice
func readGameFile(filename string) (file []string, err error) {
	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return file, err
	}
	file = strings.Split(string(fileBytes), "\n")
	return file, err
}
