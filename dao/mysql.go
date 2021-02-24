package dao

import (
	"bubble/setting"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	DB *gorm.DB
)

func InitMySQL(cfg *setting.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	dbDialector := mysql.Open(dsn)
	DB, err = gorm.Open(dbDialector, &gorm.Config{
		//SkipDefaultTransaction: true,
		//PrepareStmt:            true,
		//Logger:                 gin.Logger(), //拦截、接管 gorm v2 自带日志
	})
	if err != nil {
		return
	}
	if rawDb, err := DB.DB(); err != nil {
		return err
	} else {
		rawDb.SetConnMaxIdleTime(time.Second * 30)
		rawDb.SetConnMaxLifetime(60 * time.Second)
		rawDb.SetMaxIdleConns(10)
		rawDb.SetMaxOpenConns(10)
	}
	return err
}
