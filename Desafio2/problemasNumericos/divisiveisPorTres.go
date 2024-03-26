package problemas

func NumerosDivisiveisPorTres(numero int) []int {
	listaDeDivisiveisPorTres := []int{}

	for i := 1; i <= numero; i++ {
		if i%3 == 0 {
			listaDeDivisiveisPorTres = append(listaDeDivisiveisPorTres, i)
		}
	}
	return listaDeDivisiveisPorTres
}
