package problemas

func JogarPinPan(numero int) []string {
	falasPinPan := []string{}

	for i := 1; i <= numero; i++ {
		if i%3 == 0 && i%5 == 0 {
			falasPinPan = append(falasPinPan, "PinPan")
		} else if i%5 == 0 {
			falasPinPan = append(falasPinPan, "Pan")
		} else if i%3 == 0 {
			falasPinPan = append(falasPinPan, "Pin")
		}
	}
	return falasPinPan
}
