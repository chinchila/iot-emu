package cmd

import (
	"fmt"

	"github.com/chinchila/iot-emu/internal/device"
)

const (
	STOP_LITERAL = "stop"
	STOP_USAGE   = STOP_LITERAL + " [name]"
)

func Stop(args []string) int {
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
	dev.Stop()
	return 0
}
