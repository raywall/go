package main

import "fmt"

func main() {
	// Declaração
	slice := []int{1, 2, 3}
	fmt.Println(slice) // [1 2 3]

	// Adicionar elementos
	slice = append(slice, 4)
	fmt.Println(slice) // [1 2 3 4]

	// Slice a partir de array
	array := [5]int{1, 2, 3, 4, 5}
	subSlice := array[1:4] // [2 3 4]
	fmt.Println(subSlice)
}
