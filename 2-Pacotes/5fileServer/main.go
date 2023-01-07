package main

import (
	"log"
	"net/http"
)

func main() {

	fileServer := http.FileServer(http.Dir("./style"))
	mux := http.NewServeMux()
	mux.Handle("/", fileServer)
	mux.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from my Blog"))
	})
	log.Fatal(http.ListenAndServe(":9090", mux))
}
