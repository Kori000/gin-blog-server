package core

import (
	"gin-blog-server/global"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitGorm() *gorm.DB {

	if global.Config.Mysql.Host == "" {
		global.Log.Warnln("未配置 mysql, 取消 gorm 连接")
		return nil
	}

	dsn := global.Config.Mysql.Dsn()

	var mysqlLogger logger.Interface

	if global.Config.System.Env == "debug" {

		mysqlLogger = logger.Default.LogMode(logger.Info) // 开发环境打印所有的sql日志
	} else {
		mysqlLogger = logger.Default.LogMode(logger.Error) // 只打印错误的sql
	}

	// 初始化 global.MysqlLog 打印所有的sql日志
	global.MysqlLog = logger.Default.LogMode(logger.Info)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})

	if err != nil {
		global.Log.Fatalf("[%s] mysql 连接失败", dsn)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetConnMaxIdleTime(10) //最大空闲连接数
	sqlDB.SetMaxIdleConns(100)   //最大可容纳
	sqlDB.SetConnMaxLifetime(time.Hour * 4)

	return db
}
