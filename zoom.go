package main

import "fmt"

func Zoomin(w, h int, n int, wn, hn *int) {
	*wn = w * n
	*hn = h * n
}

func Zoomout(w, h int, n int, wn, hn *int) {
	*wn = w / n
	*hn = h / n
}

func main() {
	var w, h, n, wn, hn int
	var s string

	fmt.Scan(&w, &h)
	fmt.Scan(&s, &n)

	fmt.Println(s, n)

	if s == "-" {
		Zoomout(w, h, n, &wn, &hn)
	}
	if s == "+" {
		Zoomin(w, h, n, &wn, &hn)
	}

}
