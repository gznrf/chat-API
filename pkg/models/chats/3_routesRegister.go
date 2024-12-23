package chats

import (
	"github.com/gorilla/mux"
)

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/get-all-chats", h.getAllChatsHandler).Methods("POST")
	router.HandleFunc("/new-chat", h.createNewChatHandler).Methods("POST")
	router.HandleFunc("/ws", h.handleWebSocket)
}
