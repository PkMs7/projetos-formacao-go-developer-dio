package main

import (
	escalas "escalasTermometricas/EscalasTermometricas"
	"fmt"
)

func main() {
	var valorEmKelvin float64 = 373.2

	fmt.Printf("O valor convertido para Kelvin é: %.2f C°", escalas.ConversaoKelvinParaCelcius(valorEmKelvin))

}
