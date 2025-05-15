package database_controller

import (
	"fmt"
	"soapstone/database_controller/message_database"
	"soapstone/message"
	"soapstone/printer"
)

func PrintSoapstone(cmd string) {
	printer.PrintMessages(cmd, message_database.GetMessages(cmd))
}

func VoteMessage(uuid string, like bool) {
}

func NewMessage(msg string, cmd string) {
	message_database.InsertMessage(message.NewMessage(msg, cmd), cmd)
	printer.PrintSuccess(fmt.Sprintf("Added '%s' soapstone message to '%s'", msg, cmd))
}

func RemoveMessage(cmd string, idx int) {
	messages := message_database.GetMessages(cmd)
	idx--

	if idx >= len(messages) {
		printer.PrintWarning("Index out of range")
		return
	}

	message_database.RemoveMessage(cmd, idx)
}

func Save() {
	message_database.WriteMessageDB()
}

func Load() {
	message_database.LoadMessageDB()
}
