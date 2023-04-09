package model

type Tag struct {
	BaseModel
	Name string `gorm:"column:name;type:text" json:"name"`
}
