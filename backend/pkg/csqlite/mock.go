package csqlite

import (
	"chatcat/backend/model"
	"gorm.io/gorm"
	"time"
)

func MockSetting(db *gorm.DB) {
	db.Create(&model.Setting{
		ApiKey:      "",
		AskModel:    "text-davinci-003",
		ChatModel:   "gpt-3.5-turbo",
		Language:    "en",
		Theme:       1,
		ProxyUrl:    "http://127.0.0.1:7890",
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
	db.Create(&[]model.ChatCate{
		{
			Name:   "未分类",
			Letter: "N",
			Color:  "#3370ff",
		}, {
			Name:   "对话",
			Letter: "C",
			Color:  "#14c9c9",
		}, {
			Name:   "问答",
			Letter: "A",
			Color:  "#ff7d00",
		},
	})
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
