package api

import (
	"fmt"
	"net/http"
)

func StartServer(port string, h http.Handler) {
	if port[0] != ':' {
		port = ":" + port
	}
	fmt.Printf("Server connected with port no %s", port)
	http.ListenAndServe(port, h)
}

