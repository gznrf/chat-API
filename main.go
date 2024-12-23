package main

import (
	"github.com/pro-cop/praktica/pkg/api"
	"github.com/pro-cop/praktica/pkg/db"
	"github.com/pro-cop/praktica/pkg/models"
	"github.com/pro-cop/praktica/pkg/viper"
	"log"
)

func main() {

	viper.SetConfiguration()
	dbHandler := db.Init()
	models.Init(dbHandler)

	server := api.NewServer(":3333", dbHandler)
	if err := server.Run(); err != nil {
		log.Println(err)
	}
}
