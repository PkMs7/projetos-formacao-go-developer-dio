package operacoes_test

import (
	"calculadora/operacoes"
	"testing"
)

func ShouldMultCorrect(t *testing.T) {
	teste := operacoes.Multiplicar(5, 3)
	rEsperado := 15
	if teste != rEsperado {
		t.Error("Valor esperado:", rEsperado, "Valor retornado:", teste)
	}
}
func ShouldMultIncorrect(t *testing.T) {
	teste := operacoes.Multiplicar(5, 3)
	rEsperado := 27
	if teste != rEsperado {
		t.Error("Valor esperado:", rEsperado, "Valor retornado:", teste)
	}
}
