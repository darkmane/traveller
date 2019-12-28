package handlers

import (
	"fmt"
	"net/http"
)

func RegisterHandlers(httpFunc func(pattern string, handler func(http.ResponseWriter, *http.Request))) {
	httpFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
