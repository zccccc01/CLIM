package svc

import (
	"CLIM/app/user/cmd/rpc/internal/config"
	"context"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	dsn := c.MySQL.DataSource
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		db.Logger.Error(context.Background(), "mysql connect error", err)
		return nil
	}
	return &ServiceContext{
		Config: c,
		DB:     db.Debug(),
	}
}
