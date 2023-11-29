package device

import (
	"io"
	"net/http"
)

func RouteHome(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	d := ctx.Value(SERVER_DEVICE_CONTEXT_KEY).(*Device)
	io.WriteString(w, d.FullAddress()+"!\n")
}
