package model

type BaseModel struct {
	ID        uint   `gorm:"primarykey"`
	CreatedAt uint64 `gorm:"column:created_at;type:bigint unsigned;NULL;" json:"created_at"`
	UpdatedAt uint64 `gorm:"column:updated_at;type:bigint unsigned;NULL;" json:"updated_at"`
}
