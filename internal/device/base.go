package device

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/chinchila/iot-emu/internal/routes"
)

const SERVER_ADDRESS_KEY = "server_address_key"

type Device struct {
	Name    string
	Port    int
	Address string
	Server  *http.Server
}

func (d *Device) Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", routes.Home)
	// mux.HandleFunc("/hello", getHello)
	d.Server = &http.Server{
		Addr:    d.FullAddress(),
		Handler: mux,
	}
	go func() {
		err := d.Server.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server one closed\n")
		} else if err != nil {
			fmt.Printf("error listening for server one: %s\n", err)
		}
	}()
	fmt.Println(d.Server)
}

func (d *Device) Stop() {
	// ctx := context.Background()
	fmt.Println(d.Server)
	if d.Server != nil {
		d.Server.Shutdown(context.TODO())
	}
}

func (d *Device) FullAddress() string {
	return fmt.Sprintf("%s:%d", d.Address, d.Port)
}
