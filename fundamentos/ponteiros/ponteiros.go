package main

import (
	"fmt"
)

func main() {

	i := 1

	var p *int = nil
	p = &i //pregando o endereço da variável
	*p++   // desreferenciando para pegar o valor

	i++

	// Go não tem aritmetica de ponteiros
	//p++

	fmt.Println(p, *p, i, &i)
}
