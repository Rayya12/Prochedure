package main

import "fmt"

func BesCil(a, b *int) {
	var perantara int
	if *a > *b {
		perantara = *a
		*a = *b
		*b = perantara
	}
}

func main() {
	var a, b int
	fmt.Scanln(&a, &b)
	for !(a == b) {
		BesCil(&a, &b)
		fmt.Println(a, b)
		fmt.Scanln(&a, &b)
	}
}
