package operacoes

func Subtrair(i ...int) int {
	resultado := 0
	for _, valor := range i {
		resultado = valor - resultado
	}
	return resultado
}
