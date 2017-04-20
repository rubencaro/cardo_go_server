/*
Package helpers contains generic helper functions
*/
package hlp

// This file contains helpers to work with config files

import (
	"github.com/BurntSushi/toml"
)

// Config and all subtypes, wrap all config options
type Config struct {
	DB database
}

type database struct {
	URL  string
	User string
	Pass string
}

// Conf holds all config options after call to 'init'
var Conf Config

// readConfig reads the contents of the 'config.toml' file
// into the 'Conf' struct
func readConfig() {
	// return already parsed config if it's present
	if Conf != (Config{}) {
		return
	}

	_, err := toml.DecodeFile("config.toml", &Conf)
	if err != nil {
		panic(err)
	}
}

func init() {
	readConfig()
}
