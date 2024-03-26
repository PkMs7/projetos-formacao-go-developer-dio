package escalas

const kelvinAZeroCelcius float64 = 273.15

func ConversaoKelvinParaCelcius(temperaturaKelvin float64) float64 {
	return temperaturaKelvin - kelvinAZeroCelcius
}
