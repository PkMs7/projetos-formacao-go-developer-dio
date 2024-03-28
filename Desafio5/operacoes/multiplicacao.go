package operacoes

func Multiplicar(i ...int) int {
	resultado := 1
	for _, valor := range i {
		resultado = resultado * valor
	}
	return resultado
}
