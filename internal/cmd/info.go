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
	fmt.Printf("Device Name: %s\nDevice Address: %s\nDevice Port: %d\n", dev.Name, dev.Address, dev.Port)
}

func Info(args []string) int {
	if len(args) >= 1 {
		for _, name := range args {
			dev := device.GetDeviceByName(name)
			if dev == nil {
				fmt.Printf("Device with name %s not found.\n", name)
			} else {
				printDevInfo(dev)
			}
			fmt.Println("===")
		}
	} else {
		allDevices := device.GetAll()
		for _, v := range allDevices {
			printDevInfo(&v)
			fmt.Println("===")
		}
	}
	return 0
}
