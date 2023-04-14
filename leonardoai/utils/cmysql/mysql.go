package cmysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

type Cfg struct {
	Host         string
	Port         string
	User         string
	Pass         string
	DbName       string
	Prefix       string
	MaxIdleConns int
	MaxOpenConns int
	MaxLifeTime  int // 分钟
	LogLevel     logger.LogLevel
}

// New
// @Description: 实例化
// @return *Cfg
func New(config *Cfg) *Cfg {
	return config
}

// WithConnect
// @Description: 数据库连接
// @receiver c
// @param config
// @return *gorm.DB
// @return error
func (c *Cfg) WithConnect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.User, c.Pass, c.Host, c.Port, c.DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   c.Prefix,
			SingularTable: true, // 是否设置单数表名，设置为 是
		},
		Logger: logger.Default.LogMode(c.LogLevel),
	})
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to the database, please check the MySQL configuration information first,the error details are:" + err.Error())
	}
	// GORM 使用 database/sql 维护连接池
	sqlDB, _ := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(c.MaxIdleConns)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(c.MaxOpenConns)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Duration(c.MaxLifeTime))
	return db, nil
}
