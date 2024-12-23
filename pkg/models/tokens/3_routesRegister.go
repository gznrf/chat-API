package tokens

import "github.com/gorilla/mux"

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/check-token", h.checkToken).Methods("POST")
	router.HandleFunc("/log-out", h.logoutHandler).Methods("POST")
}
