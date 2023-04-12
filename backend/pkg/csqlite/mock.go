package csqlite

import (
	"chatcat/backend/model"
	"gorm.io/gorm"
	"time"
)

func MockSetting(db *gorm.DB) {
	db.Create(&model.Setting{
		ApiKey:      "",
		ModelName:   "gpt-3.5-turbo",
		Language:    "en",
		Theme:       1,
		Account:     "",
		AccessToken: "",
		IsSync:      2,
		SyncTime:    5,
		BaseModel: model.BaseModel{
			CreatedAt: uint64(time.Now().Unix()),
			UpdatedAt: uint64(time.Now().Unix()),
		},
	})
}

// MockChatCate ...
func MockChatCate(db *gorm.DB) {

}

// MockTagList
// @Description: tag list
// @param db
// @author cx
func MockTagList(db *gorm.DB) {

}

// MockPromptList
// @Description: prompt list
// @param db
// @author cx
func MockPromptList(db *gorm.DB) {

}
