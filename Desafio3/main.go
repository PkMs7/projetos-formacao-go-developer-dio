package main

import (
	"fmt"
	pingpong "pingPong/pingPong"
)

func main() {
	var canal chan string = make(chan string)

	go pingpong.ChamarPing(canal)
	go pingpong.ImprimirPinPong(canal)
	go pingpong.ChamarPong(canal)

	var parar string
	fmt.Scanln(&parar)

}
