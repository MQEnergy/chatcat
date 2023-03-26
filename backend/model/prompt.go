package model

// Prompt
// @Description: 聊天表 主要包含聊天内容列表
type Prompt struct {
	ID        uint   `gorm:"primaryKey"`
	Word      string `json:"word"` // 关键词
	CreatedAt string `json:"created_at"`
}
