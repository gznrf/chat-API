package api

import (
	"github.com/pro-cop/praktica/pkg/models/chats"
	"github.com/pro-cop/praktica/pkg/models/tokens"
	"github.com/pro-cop/praktica/pkg/models/users"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Server struct {
	port      string
	dbHandler *gorm.DB
}

func NewServer(port string, dbHandler *gorm.DB) *Server {
	return &Server{
		port:      port,
		dbHandler: dbHandler,
	}
}

func (s *Server) Run() error {

	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userHandler := users.NewHandler(s.dbHandler)
	userHandler.RegisterRoutes(subrouter)

	chatHandler := chats.NewHandler(s.dbHandler)
	chatHandler.RegisterRoutes(subrouter)

	tokenHandler := tokens.NewHandler(s.dbHandler)
	tokenHandler.RegisterRoutes(subrouter)

	log.Println("Listening on localhost" + s.port)
	return http.ListenAndServe(s.port, router)
}
