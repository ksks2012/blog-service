package config

import (
	bsinterfaces "github.com/blog-service/interfaces"
	mysqlstorage "github.com/blog-service/internal/dao/dbversion/mysql"
	"github.com/blog-service/pkg/setting"
)

func setupMySQLStorageEngine(databaseSetting *setting.DatabaseSettingS) (storage bsinterfaces.StorageEngine, err error) {
	if len(databaseSetting.Host) == 0 {
		return
	}
	return mysqlstorage.NewMySQLStorageEngine(databaseSetting.UserName, databaseSetting.Password, databaseSetting.Host[0], databaseSetting.SocketPath, databaseSetting.DBName)
}

func setupMySQLRoundRobinStorageEngine(databaseSetting *setting.DatabaseSettingS) (storage bsinterfaces.StorageEngine, err error) {
	return mysqlstorage.NewMySQLRoundRobinStorageEngine(databaseSetting.UserName, databaseSetting.Password, databaseSetting.Host, databaseSetting.SocketPath, databaseSetting.DBName)
}
