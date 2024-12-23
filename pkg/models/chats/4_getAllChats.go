package chats

import (
	"github.com/pro-cop/praktica/pkg/utils"
	"gorm.io/gorm"
	"net/http"
)

type getAllChatsInput struct {
	UserId int64 `json:"user_id"`
}

type chatListOutput struct {
	ChatId      int64  `json:"chat_id" gorm:"column:id"`
	ChatName    string `json:"chat_name" gorm:"column:name"`
	LastMessage string `json:"last_message"`
	ChatAvatar  string `json:"chat_avatar"`
}

func (h *Handler) getAllChatsHandler(w http.ResponseWriter, r *http.Request) {
	var input getAllChatsInput

	if err := utils.DecodeJson(r, &input); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	chatList := getChatsByUserId(h.dbHandler, input.UserId)

	if err := utils.WriteJson(w, http.StatusOK, chatList); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}

func getChatsByUserId(dbHandler *gorm.DB, userId int64) []chatListOutput {
	var chats []chatListOutput
	var userChatsIds []int

	//получаем список айди чатов нашего пользователя
	if err := dbHandler.Table("users").
		Select("chats.id").
		Joins("JOIN user_x_chats ON users.id = user_x_chats.from_id").
		Joins("JOIN chats ON chats.id = user_x_chats.to_id").
		Where("users.id = ?", userId).
		Scan(&userChatsIds).Error; err != nil {
		return nil
	}

	for _, chatId := range userChatsIds {
		//делаем выборку для каждого чата в котором состоит наш пользователь
		var newChat chatListOutput
		if err := dbHandler.Table("users").
			Select("chats.id, users.name").
			Joins("JOIN user_x_chats ON users.id = user_x_chats.from_id").
			Joins("JOIN chats ON chats.id = user_x_chats.to_id").
			Where("chats.id = ? AND users.id != ?", chatId, userId).
			Scan(&newChat).Error; err != nil {
			return nil
		}
		//добавляем чат в наш список чатов
		chats = append(chats, newChat)
	}

	return chats
}
