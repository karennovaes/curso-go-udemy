# Tipos


Em Go, assim como em muitas outras linguagens de programação, existem vários tipos de dados para representar diferentes tipos de informações. Alguns dos tipos de dados fundamentais em Go incluem:

* **Tipos Numéricos**:

    - **int:** Números inteiros, cujo tamanho depende da arquitetura do sistema (32 ou 64 bits).
    - **int8, int16, int32, int64:** Inteiros com tamanhos específicos em bits (com sinal).
    - **uint:** Números inteiros sem sinal, com tamanho dependente da arquitetura.
    - **uint8 (byte), uint16, uint32, uint64:** Inteiros sem sinal com tamanhos específicos em bits.
    - **float32, float64:** Números de ponto flutuante com precisão simples e dupla, respectivamente.
    - **complex64, complex128:** Números complexos com partes reais e imaginárias representadas por float32 e float64.

* **Tipos de Texto**:

    - **string:** Sequência de caracteres Unicode.

* **Tipos Booleanos**:

    - **bool:** Representa valores booleanos, ou seja, true ou false.

* **Tipos de Dados Compostos:**

    - **arrays:** Coleção de elementos de tamanho fixo e do mesmo tipo.
    - **slices:** Segmento flexível de um array.
    - **maps:** Coleção não ordenada de pares chave-valor.
    - **structs:** Tipo de dados que permite agrupar diferentes tipos de dados com chaves nomeadas.

* **Tipos Especiais:**

    - **nil:** Valor zero de vários tipos de dados, usado para representar a ausência de valor em ponteiros, slices, maps, canais e interfaces.

* **Outros Tipos:**

    - **byte:** Um alias para uint8.
    - **rune:** Um alias para int32, usado para representar um caractere Unicode.