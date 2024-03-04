
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
	var m, n, o, p, q, r, s, t int
	fmt.Scanln(&m, &n)
	fmt.Scanln(&o, &p)
	fmt.Scanln(&q, &r)
	fmt.Scanln(&s, &t)

	BesCil(&m, &n)
	BesCil(&o, &p)
	BesCil(&q, &r)
	BesCil(&s, &t)
	fmt.Println(m, n)
	fmt.Println(o, p)
	fmt.Println(q, r)
	fmt.Println(s, t)
}
