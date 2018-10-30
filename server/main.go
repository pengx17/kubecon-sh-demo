package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// loggingHandler enhances a handler with access logging
type loggingHandler struct {
	handler http.Handler
}

func (h loggingHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Printf("path: %s header: %s", req.Header, req.URL.Path)
	h.handler.ServeHTTP(w, req)
}

var tooMuchMemory bool
var mode string

func selectModeFromEnv(w http.ResponseWriter, req *http.Request) {
	if tooMuchMemory {
		mode = "ERROR"
	}
	fmt.Fprintf(w, mode)
}

func main() {
	fs := http.FileServer(http.Dir("."))
	mode = os.Getenv("APP_MODE")
	http.HandleFunc("/api/mode", selectModeFromEnv)
	http.Handle("/", loggingHandler{fs})

	if mode == "NEW" {
		go func() {
			time.Sleep(time.Second * 30)
			slices := []int{}
			tooMuchMemory = true
			fmt.Println("I am smashing your memory and CPU!")
			slices = append(slices, make([]int, 10000000)...)
			for true {
				slices = append(slices, make([]int, 10)...)
				time.Sleep(time.Microsecond * 100)
			}
		}()
	}

	log.Printf("Serving at HTTP port 8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
