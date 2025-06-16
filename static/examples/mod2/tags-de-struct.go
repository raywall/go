package main

import (
	"encoding/json"
	"fmt"
)

type Produto struct {
	ID    int     `json:"id"`
	Nome  string  `json:"nome"`
	Preco float64 `json:"preco,omitempty"`
}

func main() {
	p := Produto{ID: 1, Nome: "Laptop"}
	jsonData, _ := json.Marshal(p)
	fmt.Println(string(jsonData)) // {"id":1,"nome":"Laptop"}
}
