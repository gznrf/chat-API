package chatData

type ChatData struct {
	ID          int64  `json:"id" gorm:"primaryKey; unique;  not null"`
	Name        string `json:"name" gorm:"type:varchar(50);"`
	Description string `json:"description" gorm:"type:varchar(255);"`
}
