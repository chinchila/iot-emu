package device

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func RouteHome(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	dev := ctx.Value(SERVER_DEVICE_CONTEXT_KEY).(*Device)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, fmt.Sprintf("{\"name\":\"%s\",\"status\":\"%s\"}", dev.Name, dev.ServerStatus))
}

func RouteHomeThing(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	dev := ctx.Value(SERVER_DEVICE_CONTEXT_KEY).(*Device)
	w.Header().Set("Content-Type", "application/json")
	id := fmt.Sprintf(URN_FORMAT_STRING, strings.Split(r.URL.Path, string(os.PathSeparator))[1])
	fmt.Println(dev.Things)
	fmt.Println(dev.Things[id])
	io.WriteString(w, dev.Name)
}
