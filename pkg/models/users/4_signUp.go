package users

import (
	"errors"
	"github.com/pro-cop/praktica/pkg/models/passwords"
	"github.com/pro-cop/praktica/pkg/models/passwordsXUsers"
	"github.com/pro-cop/praktica/pkg/utils"
	"gorm.io/gorm"
	"net/http"
)

type signUpInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signUpHandler(w http.ResponseWriter, r *http.Request) {
	var input signUpInput

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	//записываем данные пользователя в структуру
	if err := utils.DecodeJson(r, &input); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	//получаем хеш пароля
	hashedPassword, err := utils.GetHash(input.Password)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	//проверяем нет ли в базе данного пользователя
	if isExists(h.dbHandler, input.Username) {
		utils.WriteError(w, http.StatusConflict, errors.New("пользователь уже существует"))
		return
	}

	//создаем пользователя с полученными логином и паролем
	if err := createUser(h.dbHandler, input.Username, hashedPassword); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	_ = utils.WriteJson(w, 200, "пользователь зарегестрирован")
}

// Функция для записи юзера в базу вместе с его паролем, связывая их
func createUser(dbHandler *gorm.DB, username, password string) error {

	var userStruct Users
	userStruct.Name = username
	if err := dbHandler.Create(&userStruct).Error; err != nil {
		return err
	}

	var passwordStruct passwords.Passwords
	passwordStruct.Password = password
	if err := dbHandler.Create(&passwordStruct).Error; err != nil {
		return err
	}

	if err := dbHandler.Create(&passwordsXUsers.PasswordsXUsers{FromId: passwordStruct.ID, ToId: userStruct.ID}).Error; err != nil {
		return err
	}

	return nil
}

// Удаляем созданного пользователя по его паролю и юзернейму, так же в связующей таблице
func deleteUser(dbHandler *gorm.DB, username, hashedPassword string) error {
	var userId int64

	if err := dbHandler.Table("users").Select("id").Where("name = ?", username).Scan(&userId).Error; err != nil {
		return err
	}

	if err := dbHandler.Table("users").Where("name = ?", username).Delete(Users{}).Error; err != nil {
		return err
	}
	if err := dbHandler.Table("passwords").Where("password = ?", hashedPassword).Delete(passwords.Passwords{}).Error; err != nil {
		return err
	}

	if err := dbHandler.Where("passwords_x_users.to_id = ?", userId).Delete(passwordsXUsers.PasswordsXUsers{}).Error; err != nil {
		return err
	}

	return nil
}
