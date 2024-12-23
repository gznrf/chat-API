package chats

import (
	"fmt"
	"github.com/pro-cop/praktica/pkg/models/users"
	"github.com/pro-cop/praktica/pkg/models/usersXChats"
	"github.com/pro-cop/praktica/pkg/utils"
	"gorm.io/gorm"
	"net/http"
)

type CreateNewChatInput struct {
	FirstUserId    int64  `json:"first-user-id"`
	SecondUsername string `json:"second-username"`
}

func (h *Handler) createNewChatHandler(w http.ResponseWriter, r *http.Request) {
	var input CreateNewChatInput

	if err := utils.DecodeJson(r, &input); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := addChatToDb(h.dbHandler, input.FirstUserId, input.SecondUsername); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if err := utils.WriteJson(w, http.StatusCreated, fmt.Sprint("чат был создан с пользователем "+input.SecondUsername)); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}
}

// функция получет айди второго пользователя с которым будет чат а затем добовляет обоих
// в связующую таблицу в базе
func addChatToDb(dbHandler *gorm.DB, firstUserId int64, secondUsername string) error {
	var newChat Chats
	var secondUser users.Users

	//получаем айди второго пользователя которому хочет написать наш пользователь
	if err := dbHandler.Table("users").
		Select("id").
		Where("name = ?", secondUsername).
		Scan(&secondUser).Error; err != nil {
		return err
	}

	//создаем новый чат создавая его через структуру чтобы сразу получить айди
	if err := dbHandler.Create(&newChat).Error; err != nil {
		return err
	}

	//создаем две связи с новым чатом
	//наш пользователь
	if err := dbHandler.Create(&usersXChats.UserXChat{FromId: firstUserId, ToId: newChat.ID}).Error; err != nil {
		return err
	}
	//пользователь которому пишет наш пользователь
	if err := dbHandler.Create(&usersXChats.UserXChat{FromId: secondUser.ID, ToId: newChat.ID}).Error; err != nil {
		return err
	}

	return nil
}
