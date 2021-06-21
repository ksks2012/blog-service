package config

import (
	yaml "gopkg.in/yaml.v3"

	bsinterfaces "github.com/blog-service/interfaces"
	mysqlstorage "github.com/blog-service/internal/dao/mysql"
)

func setupMySQLStorageEngine(nodeValue *yaml.Node) (storage bsinterfaces.StorageEngine, err error) {
	var aux struct {
		UserName       string `yaml:"username"`
		Password       string `yaml:"password"`
		NetworkAddress string `yaml:"address"`
		SocketPath     string `yaml:"socket"`
		DatabaseName   string `yaml:"database"`
	}
	if err = nodeValue.Decode(&aux); nil != err {
		return
	}
	return mysqlstorage.NewMySQLStorageEngine(aux.UserName, aux.Password, aux.NetworkAddress, aux.SocketPath, aux.DatabaseName)
}

func setupMySQLRoundRobinStorageEngine(nodeValue *yaml.Node) (storage bsinterfaces.StorageEngine, err error) {
	var aux struct {
		UserName         string   `yaml:"username"`
		Password         string   `yaml:"password"`
		NetworkAddresses []string `yaml:"addresses"`
		SocketPath       string   `yaml:"socket"`
		DatabaseName     string   `yaml:"database"`
	}
	if err = nodeValue.Decode(&aux); nil != err {
		return
	}
	return mysqlstorage.NewMySQLRoundRobinStorageEngine(aux.UserName, aux.Password, aux.NetworkAddresses, aux.SocketPath, aux.DatabaseName)
}
