package personalData

import "gorm.io/gorm"

type PersonalData struct {
	gorm.Model
	ID          int64  `json:"id" gorm:"primaryKey; unique;  not null"`
	Email       string `json:"email" gorm:"type:varchar(50); unique;  not null"`
	AltEmail    string `json:"alt_email" gorm:"type:varchar(50); unique;  not null"`
	PhoneNumber string `json:"phone_number" gorm:"type:varchar(20); unique;  not null"`
}
