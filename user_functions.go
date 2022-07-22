package main

import (
	"database/sql"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/lib/pq"
)

var db, db_err = sql.Open("postgres", "host=localhost port=5433 user=admin password=admin dbname=gotest sslmode=disable")

func UserHandler(update tgbotapi.Update) {
	// if !update.Message.IsCommand() {
	// 	return
	// }

	chat := Chat{update.Message.Chat.ID, update.Message.MessageID, strconv.FormatInt(update.Message.From.ID, 10), string(update.Message.From.UserName), string(update.Message.From.LanguageCode), string(update.Message.From.FirstName)}

	if update.Message.Location != nil {
		go SendMessage(chat.ID, getLocation(update.Message.Location.Latitude, update.Message.Location.Longitude))
		return
	}

	switch update.Message.Command() {
	case "solve":
		args := update.Message.CommandArguments()
		if args != "" {
			go SendMessage(chat.ID, solveAnagram(update.Message.CommandArguments()))
		} else {
			go SendMessage(chat.ID, "Используйте: /solve [слово]\nНапример: /solve нисьлепа")
		}

	case "make":
		args := update.Message.CommandArguments()
		if args != "" {
			go SendMessage(chat.ID, makeAnagram(update.Message.CommandArguments()))
		} else {
			go SendMessage(chat.ID, "Используйте: /make [слово]\nНапример: /make апельсин")
		}

	case "start":
		response, _ := db.Query("SELECT * FROM users WHERE user_id = $1", chat.user_id)
		if isUserNew(response, chat) {
			SendMessage(chat.ID, "Добро пожаловать, "+chat.first_name+"!")
		} else {
			go updateUserInfo(response, chat)
		}
		go SendHelpMessage(chat.ID)

	case "help":
		response, _ := db.Query("SELECT * FROM users WHERE user_id = $1", chat.user_id)
		go updateUserInfo(response, chat)
		go SendHelpMessage(chat.ID)

	default:
		go ReplyToMessageID(chat.ID, chat.messageID, "Неизвестная команда.")
	}
}

func isUserNew(response *sql.Rows, chat Chat) bool {
	err_count := 0
	if !response.Next() {
		for {
			_, err := db.Exec("INSERT INTO users VALUES ($1, $2, $3, $4)", chat.user_id, chat.username, chat.language, chat.first_name)
			if err == nil {
				break
			} else {
				err_count += 1
			}
			if err_count == 10 {
				break
			}
		}
		return true
	}
	return false
}

func updateUserInfo(response *sql.Rows, chat Chat) {
	res := User{}
	response.Scan(&res.user_id, &res.username, &res.language, &res.first_name)
	response.Close()

	err_count := 0
	if res.username != chat.username || res.language != chat.language || res.first_name != chat.first_name {
		for {
			_, err := db.Exec("UPDATE users SET username=$1, language=$2, first_name=$3 WHERE user_id=$4", chat.username, chat.language, chat.first_name, chat.user_id)
			if err == nil {
				break
			} else {
				err_count += 1
			}
			if err_count == 10 {
				break
			}
		}
	}
}
