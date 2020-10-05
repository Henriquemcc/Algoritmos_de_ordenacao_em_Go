package ordenacao

// insercaoPorCor insere elementos do arranjo de acordo com a cor da parte do arranjo.
// Parametro cor: Numero inteiro indicando a cor da parte do arranjo.
// Parametro h: Quantidade de partes do arranjo.
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

// ShellSort executa o algoritmo de ordenacao shellsort.
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
