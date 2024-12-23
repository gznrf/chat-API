package tokens

import (
	"github.com/pro-cop/praktica/pkg/utils"
	"net/http"
	"time"
)

func (h *Handler) logoutHandler(w http.ResponseWriter, r *http.Request) {

	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    "",
		Expires:  time.Unix(0, 0),
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
	})

	utils.WriteJson(w, 200, map[string]string{"message": "Logged out successfully"})
}
