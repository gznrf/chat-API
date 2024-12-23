package chats

import "gorm.io/gorm"

type Handler struct {
	dbHandler *gorm.DB
}

func NewHandler(dbHandler *gorm.DB) *Handler {
	return &Handler{dbHandler: dbHandler}
}
