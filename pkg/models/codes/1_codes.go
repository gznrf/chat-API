package codes

import (
	"time"
)

type Codes struct {
	ID             int64     `json:"id" gorm:"primaryKey; unique;  not null"`
	Code           string    `json:"name" gorm:"type:varchar(40); unique;  not null"`
	ExpirationTime time.Time `json:"expiration_time" gorm:"type:timestamp;not null"`
}
