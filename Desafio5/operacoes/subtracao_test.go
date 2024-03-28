package operacoes_test

import (
	"calculadora/operacoes"
	"testing"
)

func ShouldSubCorrect(t *testing.T) {
	teste := operacoes.Subtrair(3, 5)
	rEsperado := 2
	if teste != rEsperado {
		t.Error("Valor esperado:", rEsperado, "Valor retornado:", teste)
	}
}

func ShouldDSubIncorrect(t *testing.T) {
	teste := operacoes.Subtrair(3, 5)
	rEsperado := 9
	if teste != rEsperado {
		t.Error("Valor esperado:", rEsperado, "Valor retornado:", teste)
	}
}
