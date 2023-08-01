package service

import (
	"devops-platform/internal/common/config"
	"devops-platform/internal/common/database/internal/domain"
	"devops-platform/internal/common/log"
	"devops-platform/pkg/beans"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

type databaseConfig interface {
	GetDataSourceName() string
	GetMaxIdle() int
	GetMaxOpen() int
	GetMaxConnectionLifetime() time.Duration
}

type DB struct {
	*gorm.DB
}

func (db *DB) StopOrder() int {
	return 1
}
func (db *DB) Stop() {
	logrus.Info("即将关闭数据库链接池")

	if sqlDB, err := db.DB.DB(); err != nil {
		logrus.WithError(err).Error("获取数据库链接池异常")
	} else if err = sqlDB.Close(); err != nil {
		logrus.WithError(err).Error("关闭数据库链接池异常")
	}

}

func (db *DB) PreInject(getBean func(string) interface{}) {

	logrus.Info("db pre inject")
	databaseConfig, ok := getBean(config.BeanDatabase).(databaseConfig)
	if !ok {
		logrus.Panicf("初始化时获取[%s]失败", config.BeanDatabase)
		return
	}

	logs, ok := getBean(log.BeanLog).(logger.Interface)
	if !ok {
		logrus.Panicf("初始化时获取[%s]失败", log.BeanLog)
		return
	}
	var err error

	db.DB, err = gorm.Open(mysql.Open(databaseConfig.GetDataSourceName()), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，启用该选项
		},
		Logger: logs,
	})
	if err != nil {
		logrus.Panic("初始化数据库链接失败", err.Error())
		return
	}

	sqlDB, err := db.DB.DB()

	if err != nil {
		logrus.Panic("初始化数据库链接时获取数据源链接池失败", err.Error())
		return
	}

	//设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(databaseConfig.GetMaxIdle())

	//设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(databaseConfig.GetMaxOpen())

	//设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(databaseConfig.GetMaxConnectionLifetime())

	beans.Register(domain.BeanDB, db.DB)

}
