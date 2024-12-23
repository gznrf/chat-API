package codesXUsers

type CodesXUsers struct {
	FromId int64 `json:"id" gorm:"references:codes.id"`
	ToId   int64 `json:"to_id" gorm:"references:users.id"`
}
