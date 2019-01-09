package main

import (
	"fmt"

	"gitlab.com/purini-to/stack_code_cov/msg"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Hello() + msg.GetMessage())
}
