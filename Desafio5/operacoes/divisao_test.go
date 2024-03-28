package operacoes_test

import (
	"calculadora/operacoes"
	"testing"
)

func ShouldDivCorrect(t *testing.T) {
	teste := operacoes.Dividir(40)
	rEsperado := 40
	if teste != rEsperado {
		t.Error("Valor esperado:", rEsperado, "Valor retornado:", teste)
	}
}

func ShouldDivIncorrect(t *testing.T) {
	teste := operacoes.Dividir(40)
	rEsperado := 50
	if teste != rEsperado {
		t.Error("Valor esperado:", rEsperado, "Valor retornado:", teste)
	}
}
