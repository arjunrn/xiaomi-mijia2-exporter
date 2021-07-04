package app

import (
    "fmt"

	"github.com/muka/go-bluetooth/bluez/profile/adapter"
)

func App() {
	defaultAdapter := adapter.GetDefaultAdapterID()
	fmt.Printf("Default Adapter: %s\n", defaultAdapter)
}
