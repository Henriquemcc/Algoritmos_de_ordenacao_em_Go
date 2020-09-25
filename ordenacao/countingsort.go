//Retorna o maior elemento do arranjo.
package ordenacao

func getMaior() int {
	maior := array[0]

	for i = 0; i < n; i++ {
		if maior < array[i] {
			maior = array[i]
		}
	}
	return maior
}

// Algoritmo de ordenacao
func CountingSort() {

	//Array para contar o numero de ocorrencias de cada elemento
	tamCount := getMaior() + 1
	var count = gerarArranjo(n)
	var ordenado = gerarArranjo(n)

	//Inicializar cada posicao do array de contagem
	for i = 0; i < tamCount; i++ {
		count = append(count, 0)

	}

	//Agora, o count[i] contem o numero de elemento iguais a i
	for i = 0; i < n; i++ {
		count[array[i]]++
	}

	//Agora, o count[i] contem o numero de elemento menores ou iguais a i
	for i = 1; i < tamCount; i++ {
		count[i] += count[i-1]
	}

	//Ordenando
	for i = n - 1; i >= 0; i-- {

		ordenado[count[array[i]]-1] = array[i]
		count[array[i]]--
	}

	//Copiando para o array original
	for i = 0; i < n; i++ {
		array[i] = ordenado[i]
	}
}
