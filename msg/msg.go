package msg

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
