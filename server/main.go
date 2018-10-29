package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// loggingHandler enhances a handler with access logging
type loggingHandler struct {
	handler http.Handler
}

func (h loggingHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Printf("path: %s header: %s", req.Header, req.URL.Path)
	h.handler.ServeHTTP(w, req)
}

func selectModeFromEnv(w http.ResponseWriter, req *http.Request) {
	mode := os.Getenv("SH_DEMO_APP_MODE")
	fmt.Fprintf(w, mode)
}

func main() {
	fs := http.FileServer(http.Dir("."))
	http.HandleFunc("/api/mode", selectModeFromEnv)
	http.Handle("/", loggingHandler{fs})

	log.Printf("Serving at HTTP port 1216")
	log.Fatal(http.ListenAndServe("0.0.0.0:1216", nil))
}
