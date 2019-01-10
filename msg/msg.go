package msg

import "github.com/google/wire"

// Message interface get message text.
type Message interface {
	GetMessage() string
}

// Hello is Hello message.
type Hello struct {
	Text string
}

// GetMessage is get Hello message text.
func (h *Hello) GetMessage() string {
	return "Hello " + h.Text + "."
}

// ProvideHello is create new instance of Hello struct.
func ProvideHello(text string) *Hello {
	return &Hello{Text: text}
}

// MessageHello provide wire of Hello struct.
var MessageHello = wire.NewSet(
	ProvideHello,
	wire.Bind(new(Message), new(Hello)))
