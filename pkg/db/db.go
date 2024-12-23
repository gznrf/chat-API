package db

import (
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Init() *gorm.DB {
	dbURL := viper.GetString("database.url")
	if dbURL == "" {
		log.Println("file is not exist")
	}

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Println(err)
	} else {
		log.Println("DataBase successfully connected!")
	}

	return db
}

func CloseConnection(db *gorm.DB) {
	sqlConn, _ := db.DB()
	if err := sqlConn.Close(); err != nil {
		log.Printf("failed to close database connection: %v\n", err)
	}
}
