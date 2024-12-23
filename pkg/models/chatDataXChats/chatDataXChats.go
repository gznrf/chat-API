package chatDataXChats

type ChatDataXChats struct {
	FromId int64 `json:"id" gorm:"foreignKey:chats_data.id"`
	ToId   int64 `json:"to_id" gorm:"foreignKey:chats.id"`
}
