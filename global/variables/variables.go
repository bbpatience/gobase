package variables

import (
	"go.uber.org/zap"
	"gobase/utils/yml_config/ymlconfig_interf"
	"gorm.io/gorm"
)

var (
	// zap logger
	ZapLog *zap.Logger
	// env config
	ConfigYml       ymlconfig_interf.YmlConfigInf
	// mysql config
	ConfigGormv2Yml ymlconfig_interf.YmlConfigInf

	//gorm 数据库客户端，如果您操作数据库使用的是gorm，请取消以下注释，在 bootstrap>init 文件，进行初始化即可使用
	GormDbMysql      *gorm.DB // 全局gorm的客户端连接
)
