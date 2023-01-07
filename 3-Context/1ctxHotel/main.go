package main

import (
	"context"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	//o statement select guarda resultados quando o resultado é chamado.
	select {
	//Caso passe de 3 segundos, a requisição é cancelada
	case <-ctx.Done():
		println("Hotel booking cancelled, timeout reached")
		return
		//Caso não seja cancelado em 3 segundos, como definido na linha 10.
		//A segunte função é executada
	case <-time.After(5 * time.Second):
		println("Hotel booked")
	}
}
