package ordenacao

func insercaoPorCor(cor int, h int) {
	for i := (h + cor); i < n; i += h {
		tmp := array[i]
		j := i - h
		for (j >= 0) && (array[j] > tmp) {
			array[j+h] = array[j]
			j -= h
		}
		array[j+h] = tmp
	}
}

// Algoritmo de ordenacao
func ShellSort() {
	h := 1

	for {
		h = (h * 3) + 1

		if h >= n {
			break
		}
	}

	for {
		h /= 3
		for cor := 0; cor < h; cor++ {
			insercaoPorCor(cor, h)
		}

		if h == 1 {
			break
		}
	}
}
