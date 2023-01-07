package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	c := http.Client{}
	//O NewBuffer recebe um slice de bytes, e retorna um buffer
	/// Eo buffer é um espaço com informações que podemos "entregar" para um Reader
	jsonVar := bytes.NewBuffer([]byte(`{"name": "alan"}`))
	resp, err := c.Post("http://google.com", "application/json", jsonVar)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//O CopyBuffer pega os dados, da onde vou pegar estes dados e escolhe onde jogar
	///No caso, pego do resp.Body e "jogá-los" no os.Stdout
	io.CopyBuffer(os.Stdout, resp.Body, nil)
}
