package printer

import (
	"fmt"
	"soapstone/message"
)

func PrintMessages(cmd string, msg []message.Message) {
	fmt.Println(BoldYellow + "Soapstone messages found for '" + cmd + "':" + Reset)

	for i, msg := range msg {
		fmt.Printf(
			"  %d. %s%s %s\n",
			i+1, Reset+BoldYellow, msg.Text, Reset,
		)
	}
}

func PrintWarning(warning string) {
	fmt.Println(BoldRed + "Soapstone error: " + Reset + Red + warning + Reset)
}

func PrintSuccess(warning string) {
	fmt.Println(BoldGreen + "Soapstone: " + Reset + Green + warning + Reset)
}
