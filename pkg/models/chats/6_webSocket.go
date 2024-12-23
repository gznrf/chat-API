package chats

import (
	"github.com/gorilla/websocket"
	"github.com/pro-cop/praktica/pkg/models/chatsXMessages"
	"github.com/pro-cop/praktica/pkg/models/messages"
	"github.com/pro-cop/praktica/pkg/utils"
	"net/http"
)

type UserConnection struct {
	Conn   *websocket.Conn
	ChatId int64
}

var connections = make(map[int64]*UserConnection)
var upgrader = websocket.Upgrader{}

func (h Handler) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	userID, err := utils.ParseToInt64(r.URL.Query().Get("user_id"))
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	chatID, err := utils.ParseToInt64(r.URL.Query().Get("chat_id"))
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		utils.WriteError(w, 500, err)
		return
	}

	connections[userID] = &UserConnection{Conn: conn, ChatId: chatID}
	defer func() {
		conn.Close()
		delete(connections, userID)
	}()

	for {
		var message messages.Messages
		conn.ReadJSON(&message)

		if message.Text == "" {
			continue
		}

		h.dbHandler.Create(&message)

		h.dbHandler.Create(&chatsXMessages.ChatsXMessages{FromId: message.ChatId, ToId: message.ID})

		for _, userConn := range connections {
			if userConn.ChatId == message.ChatId { // Отправляем только тем пользователям, которые находятся в этом чате
				userConn.Conn.WriteJSON(message)
			}
		}

	}

}
