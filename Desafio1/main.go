package main

import (
	escalas "escalasTermometricas/EscalasTermometricas"
	"fmt"
)

func main() {
	var valorEmKelvinParaConversaoEmCelcius float64 = 373.2

	fmt.Printf("O valor convertido para Kelvin é: %.2f C°", escalas.ConversaoKelvinParaCelcius(valorEmKelvinParaConversaoEmCelcius))

}
