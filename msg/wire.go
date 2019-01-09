//+build wireinject

package msg

import (
	"github.com/google/wire"
)

// MessageHello provide wire of Hello struct.
var MessageHello = wire.NewSet(
	ProvideHello,
	wire.Bind(new(Message), new(Hello)))

// InitializeMessage create instance of Message interface.
func InitializeMessage(text string) (Message, error) {
	wire.Build(MessageHello)
	return new(Hello), nil
}
