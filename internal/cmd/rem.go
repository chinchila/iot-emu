package cmd

import (
	"fmt"

	"github.com/chinchila/iot-emu/internal/device"
)

const (
	REM_LITERAL    = "rem"
	REMOVE_LITERAL = "remove"
	REM_USAGE      = REM_LITERAL + " [name]"
)

func Rem(args []string) int {
	if len(args) < 1 {
		fmt.Println("You need to specify at least the name of your device.")
		return 1
	}
	name := args[0]
	if err := device.Remove(name); err != nil {
		fmt.Printf("There was an error running: %v.\n", err)
		return 2
	}
	fmt.Printf("Removed device with name %s\n", name)
	return 0
}
