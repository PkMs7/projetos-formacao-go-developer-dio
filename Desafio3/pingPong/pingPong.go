package pingpong

import (
	"fmt"
	"time"
)

func ChamarPing(canal chan string) {
	for i := 0; ; i++ {
		canal <- "Ping"
	}
}

func ChamarPong(canal chan string) {
	for i := 0; ; i++ {
		canal <- "Pong"
	}
}

func ImprimirPinPong(canal chan string) {
	for {
		mensagem := <-canal
		fmt.Println(mensagem)
		time.Sleep(time.Second * 1)
	}
}
