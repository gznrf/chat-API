package tokens

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/pro-cop/praktica/pkg/utils"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type UserPassword struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) checkToken(w http.ResponseWriter, r *http.Request) {

	//получаем токен из куки
	cookie, err := r.Cookie("auth_token")

	//если при получении произошла ошибка то отправляем 401
	if err != nil {
		utils.WriteError(w, 401, errors.New("ошибка в получении токена"+
			"возможно его нет"))
		return
	}

	//записываем значение токена в переменную
	tokenString := cookie.Value

	//проверяем токен на валидность
	isTValid, err := isTokenValid(tokenString)
	if err != nil {
		//выводим причину невалидности токена
		utils.WriteError(w, 401, err)
	}

	if !isTValid {
		utils.WriteError(w, 401, errors.New("токен не валидный"))
		return
	}

	_ = utils.WriteJson(w, 200, "token is valid")
}

func isTokenValid(tokenString string) (bool, error) {
	secret := viper.GetString("jwt.secret")
	claims := &jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	// Ошибка парсинга токена
	if err != nil {
		return false, err
	}

	// Проверяем валиден ли токен
	if !token.Valid {
		return false, errors.New("invalid token")
	}

	// проверяем не истекло ли его время
	if claims.ExpiresAt < time.Now().Unix() {
		return false, errors.New("tokens time out")
	}

	return true, nil
}

func getUserByToken(dbHandler *gorm.DB, token string) UserPassword {
	var userPassword UserPassword

	//получаем имя пользователя по его токену
	dbHandler.Table("tokens").
		Select("users.name").
		Joins("JOIN tokens_x_users ON tokens_x_users.from_id = tokens.id").
		Joins("JOIN users ON tokens_x_users.to_id = users.id").
		Where("tokens.access_token = ?", token).
		Scan(&userPassword.Username)

	//получаем пароль по имени пользователя
	dbHandler.Table("passwords").
		Select("passwords.password").
		Joins("JOIN passwords_x_users ON passwords_x_users.from_id = passwords.id").
		Joins("JOIN users ON passwords_x_users.to_id = users.id").
		Where("users.name = ?", userPassword.Username).
		Scan(&userPassword.Password)

	return userPassword
}
