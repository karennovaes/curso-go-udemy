package main

import "fmt"

func notaParaConteito(n float64) string {
	var nota = int(n)

	switch nota {
	case 10:

		fallthrough

	case 9:

		return "A"

	case 8, 7:

		return "B"

	case 6, 5:

		return "C"

	case 4, 3:
		return "D"

	case 2, 1, 0:

		return "E"

	default:

		return "Nota inválida"

	}
}

func main() {
	fmt.Println(notaParaConteito(9.8))
	fmt.Println(notaParaConteito(6.9))
	fmt.Println(notaParaConteito(2.1))
}
