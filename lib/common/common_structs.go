/*
 * common_structs.go
 *
 */
package common

// struct to store command line arguments
type CmdArgs struct {
	WorldName         string
	WorldFileLocation string // TODO: passing world from non default location
}
