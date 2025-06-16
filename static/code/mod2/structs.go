package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

func main() {
	p := Pessoa{Nome: "Raywall", Idade: 45}
	fmt.Println(p) // {Raywall 45}

	// Acesso a campos
	fmt.Println(p.Nome) // Raywall

	// Struct anônima
	anon := struct {
		Valor string
	}{Valor: "Teste"}

	fmt.Println(anon) // {Teste}
}
