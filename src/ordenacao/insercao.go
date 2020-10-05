package ordenacao

// InsertionSort executa o algoritmo de ordenacao insertion sort.
func InsertionSort() {
	for i := 1; i < n; i++ {
		tmp := array[i]
		j := i - 1

		for (j >= 0) && (array[j] > tmp) {
			array[j+1] = array[j]
			j--
		}
		array[j+1] = tmp

	}
}
