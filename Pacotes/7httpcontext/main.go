package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	ctx := context.Background()
	// context.WithTimeout() precisa passar o ctx, (linha 12) e o tempo máximo que a chamada http deve retornar
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	// // ctx, cancel := context.WithCancel(ctx, time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "http://google.com", nil)
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
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
