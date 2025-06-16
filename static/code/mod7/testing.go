package main

import "testing"

func soma(a, b int) int {
	return a + b
}

func TestSoma(t *testing.T) {
	resultado := soma(2, 3)
	if resultado != 5 {
		t.Errorf("soma(2, 3) = %d; esperado 5", resultado)
	}
}
