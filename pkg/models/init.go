package models

import (
	"github.com/pro-cop/praktica/pkg/models/chatData"
	"github.com/pro-cop/praktica/pkg/models/chatDataXChats"
	"github.com/pro-cop/praktica/pkg/models/chats"
	"github.com/pro-cop/praktica/pkg/models/chatsXMessages"
	"github.com/pro-cop/praktica/pkg/models/codes"
	"github.com/pro-cop/praktica/pkg/models/codesXUsers"
	"github.com/pro-cop/praktica/pkg/models/files"
	"github.com/pro-cop/praktica/pkg/models/messages"
	"github.com/pro-cop/praktica/pkg/models/messagesXFiles"
	"github.com/pro-cop/praktica/pkg/models/passwords"
	"github.com/pro-cop/praktica/pkg/models/passwordsXUsers"
	"github.com/pro-cop/praktica/pkg/models/personalData"
	"github.com/pro-cop/praktica/pkg/models/personalDataXUsers"
	"github.com/pro-cop/praktica/pkg/models/tokens"
	"github.com/pro-cop/praktica/pkg/models/tokensXUsers"
	"github.com/pro-cop/praktica/pkg/models/users"
	"github.com/pro-cop/praktica/pkg/models/usersXChats"
	"github.com/pro-cop/praktica/pkg/models/usersXFiles"
	"github.com/pro-cop/praktica/pkg/models/usersXMessages"
	"reflect"

	"gorm.io/gorm"

	"log"
)

type Tables interface {
}

func Init(db *gorm.DB) {

	//Список таблиц
	tablesList := []Tables{users.Users{}, chats.Chats{}, messages.Messages{}, files.Files{},
		passwords.Passwords{}, personalData.PersonalData{}, chatsXMessages.ChatsXMessages{},
		messagesXFiles.MessagesXFiles{}, passwordsXUsers.PasswordsXUsers{}, personalDataXUsers.PersonalDataXUsers{},
		usersXChats.UserXChat{}, usersXFiles.UserXFile{}, usersXMessages.UsersXMessages{}, chatData.ChatData{},
		chatDataXChats.ChatDataXChats{}, codes.Codes{}, codesXUsers.CodesXUsers{}, tokens.Tokens{}, tokensXUsers.TokensXUsers{}}

	initTables(tablesList, db)
}

func initTables(tablesList []Tables, dbHandler *gorm.DB) {

	for _, table := range tablesList {

		if !dbHandler.Migrator().HasTable(&table) {

			if err := dbHandler.Migrator().CreateTable(&table); err != nil {
				log.Println(err)
				log.Println("ошибка создания таблицы {", reflect.TypeOf(table).Name(), "}")
			}

		} else {
			log.Println("таблица уже существует {", reflect.TypeOf(table).Name(), "}")
		}

	}

}
