package model

type Setting struct {
	BaseModel
	ApiKey      string `gorm:"column:api_key;type:text" json:"api_key"`
	ModelName   string `gorm:"column:model_name;type:text" json:"model_name"`
	Language    string `gorm:"column:language;type:text" json:"language"`
	Theme       int    `gorm:"column:theme;type:int" json:"theme"`                // 主题 1：跟随系统 2：亮色模式 3：暗黑模式
	Account     string `gorm:"column:account;type:text" json:"account"`           // 账号
	AccessToken string `gorm:"column:access_token;type:text" json:"access_token"` // access_token
	IsSync      int    `gorm:"column:is_sync;type:int" json:"is_sync"`            // 是否同步 1：同步 2：不同步
	SyncTime    int    `gorm:"column:sync_time;type:int" json:"sync_time"`        // 同步时间 默认 5s
}
