package users

import (
	"errors"
	"github.com/pro-cop/praktica/pkg/models/tokens"
	"github.com/pro-cop/praktica/pkg/models/tokensXUsers"
	"github.com/pro-cop/praktica/pkg/utils"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type signInInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) signInHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:9001")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var input signInInput

	if err := utils.DecodeJson(r, &input); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if !isExists(h.dbHandler, input.Username) {
		utils.WriteError(w, http.StatusConflict, errors.New("данного пользователя не существует"))
		return
	}

	if !isPasswordCorrect(h.dbHandler, input.Username, input.Password) {
		utils.WriteError(w, http.StatusConflict, errors.New("пароль неверный"))
		return
	}

	//получаем токен из куки
	cookie, _ := r.Cookie("auth_token")

	//если токен уже есть то просто логиним пользователя
	if cookie != nil {
		//отправляем успешную авторизацию
		response := map[string]string{"is_authenticated": "true"}
		_ = utils.WriteJson(w, 200, response)
		return
	}

	//если нет токена то генерируем его
	//удаляем старый токен из базы если он есть
	deleteToken(h.dbHandler, input.Username)

	//создаем новый токен
	token, err := generateToken(h.dbHandler, input.Username)
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, err)
		return
	}

	//ставим его в куки
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    token,
		Expires:  time.Now().Add(12 * time.Hour),
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
	})

	response := map[string]string{"is_authenticated": "true"}
	_ = utils.WriteJson(w, 200, response)
}

func isPasswordCorrect(dbHandler *gorm.DB, username, password string) bool {
	var correctPassword string

	dbHandler.Table("passwords").Select("password").
		Joins("JOIN passwords_x_users ON passwords_x_users.from_id = passwords.id").
		Joins("JOIN users ON passwords_x_users.to_id = users.id").
		Where("users.name = ? ", username).
		Scan(&correctPassword)

	if utils.CheckHash(password, correctPassword) {

		return true
	}

	if checkTwoHashedPassword(correctPassword, password) {
		return true
	}

	return false
}

// так как при автоматическом логине пароль приходит в хеше есть такой метод который проверяет на сходство два пароля
func checkTwoHashedPassword(passwordFromDb, inputPassword string) bool {
	if passwordFromDb != inputPassword {
		return false
	}
	return true
}

func deleteToken(dbHandler *gorm.DB, username string) {
	var userId int64
	var tokenId int64

	//получаем айди юзера
	dbHandler.Table("users").
		Select("users.id").
		Where("users.name = ?", username).
		Scan(&userId)

	//получаем айди токена
	dbHandler.Table("tokens_x_users").
		Select("tokens_x_users.to_id").
		Where("tokens_x_users.from_id = ?", userId).
		Scan(&tokenId)

	//удаляем запись в связующей таблице
	dbHandler.Table("tokens_x_users").
		Where("to_id = ?", tokenId).
		Delete(tokensXUsers.TokensXUsers{})

	//удаляем токен из бд
	dbHandler.Table("tokens").
		Where("id = ?", tokenId).
		Delete(tokens.Tokens{})
}

func generateToken(dbHandler *gorm.DB, username string) (string, error) {
	var user Users
	var tokensStruct tokens.Tokens
	var tokensUsers tokensXUsers.TokensXUsers

	//получаем пользователя
	user, err := getUser(dbHandler, username)
	if err != nil {
		return "", err
	}

	//генерируем аксес токен и записываем его в структуру токена
	tokensStruct.AccessToken, err = tokens.NewJWT()
	if err != nil {
		return "", err
	}

	/*//генерируем рефреш токен и записываем его в структуру рефреш токена
	tokensStruct.RefreshToken, err = tokens.NewRefreshToken()
	if err != nil {
		return "", err
	}*/

	//добавляем время жизни токена
	tokensStruct.ExpiresAt = time.Now().Add(12 * time.Hour)

	//вносим в таблицу токены
	dbHandler.Table("tokens").Create(&tokensStruct)

	//записываем айди пользователя и токенов в структуру связующей таблицы
	tokensUsers.FromId = tokensStruct.ID
	tokensUsers.ToId = user.ID

	//записываем структуру в таблицу
	dbHandler.Table("tokens_x_users").Create(&tokensUsers)

	return tokensStruct.AccessToken, nil
}
