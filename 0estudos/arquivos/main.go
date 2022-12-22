package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	a, err := os.Create("arquivo1.txt")
	if err != nil {
		panic(err)
	}

	tamanho, err := a.Write([]byte("Escrevendo dados no arquivo"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso, tamanho %d bytes\n", tamanho)
	a.Close()

	arquivo, err := os.ReadFile("arquivo1.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arquivo))

	arquivo1, err := os.Open("arquivo1.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(arquivo1)
	buffer := make([]byte, 3)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

}
