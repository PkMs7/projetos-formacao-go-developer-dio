package operacoes_test

import (
	"calculadora/operacoes"
	"testing"
)

func ShouldSumCorrect(t *testing.T) {
	teste := operacoes.Somar(5, 7, 9)
	rEsperado := 21
	if teste != rEsperado {
		t.Error("Valor esperado:", rEsperado, "Valor retornado:", teste)
	}
}

func ShouldSumIncorrect(t *testing.T) {
	teste := operacoes.Somar(5, 7, 9)
	rEsperado := 73
	if teste != rEsperado {
		t.Error("Valor esperado:", rEsperado, "Valor retornado:", teste)
	}
}
