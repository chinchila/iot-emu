package cmd

import (
	"fmt"

	"github.com/chinchila/iot-emu/internal/device"
)

const (
	INFO_LITERAL = "info"
	INFO_USAGE   = INFO_LITERAL + " [name: optional]"
)

func printDevInfo(dev *device.Device) {
	fmt.Printf("Device Name: %s\n"+
		"Device Address: %s\n",
		dev.Name, dev.FullAddress(),
	)
}

func Info(args []string) int {
	if len(args) >= 1 {
		for _, name := range args {
			dev := device.GetDeviceByName(name)
			fmt.Println("===")
			if dev == nil {
				fmt.Printf("Device with name %s not found.\n", name)
			} else {
				printDevInfo(dev)
			}
		}
	} else {
		allDevices := device.GetAll()
		for _, v := range allDevices {
			fmt.Println("===")
			printDevInfo(v)
		}
	}
	return 0
}
