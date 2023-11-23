package cmd

import (
	"fmt"

	"github.com/chinchila/iot-emu/internal/device"
)

const (
	START_LITERAL = "start"
	START_USAGE   = START_LITERAL + " [name]"
)

func Start(args []string) int {
	if len(args) < 1 {
		fmt.Println("You need to specify at least the name of your device.")
		return 1
	}
	name := args[0]
	dev := device.GetDeviceByName(name)
	if dev == nil {
		fmt.Printf("Could not find device with name %s.\n", name)
		return 2
	}
	dev.Start()
	fmt.Printf("Started server %s with address %s\n", dev.Name, dev.FullAddress())
	return 0
}
