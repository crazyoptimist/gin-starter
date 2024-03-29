package common

import (
	"time"
)

// Same as gorm.Model, but including column and json tags
type BaseModel struct {
	ID        uint       `gorm:"primaryKey;column:id" json:"id"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}
