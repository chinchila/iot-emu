package cmd

import (
	"fmt"

	"github.com/chinchila/iot-emu/internal/device"
)

const (
	TADD_LITERAL      = "tadd"
	THING_ADD_LITERAL = "thing_add"
	TADD_USAGE        = TADD_LITERAL + " [device name] [thing title] [thing version:optional] [thing description:optional]"
)

func ThingAdd(args []string) int {
	if len(args) < 2 {
		fmt.Println("You need to specify the device name and the thing title at least.")
		return 1
	}
	version := "1.0"
	if len(args) >= 3 {
		version = args[2]
	}
	description := "Sample description"
	if len(args) >= 4 {
		description = args[3]
	}
	device_name := args[0]
	title := args[1]
	dev := device.GetDeviceByName(device_name)
	if dev == nil {
		fmt.Printf("Could not find device with name %s.\n", device_name)
		return 2
	}
	thing, err := dev.NewThing(title, version, description)
	if err != nil {
		fmt.Printf("There was an error: %v.\n", err)
		return 2
	}
	fmt.Printf("Attached thing %s (%s) (%p) on device %s, address %s/%s/\n", thing.ID, thing.Title, thing, dev.Name, dev.FullAddress(), thing.ID)
	return 0
}
