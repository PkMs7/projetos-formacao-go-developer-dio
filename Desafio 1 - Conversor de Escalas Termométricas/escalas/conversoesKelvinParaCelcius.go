package escalas

const ebulicaoAguaEmKelvin float64 = 373.2

func ConversaoKelvinParaCelcius(temperaturaKelvin float64) float64 {
	return temperaturaKelvin - 273.15
}
