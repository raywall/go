package main

import "fmt"

func main() {
	defer fmt.Println("Executado no final")
	fmt.Println("Executado primeiro")
}
