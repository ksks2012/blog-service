package config

import (
	"errors"

	yaml "gopkg.in/yaml.v3"

	bsinterfaces "github.com/blog-service/interfaces"
)

// StorageSetup contains storage type and storage instance
type StorageSetup struct {
	Type     string
	Instance bsinterfaces.StorageEngine
}

// UnmarshalYAML implements yaml.Unmarshaler interface
func (s *StorageSetup) UnmarshalYAML(nodeValue *yaml.Node) (err error) {
	var aux struct {
		StorageType string `yaml:"type"`
	}
	if err = nodeValue.Decode(&aux); nil != err {
		return
	}
	switch aux.StorageType {
	case "pxc":
		s.Instance, err = setupMySQLRoundRobinStorageEngine(nodeValue)
	case "mysql", "mariadb":
		s.Instance, err = setupMySQLStorageEngine(nodeValue)
	default:
		err = errors.New("unknown storage engine type: " + aux.StorageType)
	}
	if nil != err {
		s.Instance = nil
	} else {
		s.Type = aux.StorageType
	}
	return err
}
