package device

import (
	"fmt"
	"math/rand"

	wothing "github.com/project-eria/go-wot/thing"
)

var deviceMap = map[string]*Device{}
var badPorts [70000]bool

func Add(name, address string, port int) (bool, error) {
	if len(deviceMap) >= MAX_DEVICES {
		return false, fmt.Errorf("maximum devices reached, please delete one before adding a new")
	}
	if _, ok := deviceMap[name]; ok {
		return false, fmt.Errorf("device with this name already exist")
	}
	deviceMap[name] = &Device{
		Name:         name,
		Address:      address,
		Port:         port,
		ServerStatus: SERVER_STATUS_CLOSED,
		Things:       make(map[string]*wothing.Thing),
	}
	badPorts[port] = true
	return true, nil
}

func Remove(name string) error {
	if _, ok := deviceMap[name]; !ok {
		return fmt.Errorf("device with name %s does not exist", name)
	}
	deviceMap[name].Stop()
	delete(deviceMap, name)
	return nil
}

func GetDeviceByName(name string) *Device {
	d, ok := deviceMap[name]
	if !ok {
		return nil
	}
	return d
}

func GetAll() map[string]*Device {
	return deviceMap
}

func GetNewPort() int {
	for {
		n := rand.Intn(20000-10000) + 10000
		if !badPorts[n] {
			return n
		}
	}
}
