package messages

import (
	"gorm.io/gorm"
	"time"
)

type Messages struct {
	ID        int64          `json:"id" gorm:"primaryKey; unique; not null"`
	ChatId    int64          `json:"chat_id" gorm:"not null"`
	SenderId  int64          `json:"sender_id" gorm:"not null"`
	Text      string         `json:"text" gorm:"type:text; not null"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime;type:timestamp"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime;type:timestamp"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"autoUpdateTime;type:timestamp"`
}
