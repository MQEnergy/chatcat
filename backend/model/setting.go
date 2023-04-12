package model

type Setting struct {
	BaseModel
	ApiKey      string `gorm:"column:api_key;type:text" json:"api_key"`
	ChatModel   string `gorm:"column:chat_model;type:text" json:"chat_model"`     // 对话模型
	AskModel    string `gorm:"column:ask_model;type:text" json:"ask_model"`       // 问答模型
	Language    string `gorm:"column:language;type:text" json:"language"`         // 默认语言
	Theme       int    `gorm:"column:theme;type:int" json:"theme"`                // 主题 1：跟随系统 2：亮色模式 3：暗黑模式
	ProxyUrl    string `gorm:"column:proxy_url;type:text" json:"proxy_url"`       // 代理地址
	Account     string `gorm:"column:account;type:text" json:"account"`           // 账号
	AccessToken string `gorm:"column:access_token;type:text" json:"access_token"` // access_token
	IsSync      int    `gorm:"column:is_sync;type:int" json:"is_sync"`            // 是否同步 1：同步 2：不同步
	SyncTime    int    `gorm:"column:sync_time;type:int" json:"sync_time"`        // 同步时间 默认 5s
}
