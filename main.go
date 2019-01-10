package main

import (
	"fmt"

	"gitlab.com/purini-to/stack_code_cov/msg"
)

// Message interface get message text.
type Message interface {
	GetMessage() string
}

func main() {
	world, _ := msg.InitializeHello("World")
	fmt.Println(getMessage(world))
}

func getMessage(message Message) string {
	return message.GetMessage()
}
