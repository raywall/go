package main

import "testing"

func soma(a, b int) int {
	return a + b
}

func TestSoma(t *testing.T) {
	testes := []struct {
		nome     string
		a, b     int
		esperado int
	}{
		{"positivo", 2, 3, 5},
		{"negativo", -1, -1, -2},
		{"zero", 0, 5, 5},
	}

	for _, tt := range testes {
		t.Run(tt.nome, func(t *testing.T) {
			resultado := soma(tt.a, tt.b)
			if resultado != tt.esperado {
				t.Errorf("soma(%d, %d) = %d; esperado %d", tt.a, tt.b, resultado, tt.esperado)
			}
		})
	}
}
