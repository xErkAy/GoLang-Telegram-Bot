package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func SendHelpMessage(ChatID int64) {
	bot.Send(tgbotapi.NewMessage(ChatID, "Доступные команды:\n/help - справка\n/solve - решить анаграмму\n/make - составить анаграмму"))
}

func SendMessage(ChatID int64, message string) {
	bot.Send(tgbotapi.NewMessage(ChatID, message))
}

func ReplyToMessageID(ChatID int64, messageID int, Text string) {
	message := tgbotapi.NewMessage(ChatID, Text)
	message.ReplyToMessageID = messageID
	bot.Send(message)
}

func SendPhoto(ChatID int64, FilePath string) {
	bot.Send(tgbotapi.NewPhoto(ChatID, tgbotapi.FilePath(FilePath)))
}

func ReplyWithPhoto(ChatID int64, messageID int, FilePath string) {
	message := tgbotapi.NewPhoto(ChatID, tgbotapi.FilePath(FilePath))
	message.ReplyToMessageID = messageID
	bot.Send(message)
}

func SendAudioMessage(ChatID int64, FilePath string) {
	bot.Send(tgbotapi.NewAudio(ChatID, tgbotapi.FilePath(FilePath)))
}

func ReplyWithAudioMessage(ChatID int64, messageID int, FilePath string) {
	message := tgbotapi.NewAudio(ChatID, tgbotapi.FilePath(FilePath))
	message.ReplyToMessageID = messageID
	bot.Send(message)
}

func SendVideo(ChatID int64, FilePath string) {
	bot.Send(tgbotapi.NewVideo(ChatID, tgbotapi.FilePath(FilePath)))
}

func ReplyWithVideo(ChatID int64, messageID int, FilePath string) {
	message := tgbotapi.NewVideo(ChatID, tgbotapi.FilePath(FilePath))
	message.ReplyToMessageID = messageID
	bot.Send(message)
}

func SendVideoURL(ChatID int64, FileURL string) {
	bot.Send(tgbotapi.NewVideo(ChatID, tgbotapi.FileURL(FileURL)))
}

func ReplyWithVideoURL(ChatID int64, messageID int, FileURL string) {
	message := tgbotapi.NewVideo(ChatID, tgbotapi.FileURL(FileURL))
	message.ReplyToMessageID = messageID
	bot.Send(message)
}

func SendDocument(ChatID int64, FilePath string) {
	bot.Send(tgbotapi.NewDocument(ChatID, tgbotapi.FilePath(FilePath)))
}

func ReplyWithDocument(ChatID int64, messageID int, FilePath string) {
	message := tgbotapi.NewDocument(ChatID, tgbotapi.FilePath(FilePath))
	message.ReplyToMessageID = messageID
	bot.Send(message)
}

func SendSticker(ChatID int64, FilePath string) {
	bot.Send(tgbotapi.NewSticker(ChatID, tgbotapi.FilePath(FilePath)))
}

func SendLocation(ChatID int64, Title string, Address string, latitude float64, longitude float64) {
	bot.Send(tgbotapi.NewVenue(ChatID, Title, Address, latitude, longitude))
}

func ReplyWithLocation(ChatID int64, messageID int, Title string, Address string, latitude float64, longitude float64) {
	message := tgbotapi.NewVenue(ChatID, Title, Address, latitude, longitude)
	message.ReplyToMessageID = messageID
	bot.Send(message)
}

func CreatePoll(ChatID int64, Title string, options ...string) {
	bot.Send(tgbotapi.NewPoll(ChatID, Title, options...))
}
