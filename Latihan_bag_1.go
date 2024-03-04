package main

import "fmt"

func pecahUang(uang int, k10, k2, k1 *int) {
	*k10 = uang / 10000
	uang = uang % 10000
	*k2 = uang / 2000
	uang = uang % 2000
	*k1 = uang / 1000

}

func main() {
	var n, l10, l2, l1 int
	fmt.Scan(&n)

	pecahUang(n, &l10, &l2, &l1)
	fmt.Println(l10, l2, l1)

}
