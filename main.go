package main

import (
	"fmt"

	"gitlab.com/purini-to/stack_code_cov/msg"

	"go.uber.org/zap"
	"rsc.io/quote"
)

func main() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", "https://localhost",
		"attempt", 3,
		"backoff", 123,
	)
	fmt.Println(quote.Hello() + msg.GetMessage())
}
