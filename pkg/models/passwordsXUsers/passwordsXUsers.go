package passwordsXUsers

type PasswordsXUsers struct {
	FromId int64 `json:"id" gorm:"references:passwords.id"`
	ToId   int64 `json:"to_id" gorm:"references:users.id"`
}
