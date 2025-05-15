package message_database

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"soapstone/message"
	"soapstone/printer"
)

const DB_PATH = "./soapstone_cache"

var message_db map[string][]message.Message

func init() {
	message_db = make(map[string][]message.Message)

	LoadMessageDB()
}

func GetMessages(cmd string) []message.Message {
	return message_db[cmd]
}

func RemoveMessage(cmd string, idx int) {
	message_db[cmd] = slices.Delete(message_db[cmd], idx, idx+1)
}

func InsertMessage(msg *message.Message, cmd string) {
	message_db[cmd] = append(message_db[cmd], *msg)
}

func insert_testing_messages() {
	fmt.Println(printer.Dim + "Overidding message_db with testing messages" + printer.Reset)

	// Command: ls
	{
		cmd := "ls"
		InsertMessage(message.NewMessage("The first step to understanding chaos is listing it.", cmd), cmd)
		InsertMessage(message.NewMessage("Use `ls -lAh` and feel like a digital archaeologist.", cmd), cmd)
		InsertMessage(message.NewMessage("Ever run it in a directory you forgot existed? Spooky.", cmd), cmd)
		InsertMessage(message.NewMessage("So many files... but none remember you.", cmd), cmd)
		InsertMessage(message.NewMessage("Who needs order when you have `ls --sort=size`?", cmd), cmd)
	}

	// Command: rm
	{
		cmd := "rm"
		InsertMessage(message.NewMessage("One command. Infinite regret.", cmd), cmd)
		InsertMessage(message.NewMessage("I removed / once. I, too, was removed.", cmd), cmd)
		InsertMessage(message.NewMessage("The quietest scream you'll never hear.", cmd), cmd)
		InsertMessage(message.NewMessage("Like death, it comes swiftly and without prompt.", cmd), cmd)
		InsertMessage(message.NewMessage("rm -rf / is how gods commit mistakes.", cmd), cmd)
	}

	// Command: grep
	{
		cmd := "grep"
		InsertMessage(message.NewMessage("I see patterns. I see meaning. I see unfinished dreams.", cmd), cmd)
		InsertMessage(message.NewMessage("Used it to find secrets once. Found my old sins instead.", cmd), cmd)
		InsertMessage(message.NewMessage("grep is like a divining rod for lost intent.", cmd), cmd)
		InsertMessage(message.NewMessage("If code is a forest, grep is the whisper that guides you.", cmd), cmd)
		InsertMessage(message.NewMessage("`grep -r TODO`... now that's how you start a panic attack.", cmd), cmd)
	}

	// Command: touch
	{
		cmd := "touch"
		InsertMessage(message.NewMessage("I touch, therefore it exists. Simple.", cmd), cmd)
		InsertMessage(message.NewMessage("Sometimes I touch files just to feel alive.", cmd), cmd)
		InsertMessage(message.NewMessage("touch is the heartbeat of a file. A pulse in a cold system.", cmd), cmd)
		InsertMessage(message.NewMessage("Created a file named `doom.txt` once. Still afraid to open it.", cmd), cmd)
	}

}

func WriteMessageDB() error {
	file, err := os.Create(DB_PATH)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	e := encoder.Encode(message_db)

	if e != nil {
		panic(e)
	}

	return e
}

func LoadMessageDB() error {
	file, err := os.Open(DB_PATH)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&message_db)

	if err != nil {
		panic(err)
	}

	return err
}
