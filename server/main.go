package main

import (
	"log"
	"net/http"
)

// loggingHandler enhances a handler with access logging
type loggingHandler struct {
	handler http.Handler
}

func (h loggingHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Printf("path: %s header: %s", req.Header, req.URL.Path)
	h.handler.ServeHTTP(w, req)
}

func main() {
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", loggingHandler{fs})

	log.Printf("Serving at HTTP port 8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
