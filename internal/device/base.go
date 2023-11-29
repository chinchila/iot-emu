package device

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
)

type contextKey string

const (
	SERVER_DEVICE_CONTEXT_KEY contextKey = "serverDeviceContextKey"
	SERVER_STATUS_LISTENING   string     = "open"
	SERVER_STATUS_CLOSED      string     = "closed"
)

type Device struct {
	Name         string
	Port         int
	Address      string
	ServerStatus string
	Server       *http.Server
}

func (d *Device) Start() {
	ctx := context.Background()
	mux := http.NewServeMux()
	mux.HandleFunc("/", RouteHome)
	d.Server = &http.Server{
		Addr:    d.FullAddress(),
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, SERVER_DEVICE_CONTEXT_KEY, d)
			return ctx
		},
	}
	go func() {
		d.ServerStatus = SERVER_STATUS_LISTENING
		err := d.Server.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			d.ServerStatus = SERVER_STATUS_CLOSED
			fmt.Printf("Stoped server %s with address %s\n", d.Name, d.FullAddress())
		} else if err != nil {
			fmt.Printf("error listening on server %s: %s\n", d.Name, err)
		}
	}()
}

func (d *Device) Stop() {
	ctx := context.Background()
	if d.Server != nil {
		d.Server.Shutdown(ctx)
	}
}

func (d *Device) FullAddress() string {
	return fmt.Sprintf("%s:%d", d.Address, d.Port)
}
