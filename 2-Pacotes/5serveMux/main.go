package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux1 := http.NewServeMux()
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello, World"))
	// })

	mux.HandleFunc("/", HomeHandler)
	mux.Handle("/blog", blog{title: "my-blog"})
	http.ListenAndServe(":9090", mux)

	mux1.HandleFunc("/", TestHandler)
	http.ListenAndServe(":7777", mux1)

	////////OBSERVAÇÃO: Os dois Mux não funcionam ao mesmo tempo
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("asdfsad"))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World"))
}

type blog struct {
	title string
}

//serve.http possui o ResponseWrite e o Request
func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}
