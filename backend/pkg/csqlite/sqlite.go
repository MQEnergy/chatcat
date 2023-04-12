package csqlite

import (
	"chatcat/backend/config"
	"chatcat/backend/model"
	"chatcat/backend/pkg/chelp"
	"chatcat/backend/pkg/clog"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

var (
	isAutoMigrate bool
	databasePath  string
)

// WithConnect
// @Description: sqlite with connect
// @param logger
// @param conf
// @return *gorm.DB
// @return error
// @author cx
func WithConnect(logger *logrus.Logger, conf *config.Conf) (*gorm.DB, error) {
	var (
		db  *gorm.DB
		err error
	)
	runtimeDir := chelp.GetRuntimeUserHomeDir()
	databasePath = runtimeDir + "/" + conf.App.AppName + "/" + conf.App.DbName
	if !chelp.IsPathExist(databasePath) {
		db, err = gorm.Open(sqlite.Open(databasePath), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		clog.PrintInfo("======== migrate start ========")
		logger.Info("======== migrate start ========")
		autoMigrate(db)
		clog.PrintInfo("======== migrate end ========")
		logger.Info("======== migrate end ========")
		isAutoMigrate = true
	} else {
		logger.Info("databasePath2", databasePath)
		db, err = gorm.Open(sqlite.Open(databasePath), &gorm.Config{})
		if err != nil {
			return nil, err
		}
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db, err
}

// GetDatabasePath get the database path
func GetDatabasePath() string {
	return databasePath
}

// GetIsAutoMigrate
// @Description: get migrate is or not executed
// @return bool
// @author cx
func GetIsAutoMigrate() bool {
	return isAutoMigrate
}

// autoMigrate ...
func autoMigrate(db *gorm.DB) {
	db.Migrator().CreateTable(
		&model.Setting{},
		&model.Chat{},
		&model.ChatCate{},
		&model.ChatRecord{},
		&model.Prompt{},
		&model.Tag{},
	)
	// 写入初始化数据
	MockSetting(db)
	MockChatCate(db)
	MockTagList(db)
	MockPromptList(db)
}
