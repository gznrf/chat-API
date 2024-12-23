package chatsXMessages

type ChatsXMessages struct {
	FromId int64 `json:"id" gorm:"foreignKey:chats.id"`
	ToId   int64 `json:"to_id" gorm:"foreignKey:messages.id"`
}
