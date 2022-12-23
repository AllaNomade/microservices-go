package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//os.Create cria um determinado arquivo
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}
	//1)Adiciona um texto dentro do arquivo determinado
	//NOTA: se não souber que é uma string, segue linha 17, se for string, segue linha 19
	tamanho, err := f.Write([]byte("Escrevendo dados no arquivo\n"))
	//Adiciona um texto dentro do arquivo determinado
	// // tamanho, err := f.WriteString("Hello,World!")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso, tamanho %d bytes\n", tamanho)
	f.Close()

	//Lê o os dados dentro do arquivo especificado.
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	//declaramos string na função pois a função somente retornará bytes caso contrário.
	fmt.Println(string(arquivo))

	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	//O pacote bufio com o reader possibilita a leitura do arquivo determinada a casa byte, como na linha 41
	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 3)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)

	}
}
