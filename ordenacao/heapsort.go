package ordenacao

func constroi(tamHeap int) {
	for i := tamHeap; i > 1 && array[i] > array[i/2]; i /= 2 {
		swap(i, i/2)
	}
}

func reconstroi(tamHeap int) {
	i := 1
	var filho int
	for i <= (tamHeap / 2) {

		if array[2*i] > array[2*i+1] || 2*i == tamHeap {
			filho = 2 * i
		} else {
			filho = 2*i + 1
		}

		if array[i] < array[filho] {
			swap(i, filho)
			i = filho
		} else {
			i = tamHeap
		}
	}
}

// Algoritmo de ordenacao
func HeapSort() {
	//Alterar o vetor ignorando a posicao zero
	for i := n; i > 0; i-- {
		array[i] = array[i-1]
	}

	//Contrucao do heap
	for tamHeap := 1; tamHeap <= n; tamHeap++ {
		constroi(tamHeap)
	}

	//Ordenacao propriamente dita
	tamHeap := n
	for tamHeap > 1 {
		swap(1, tamHeap)
		tamHeap--
		reconstroi(tamHeap)
	}

	//Alterar o vetor para voltar a posicao zero
	for i := 0; i < n; i++ {
		array[i] = array[i+1]
	}
}
