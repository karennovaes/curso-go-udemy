#Condicionais

Em Go, os condicionais são expressos através das estruturas de controle de fluxo **if**, **else if**, **else**. Aqui está um exemplo que demonstra como usar condicionais em Go:


```
	// Exemplo de estrutura condicional usando if-else
	if idade < 18 {
		fmt.Println("Você é menor de idade")
	} else if idade == 18 {
		fmt.Println("Você acabou de atingir a maioridade")
	} else {
		fmt.Println("Você é maior de idade")
	}

    // Switch case em Go
	switch idade {
	case 10:
		fmt.Println("Você tem 10 anos")
	case 18:
		fmt.Println("Você tem 18 anos")
	default:
		fmt.Println("Sua idade não é 10 nem 18")
	}

```

* A estrutura **if** é usada para testar uma condição. Se a condição for verdadeira, o bloco de código associado ao if é executado. Caso contrário, se houver uma cláusula **else if** ou **else**, a condição associada será testada até que uma condição seja verdadeira ou até chegar ao else (caso nenhuma condição seja verdadeira).

* A estrutura **switch** é usada para avaliar várias condições diferentes de maneira mais concisa do que várias declarações **if-else**. Ela compara a expressão especificada no switch com os casos listados e executa o bloco de código correspondente ao primeiro caso que corresponde à expressão.



    > Lembre-se de que em Go não é necessário utilizar parênteses ao redor da expressão condicional em **if** ou **switch**, mas os blocos de código devem estar delimitados por chaves **{ }**. Além disso, a sintaxe **switch** permite usar expressões de tipo além de valores, o que oferece mais flexibilidade em comparações.