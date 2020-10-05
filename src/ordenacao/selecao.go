package ordenacao

// SelectionSort executa o algoritmo de ordenacao selection sort.
func SelectionSort() {
	for i = 0; i < (n - 1); i++ {
		indice := i
		for j = (i + 1); j < n; j++ {
			if array[indice] > array[j] {
				indice = j
			}
		}
		swap(indice, i)
	}
}
