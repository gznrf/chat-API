package tokens

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"time"
)

type Tokens struct {
	ID          int64  `json:"id" gorm:"primaryKey; unique;  not null"`
	AccessToken string `json:"access_token" gorm:"type:text"`
	//RefreshToken string    `json:"refresh_token" gorm:"type:text"`
	ExpiresAt time.Time `json:"expires_at" gorm:"autoCreateTime;type:timestamp"`
}

func NewJWT() (string, error) {
	secret := viper.GetString("jwt.secret")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		})

	return token.SignedString([]byte(secret))
}

/*func NewRefreshToken() (string, error) {
	b := make([]byte, 32)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	_, err := r.Read(b)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}*/
