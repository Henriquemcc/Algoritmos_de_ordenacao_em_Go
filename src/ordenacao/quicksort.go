package ordenacao

// quicksortRec executa o algoritmo de ordenacao quicksort de forma recursiva, passando parametro dos delimitadores esquerdo e direito.
// Parametro esq: Numero inteiro indicando o delimitador esquerdo.
// Parametro dir: Numero inteiro indicando o delimitador direito.
func quicksortRec(esq int, dir int) {
	i := esq
	j := dir
	pivo := array[(dir+esq)/2]
	for i <= j {
		for array[i] < pivo {
			i++
		}
		for array[j] > pivo {
			j--
		}
		if i <= j {
			swap(i, j)
			i++
			j--
		}
	}
	if esq < j {
		quicksortRec(esq, j)
	}
	if i < dir {
		quicksortRec(i, dir)
	}
}

// QuickSort executa o algoritmo de ordenacao quicksort.
func QuickSort() {
	quicksortRec(0, n-1)
}
