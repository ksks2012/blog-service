package dao

import (
	"github.com/blog-service/internal/model"
)

func (d *Dao) GetAuth(appKey string, appSecret string) (model.Auth, error) {
	auth := model.Auth{AppKey: appKey, AppSecret: appSecret}
	return auth.Get(d.engine)
}
