//+build wireinject

package msg

import (
	"github.com/google/wire"
)

// InitializeMessage create instance of Message interface.
func InitializeMessage(text string) (Message, error) {
	wire.Build(MessageHello)
	return new(Hello), nil
}
