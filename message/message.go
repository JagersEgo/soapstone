package message

type Message struct {
	Text    string // The left message
	Command string // Command message is for
}

var EMPTY_MESSAGE = Message{Text: "", Command: ""}

func NewMessage(text string, command string) *Message {
	n := Message{Text: text, Command: command}
	return &n
}
