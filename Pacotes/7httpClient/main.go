package main

import (
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	//Timeout determina o tempo máximo de retorno http, caso ultrapasse o tempo, vai "estourar" o erro.
	c := http.Client{Timeout: time.Duration(1) * time.Microsecond}
	resp, err := c.Get("http://google.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}
