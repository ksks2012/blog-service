package main

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v3"

	storageconfig "github.com/blog-service/internal/dao/config"
)

const (
	minimumRedoQueueSize    = 2
	minimumStorageQueueSize = 8
	defaultR8AServer        = "dns:nsevm.miccwb.midev"
)

type configuration struct {
	StorageSetup *storageconfig.StorageSetup `yaml:"storage"`
}

func loadConfiguration(filename string) (cfg *configuration, err error) {
	buf, err := ioutil.ReadFile(filename)
	if nil != err {
		return
	}
	cfg = &configuration{}
	if err = yaml.Unmarshal(buf, cfg); nil != err {
		cfg = nil
		return
	}

	return cfg, nil
}
