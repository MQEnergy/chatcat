package model

// ChatCate chat分类表
type ChatCate struct {
	BaseModel
	Name   string `gorm:"column:name;type:text" json:"name"`
	Letter string `gorm:"column:letter" json:"letter"` // 首字母 或者 首字
	Color  string `gorm:"column:color" json:"color"`   // 颜色值
	Desc   string `gorm:"column:desc;type:text" json:"desc"`
}

// Chat chat对话表
type Chat struct {
	BaseModel
	CateId uint   `gorm:"column:cate_id;type:int unsigned" json:"cate_id"`
	Name   string `gorm:"column:name;type:text" json:"name"`
	Sort   int    `gorm:"column:sort" json:"sort"`
}

// ChatRecord chat对话记录表
type ChatRecord struct {
	BaseModel
	CateId  uint   `gorm:"column:cate_id;type:int unsigned" json:"cate_id"`
	ChatId  uint   `gorm:"column:chat_id;type:int unsigned" json:"chat_id"`
	Name    string `gorm:"column:name;type:text" json:"name"`
	Content string `gorm:"column:content;type:text" json:"content"`
	Role    string `gorm:"column:role;type:text" json:"role"` // 角色类型 system user assistant
}
