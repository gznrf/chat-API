package usersXFiles

type UserXFile struct {
	FromId int64 `json:"id" gorm:"references:users.id"`
	ToId   int64 `json:"to_id" gorm:"references:files.id"`
}