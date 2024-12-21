package main

import (
	"fmt"
	"main/internal/message/domain/model/message"
	"main/internal/message/domain/model/message/system"
	"main/internal/message/domain/model/message/text"
	inmemory2 "main/internal/message/infra/inmemory"
	"main/internal/model/user"
	user2 "main/internal/user/domain/model"
	"os"
	"time"
)

var users = map[user2.Login]*user2.User{
	"Test1": user2.New("Test1"),
	"Test2": user2.New("Test2"),
	"Test3": user2.New("Test3"),
	"Test4": user2.New("Test4"),
}

var message1 = text.NewMessage(users[1].Login(), []user2.Login{users[2].Login()}, time.Now(), "Первы")
var message2 = system.NewMessage([]user2.Login{users[2].Login(), users[3].Login(), users[4].Login()}, time.Now(), "test")

var messages = map[user2.Login][]message.Message{
	"Test2": {message1, message2},
	"Test3": {message2},
	"Test4": {message2},
}

func main() {
	startUpData()

	messageStorage := inmemory2.NewMessage(messages)
	userStorage := inmemory2.NewUserStorage(users)

	senderService := services.NewMessageSenderService(userStorage, messageStorage)
	editorService := services.NewEditorService(messageStorage)
	printerService := services.NewPrinterService(userStorage, messageStorage)

	fmt.Println("LOGIN: [Test1]")
	printerService.PrintMessagesForUser("Test1", os.Stdout)
	fmt.Println("LOGIN: [Test2]")
	printerService.PrintMessagesForUser("Test2", os.Stdout)
	fmt.Println("LOGIN: [Test3]")
	printerService.PrintMessagesForUser("Test3", os.Stdout)
	fmt.Println("LOGIN: [Test4]")
	printerService.PrintMessagesForUser("Test4", os.Stdout)
	var message4 = text.NewMessage("Четвертый курлык", time.Now(), users[2], []*user2.User{users[1]})
	senderService.SendMessage(message4)

	fmt.Println("------------------------------------------------------------------------------")
	fmt.Println("------------------------------------------------------------------------------")

	fmt.Println("LOGIN: [Test1]")
	printerService.PrintMessagesForUser("Test1", os.Stdout)
	fmt.Println("LOGIN: [Test2]")
	printerService.PrintMessagesForUser("Test2", os.Stdout)
	fmt.Println("LOGIN: [Test3]")
	printerService.PrintMessagesForUser("Test3", os.Stdout)
	fmt.Println("LOGIN: [Test4]")
	printerService.PrintMessagesForUser("Test4", os.Stdout)
	editorService.ChangeMessage(message4.Id(), "Пятый курлык")

	fmt.Println("------------------------------------------------------------------------------")
	fmt.Println("------------------------------------------------------------------------------")

	fmt.Println("Test1")
	printerService.PrintMessagesForUser("Test1", os.Stdout)
	fmt.Println("Test2")
	printerService.PrintMessagesForUser("Test2", os.Stdout)
	fmt.Println("Test3")
	printerService.PrintMessagesForUser("Test3", os.Stdout)
	fmt.Println("Test4")
	printerService.PrintMessagesForUser("Test4", os.Stdout)
}

func startUpData() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory: ", err)
		return
	}

	file, err := os.Open(cwd + "\\test.mp4")
	defer file.Close()
	if err != nil {
		fmt.Println("Error getting working directory: ", err)
		return
	}

	message3, err := multimedia.NewMessage(file, time.Now(), users[3], []*user2.User{users[1]})
	if err != nil {
		fmt.Println("Error getting working directory: ", err)
		return
	}

	messages["Test1"] = []user.Message{message3}
}
