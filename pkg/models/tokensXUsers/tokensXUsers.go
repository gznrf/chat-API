package tokensXUsers

type TokensXUsers struct {
	FromId int64 `json:"id" gorm:"references:tokens.id"`
	ToId   int64 `json:"to_id" gorm:"references:users.id"`
}
