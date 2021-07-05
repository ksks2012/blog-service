package config

import (
	"errors"

	bsinterfaces "github.com/blog-service/interfaces"
	"github.com/blog-service/pkg/setting"
)

// StorageSetup contains storage type and storage instance
type StorageSetup struct {
	Type     string
	Instance bsinterfaces.StorageEngine
}

func (s *StorageSetup) NewDBEngine(databaseSetting *setting.DatabaseSettingS) (err error) {
	switch databaseSetting.DBType {
	case "pxc":
		s.Instance, err = setupMySQLRoundRobinStorageEngine(databaseSetting)
	case "mysql", "mariadb":
		s.Instance, err = setupMySQLStorageEngine(databaseSetting)
	default:
		err = errors.New("unknown storage engine type: " + databaseSetting.DBType)
	}
	if nil != err {
		s.Instance = nil
	} else {
		s.Type = databaseSetting.DBType
	}
	return err
}
