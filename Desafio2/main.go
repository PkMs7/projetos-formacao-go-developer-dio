package main

import (
	"fmt"
	problemas "problemasNumericos/problemasNumericos"
)

func main() {
	// Desafio 1
	var valorParaEncontrarMultiplosDeTres int = 100
	fmt.Printf("A lista de números divisiveis por 3 de %v é: %v", valorParaEncontrarMultiplosDeTres, problemas.NumerosDivisiveisPorTres(valorParaEncontrarMultiplosDeTres))

	// Desafio 2
	var valorBrincadeiraPinPan int = 100
	fmt.Printf("As falas da brincadeira com o valor %v são %v", valorBrincadeiraPinPan, problemas.JogarPinPan(valorBrincadeiraPinPan))
}
