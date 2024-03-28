package operacoes

func Dividir(i ...int) int {
	resultado := 1
	for _, valor := range i {
		resultado = valor / resultado
	}
	return resultado
}
