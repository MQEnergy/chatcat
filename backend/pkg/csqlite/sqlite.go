package csqlite

import (
	"chatcat/backend/config"
	"chatcat/backend/model"
	"chatcat/backend/pkg/chelp"
	"chatcat/backend/pkg/clog"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	isAutoMigrate bool
	databasePath  string
	db            *gorm.DB
	err           error
)

// WithConnect
// @Description: sqlite with connect
// @param conf
// @return *gorm.DB
// @return error
// @author cx
func WithConnect(conf *config.Conf) (*gorm.DB, error) {
	runtimeDir := chelp.GetRuntimeUserHomeDir()
	// ${userHome}/chatcat/chatcat.db
	databasePath := runtimeDir + "/" + conf.App.AppName + "/" + conf.App.DbName
	clog.PrintInfo("sqlite db path: ", databasePath)
	// check if db already exists Todo
	if !chelp.IsPathExist(databasePath) {
		_, err := chelp.MakeFileOrPath(databasePath)
		if err != nil {
			return nil, err
		}
		db, err = gorm.Open(sqlite.Open(databasePath), &gorm.Config{})
		clog.PrintInfo("======== migrate start ========")
		autoMigrate()
		clog.PrintInfo("======== migrate end ========")
		isAutoMigrate = true
	} else {
		db, err = gorm.Open(sqlite.Open(databasePath), &gorm.Config{})
	}
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
func autoMigrate() {
	db.Migrator().CreateTable(
		&model.Chat{},
		&model.ChatCate{},
		&model.ChatRecord{},
		&model.Prompt{},
		&model.Setting{},
		&model.Tag{},
	)
	// 写入初始化数据
	MockTagList(db)
	MockChatCate(db)
	MockPromptList(db)
}
