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

	// go func() {
	// 	slices := []int{}
	// 	for true {
	// 		fmt.Println("I am eating your memory!")
	// 		slices = append(slices, make([]int, 1000)...)
	// 		time.Sleep(time.Second)
	// 	}
	// }()

	// go func() {
	// 	fmt.Println("I am smashing your CPU!")
	// 	for true {
	// 		time.Sleep(time.Microsecond * 100)
	// 	}
	// }()

	log.Printf("Serving at HTTP port 8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
