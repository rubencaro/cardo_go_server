/*
Package helpers contains generic helper functions
*/
package helpers

// This file contains helpers to work with config files

import (
	"github.com/BurntSushi/toml"
)

// configStruct and all subtypes, wrap all config options
type configStruct struct {
	DB database
}

type database struct {
	URL  string
	User string
	Pass string
}

// Config holds all config options after call to 'init'
var Config configStruct

// readConfig reads the contents of the 'config.toml' file
// into the 'Config' struct
func readConfig() {
	// return already parsed config if it's present
	if Config != (configStruct{}) {
		return
	}

	_, err := toml.DecodeFile("config.toml", &Config)
	if err != nil {
		panic(err)
	}
}

func init() {
	readConfig()
}
