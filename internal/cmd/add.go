package cmd

import (
	"fmt"

	"github.com/chinchila/iot-emu/internal/device"
)

const (
	ADD_LITERAL = "add"
	ADD_USAGE   = ADD_LITERAL + " [name] [address] [port]"
)

func Add(args []string) int {
	if len(args) < 1 {
		fmt.Println("You need to specify at least the name of your device.")
		return 1
	}
	address := "localhost"
	port := device.GetNewPort()
	name := args[0]
	if _, err := device.Add(name, address, port); err != nil {
		fmt.Printf("There was an error running: %v.\n", err)
		return 2
	}
	fmt.Printf("Added device with name %s, address %s:%d\n", name, address, port)
	return 0
}
