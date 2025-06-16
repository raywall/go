package main

import "fmt"

func main() {
	numero := 45
	if x := numero % 2; x == 0 {
		fmt.Println("Par")
	} else {
		fmt.Println("Ímpar")
	}
}
