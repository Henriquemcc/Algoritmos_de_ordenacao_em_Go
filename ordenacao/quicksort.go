package ordenacao

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

// Algoritmo de ordenacao
func QuickSort() {
	quicksortRec(0, n-1)
}
