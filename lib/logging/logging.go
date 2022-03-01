/*
 * logging.go
 *
 *
 */
package logging

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger
var err error

// GetLogger() accepts logger initializer variables and returns log handle
func GetLogger(debug bool) (*logrus.Logger, error) {
	if log != nil {
		return log, nil
	}
	initLogger(debug)
	return log, err
}

// func initLogger() initialized log handle {
func initLogger(debug bool) {
	log = logrus.New()
	log.SetFormatter(&nested.Formatter{
		HideKeys:    true,
		FieldsOrder: []string{"component", "category"},
	})
	if debug {
		log.SetLevel(logrus.DebugLevel)
	}
}
