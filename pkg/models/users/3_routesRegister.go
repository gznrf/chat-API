package users

import (
	"github.com/gorilla/mux"
)

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/sign-up", h.signUpHandler).Methods("POST")
	router.HandleFunc("/sign-in", h.signInHandler).Methods("POST")
}
