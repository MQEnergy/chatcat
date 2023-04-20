package csqlite

import (
	"chatcat/backend/model"
	"gorm.io/gorm"
	"time"
)

func MockSetting(db *gorm.DB) {
	db.Create(&model.Setting{
		ApiKey:           "",
		AskModel:         "gpt-3.5-turbo",
		ChatModel:        "gpt-3.5-turbo",
		Language:         "zh-CN",
		Theme:            1,
		ProxyUrl:         "http://127.0.0.1:7890",
		Account:          "",
		AccessToken:      "",
		IsSync:           2,
		SyncTime:         5,
		Temperature:      "0.7",
		MaxTokens:        0,
		PresencePenalty:  "0.6",
		FrequencyPenalty: "0.6",
		N:                1,
		TopP:             "0.1",
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
