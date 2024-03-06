package main

import "fmt"

func zoomIn(w, h, n int, wni, hni *int) {
	*wni = w * n
	*hni = h * n
}

func zoomOut(w, h, n int, wno, hno *int) {
	*wno = w / n
	*hno = h / n
}

func zoom(w, h, n int, s string, wn, hn *int) {
	var wni, hni int
	var wno, hno int
	if s == "+" {
		zoomIn(w, h, n, &wni, &hni)
		*wn = wni
		*hn = hni
	}
	if s == "-" {
		zoomOut(w, h, n, &wno, &hno)
		*wn = wno
		*hn = hno
	}
}

func main() {
	var w, h, n, wn, hn int
	var s string
	fmt.Scanln(&w, &h)
	fmt.Scanln(&s, &n)
	zoom(w, h, n, s, &wn, &hn)
	fmt.Println(wn, hn)
}
