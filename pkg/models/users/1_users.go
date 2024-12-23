package users

import (
	"gorm.io/gorm"
	"time"
)

type Users struct {
	ID         int64     `json:"id" gorm:"primaryKey; unique;  not null"`
	Name       string    `json:"name" gorm:"type:varchar(40); unique;  not null"`
	IsVerified bool      `json:"is_verified" gorm:"default:false; not null"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime;type:timestamp"`
}

func getUser(dbHandler *gorm.DB, name string) (Users, error) {
	var user Users

	err := dbHandler.Table("users").Where("name = ?", name).Scan(&user).Error

	return user, err
}

func isExists(dbHandler *gorm.DB, username string) bool {
	var userStatus Users

	if err := dbHandler.Table("users").
		Select("name").
		Where("name = ?", username).
		Scan(&userStatus).Error; err != nil {
		return false
	}

	if userStatus.Name == "" {
		return false
	}

	return true
}
