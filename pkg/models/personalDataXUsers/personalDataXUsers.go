package personalDataXUsers

type PersonalDataXUsers struct {
	FromId int64 `json:"id" gorm:"references:personal_data.id"`
	ToId   int64 `json:"to_id" gorm:"references:users.id"`
}
