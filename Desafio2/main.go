package main

import (
	"fmt"
	problemas "problemasNumericos/problemasNumericos"
)

func main() {
	var valor int = 100
	fmt.Printf("A lista de números divisiveis por 3 de %v é: %v", valor, problemas.NumerosDivisiveisPorTres(valor))
}
