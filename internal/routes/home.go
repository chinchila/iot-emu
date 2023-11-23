package routes

import (
	"io"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	// fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}
