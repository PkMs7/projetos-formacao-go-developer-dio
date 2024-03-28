package main

import (
	"calculadora/operacoes"
	"fmt"
)

func main() {

	divisao := operacoes.Dividir(40)
	multiplicacao := operacoes.Multiplicar(5, 3)
	soma := operacoes.Somar(5, 7, 9)
	subtracao := operacoes.Subtrair(3, 5)

	fmt.Println(divisao, multiplicacao, soma, subtracao)
}
