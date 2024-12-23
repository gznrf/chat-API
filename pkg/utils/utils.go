package utils

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

// получение хеша пароля
func GetHash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// метод для сверки хешей принимает обычный пароль и его заранее захешированную версию далее
// далее возвращает совпадают пароли или нет
func CheckHash(originalPassword, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(originalPassword))
	return err == nil
}

// декодировка jsona
func DecodeJson(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("request body is nil")
	}

	return json.NewDecoder(r.Body).Decode(payload)
}

// Pishet json
func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func ParseToInt64(s string) (int64, error) {
	i64, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return i64, nil
}

// Pishet oshibku
func WriteError(w http.ResponseWriter, status int, err error) {
	if err := WriteJson(w, status, map[string]string{"error": err.Error()}); err != nil {
		WriteError(w, http.StatusInternalServerError, err)
	}
}
