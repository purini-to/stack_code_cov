package main

import (
	"fmt"

	"gitlab.com/purini-to/stack_code_cov/msg"
)

func main() {
	world, _ := msg.InitializeMessage("World")
	fmt.Println(world.GetMessage())
}
