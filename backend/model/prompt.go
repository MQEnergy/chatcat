package model

// Prompt
// @Description: 聊天表 主要包含聊天内容列表
type Prompt struct {
	BaseModel
	Name     string `gorm:"column:name;type:text" json:"name"`
	Desc     string `gorm:"column:desc;type:text" json:"desc"`          // 描述
	Prompt   string `gorm:"column:prompt;type:text" json:"prompt"`      // 提示词（CN）
	Enprompt string `gorm:"column:enprompt;type:text" json:"enprompt"`  // 提示词（EN）
	TagName  string `gorm:"column:tag_name;type:text" json:"tag_name"`  // tag列表 逗号分割 如：AI,常用...
	IsCommon int    `gorm:"column:is_common;type:int" json:"is_common"` // 是否常用 1：是 2：不是
}
