package messagesXFiles

type MessagesXFiles struct {
	FromId int64 `json:"id" gorm:"references:messages.id"`
	ToId   int64 `json:"to_id" gorm:"references:files.id"`
}
