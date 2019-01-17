//+build wireinject

package msg

import (
	"github.com/google/wire"
)

// InitializeHello create instance of Message interface.
func InitializeHello(text string) (*Hello, error) {
	panic(wire.Build(ProvideHello))
}
