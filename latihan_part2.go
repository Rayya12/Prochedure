package main

import "fmt"

func tempPara(c float64, r, f, k *float64) {
	*r = c * 4.00 / 5.00
	*f = c*9.00/5.00 + 32.00
	*k = c + 273.15
}

func main() {
	var cel, ream, fahren, kelv float64
	fmt.Scanln(&cel)
	tempPara(cel, &ream, &fahren, &kelv)
	fmt.Println(ream, fahren, kelv)
}
