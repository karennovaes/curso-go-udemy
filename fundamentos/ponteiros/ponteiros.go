package main

import (
	"fmt"
)

func main() {

	i := 1

	// Go não tem aritmetica de ponteiros

	var p *int = nil
	p = &i //pregando o endereçi da variável
	*p++   // desreferenciando para pegar o valor
	i++
	fmt.Println(p, *p, i, &i)
}
