package operacoes

func Somar(i ...int) int {
	var resultado int
	for _, valor := range i {
		resultado += valor
	}
	return resultado
}
