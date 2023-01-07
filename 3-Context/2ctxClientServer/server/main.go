package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("request iniciada")
	defer log.Println("request finalizada")
	select {
	case <-time.After(5 * time.Second):
		//imprime no command line: os.stdout
		log.Println("Request processada com sucesso")
		//imprime no Browser
		w.Write([]byte("Request processada com sucesso!"))
	case <-ctx.Done():
		//imprime no command line: os.stdout
		log.Println("Request cancelada pelo cliente")

	}
}
