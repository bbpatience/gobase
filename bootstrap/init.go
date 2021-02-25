package bootstrap

import (
	"fmt"
	"go.uber.org/zap/zapcore"
	"gobase/global/variables"
	"gobase/global/my_errors"
	"gobase/utils/gorm_v2"
	zap_factory "gobase/utils/logger"
	"gobase/utils/yml_config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

func checkRequiredConfigs() {
	//check configs
	if _, err := os.Stat("./config/config.yml"); err != nil {
		log.Fatal(my_errors.ErrorsConfigYamlNotExists + err.Error())
	}
	if _, err := os.Stat("./config/gorm_v2.yml"); err != nil {
		log.Fatal(my_errors.ErrorsConfigGormNotExists + err.Error())
	}
}

func init() {
	//1.check deploy folder , config files.
	checkRequiredConfigs()

	// 2.config parse, and add listener
	variables.ConfigYml = yml_config.CreateYamlFactory()
	variables.ConfigYml.ConfigFileChangeListen()
	// 3. db config parse, add listener.
	variables.ConfigGormv2Yml = variables.ConfigYml.Clone("gorm_v2")
	variables.ConfigGormv2Yml.ConfigFileChangeListen()

	// 3.init log.  Zap.
	variables.ZapLog = zap_factory.CreateZapFactory(func(entry zapcore.Entry) error {
		go func(paramEntry zapcore.Entry) {
		}(entry)
		return nil
	})

	// 4.init db.
	if dbMysql, err := GetMysqlDriver(); err != nil {
		log.Fatal(my_errors.ErrorsGormInitFail + err.Error())
	} else {
		variables.GormDbMysql = dbMysql
	}

}

func GetMysqlDriver(dbConf ...gorm_v2.ConfigParams) (*gorm.DB, error) {
	dsn := getDsn(dbConf...)
	gormDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//SkipDefaultTransaction: true,
		//PrepareStmt:            true,
		//Logger:                 redefineLog(sqlType),
	})
	if err != nil {
		return nil, err
	}

	// 查询没有数据，屏蔽 gorm v2 包中会爆出的错误
	// https://github.com/go-gorm/gorm/issues/3789  此 issue 所反映的问题就是我们本次解决掉的
	_ = gormDb.Callback().Query().Before("gorm:query").Register("disable_raise_record_not_found", func(d *gorm.DB) {
		d.Statement.RaiseErrorOnNotFound = false
	})

	// config db connection pool.
	if rawDb, err := gormDb.DB(); err != nil {
		return nil, err
	} else {
		rawDb.SetConnMaxIdleTime(time.Second * 30)
		rawDb.SetConnMaxLifetime(variables.ConfigGormv2Yml.GetDuration("Gormv2.Mysql.SetConnMaxLifetime") * time.Second)
		rawDb.SetMaxIdleConns(variables.ConfigGormv2Yml.GetInt("Gormv2.Mysql.SetMaxIdleConns"))
		rawDb.SetMaxOpenConns(variables.ConfigGormv2Yml.GetInt("Gormv2.Mysql.SetMaxOpenConns"))
		return gormDb, nil
	}
}

//  根据配置参数生成数据库驱动 dsn
func getDsn(dbConf ...gorm_v2.ConfigParams) string {
	Host := variables.ConfigGormv2Yml.GetString("Gormv2.Mysql.Host")
	DataBase := variables.ConfigGormv2Yml.GetString("Gormv2.Mysql.DataBase")
	Port := variables.ConfigGormv2Yml.GetInt("Gormv2.Mysql.Port")
	User := variables.ConfigGormv2Yml.GetString("Gormv2.Mysql.User")
	Pass := variables.ConfigGormv2Yml.GetString("Gormv2.Mysql.Pass")
	Charset := variables.ConfigGormv2Yml.GetString("Gormv2.Mysql.Charset")

	if len(dbConf) > 0 {
		if len(dbConf[0].Host) > 0 {
			Host = dbConf[0].Host
		}
		if len(dbConf[0].DataBase) > 0 {
			DataBase = dbConf[0].DataBase
		}
		if dbConf[0].Port > 0 {
			Port = dbConf[0].Port
		}
		if len(dbConf[0].User) > 0 {
			User = dbConf[0].User
		}
		if len(dbConf[0].Pass) > 0 {
			Pass = dbConf[0].Pass
		}
		if len(dbConf[0].Charset) > 0 {
			Charset = dbConf[0].Charset
		}
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", User, Pass, Host, Port, DataBase, Charset)
}
