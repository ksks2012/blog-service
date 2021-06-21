package main

import (
	"errors"
	"flag"
)

// ErrRequireConfigurationFile indicates configuration file is missing.
var ErrRequireConfigurationFile = errors.New("configuration file is required")

func parseCommandParam() (cfg *configuration, err error) {
	var configFilePath string
	flag.StringVar(&configFilePath, "conf", "", "path to collect configuration")
	flag.Parse()
	if "" == configFilePath {
		return nil, ErrRequireConfigurationFile
	}
	return loadConfiguration(configFilePath)
}
