package chats

import (
	"gorm.io/gorm"
	"time"
)

type Chats struct {
	ID        int64          `json:"id" gorm:"primaryKey; unique;  not null"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime;type:timestamp"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime;type:timestamp"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"autoUpdateTime;type:timestamp"`
}
