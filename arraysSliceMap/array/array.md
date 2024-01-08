# Array

Em Go, um array é uma estrutura de dados que armazena uma coleção fixa de elementos do mesmo tipo. A sintaxe para declarar um array em Go é:


```var nomeArray [tamanho]tipo```

* **nomeArray:** É o nome do array que você está declarando.
* **tamanho:** É o número de elementos que o array pode armazenar. O tamanho do array deve ser um valor constante e não pode ser alterado após a declaração.
* **tipo:** É o tipo de dado dos elementos que serão armazenados no array.


Aqui está um exemplo simples de como declarar e usar um array em Go:

```
package main

import "fmt"

func main() {
	// Declarando um array de inteiros com tamanho 5
	var numeros [5]int

	// Atribuindo valores ao array
	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30
	numeros[3] = 40
	numeros[4] = 50

	// Acessando e imprimindo os valores do array
	fmt.Println("Elementos do array:", numeros)
	fmt.Println("Primeiro elemento:", numeros[0])
	fmt.Println("Tamanho do array:", len(numeros))
}
```


> É importante observar que os arrays em Go são estruturas estáticas e têm tamanho fixo, o que significa que o número de elementos é determinado durante a compilação e não pode ser alterado durante a execução do programa. Se você precisa de uma coleção dinâmica de elementos, geralmente é preferível usar slices, que são uma abstração mais poderosa sobre arrays em Go.