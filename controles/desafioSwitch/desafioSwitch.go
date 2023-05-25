package main

import "fmt"

func notaParaConteito(n float64) string {

	switch {
	case n >= 9 && n <= 10:
		return "A"
	case n >= 7 && n < 9:
		return "B"
	case n >= 5 && n < 8:
		return "C"
	case n >= 3 && n < 5:
		return "D"
	default:
		return "E"
	}

}

func main() {
	fmt.Println(notaParaConteito(9.8))
	fmt.Println(notaParaConteito(6.9))
	fmt.Println(notaParaConteito(2.1))
}
