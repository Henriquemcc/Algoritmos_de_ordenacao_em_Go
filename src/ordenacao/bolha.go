package ordenacao

// BubbleSort executa o algoritmo de ordenacao bolha no arranjo de numeros.
func BubbleSort() {
	for i = (n - 1); i > 0; i-- {
		for j = 0; j < i; j++ {
			if array[j] > array[j+1] {
				auxiliar := array[j+1]
				array[j+1] = array[j]
				array[j] = auxiliar
			}
		}
	}
}
