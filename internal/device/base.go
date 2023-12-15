package device

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"

	wothing "github.com/project-eria/go-wot/thing"
)

type contextKey string

const (
	SERVER_DEVICE_CONTEXT_KEY contextKey = "serverDeviceContextKey"
	SERVER_STATUS_LISTENING   string     = "open"
	SERVER_STATUS_CLOSED      string     = "closed"
	URN_FORMAT_STRING         string     = "urn:uuid:%s"
)

type Device struct {
	Name         string
	Port         int
	Address      string
	ServerStatus string
	Things       map[string]*wothing.Thing
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

func (d *Device) NewThing(title, version, description string) (*wothing.Thing, error) {
	id := fmt.Sprintf("uuid:%d", len(d.Things))
	thing, err := wothing.New(id, version, title, description, nil)
	if err != nil {
		return nil, err
	}
	d.Things[id] = thing
	mux := d.Server.Handler.(*http.ServeMux)
	routes := fmt.Sprintf("/%s/", thing.ID)
	mux.HandleFunc(routes, RouteHomeThing)
	return thing, nil
}

func (d *Device) GetThingById(id string) *wothing.Thing {
	t, ok := d.Things[id]
	if !ok {
		return nil
	}
	return t
}

// func (d *Device) Detach(id string) {
// 	id := fmt.Sprintf(URN_FORMAT_STRING, id)

// }

func (d *Device) Stop() {
	ctx := context.Background()
	if d.Server != nil {
		d.Server.Shutdown(ctx)
	}
}

func (d *Device) FullAddress() string {
	return fmt.Sprintf("%s:%d", d.Address, d.Port)
}
