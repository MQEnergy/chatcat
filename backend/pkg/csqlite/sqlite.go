package csqlite

import (
	"chatcat/backend/model"
	"chatcat/backend/pkg/chelp"
	"chatcat/backend/pkg/clog"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	isAutoMigrate = false
	databasePath  = "./runtime/chatcat.db"
	db            *gorm.DB
	err           error
)

func WithConnect() (*gorm.DB, error) {
	clog.PrintInfo("sqlite db path: ", databasePath)
	if !chelp.IsPathExist(databasePath) {
		_, err := chelp.MakeFileOrPath(databasePath)
		if err != nil {
			return nil, err
		}
	}
	db, err = gorm.Open(sqlite.Open(databasePath), &gorm.Config{})
	clog.PrintInfo("======== 迁移数据开始 ========")
	// 迁移数据
	autoMigrate()
	clog.PrintInfo("======== 迁移数据结束 ========")
	// 迁移完成给出迁移完成标识
	isAutoMigrate = true
	return db, err
}

// GetIsAutoMigrate
// @Description: 获取是否已经执行migrate
// @return bool
// @author cx
func GetIsAutoMigrate() bool {
	return isAutoMigrate
}

func autoMigrate() {
	db.Migrator().CreateTable(
		&model.Prompt{},
	)
	// 写入初始化数据
}
