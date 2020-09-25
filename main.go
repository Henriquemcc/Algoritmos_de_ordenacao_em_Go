package main

import (
	"fmt"
	"time"
	"trabalho/fronteira"
	"trabalho/ordenacao"
)

func main() {

	for {
		var algoritmo = obterAlgoritmo()

		if algoritmo != 0 {
			ordenacao.SetTamanho(obterTamanhoArranjo())
			var ordem = obterOrdemArranjo()
			executarOrdenacao(ordem, algoritmo)
		} else {
			break
		}

	}

}

func obterTamanhoArranjo() int {
	fmt.Println("Qual o tamanho do arranjo?")
	return int(fronteira.ReadInt64())

}

func obterAlgoritmo() uint64 {
	var algoritmo uint64

	for {
		fmt.Println("Qual algoritmo de ordenação deseja executar?")
		fmt.Println("0 - Sair")
		fmt.Println("1 - Bolha")
		fmt.Println("2 - Counting Sort")
		fmt.Println("3 - Heap Sort")
		fmt.Println("4 - Insertion Sort")
		fmt.Println("5 - Quick Sort")
		fmt.Println("6 - Selection Sort")
		fmt.Println("7 - Shell Sort")

		algoritmo = fronteira.ReadUint64()

		if algoritmo <= 7 {
			break
		}
	}

	return algoritmo
}

func obterOrdemArranjo() uint64 {
	var ordem uint64

	for {
		fmt.Println("Como deseja que o arranjo de numeros esteja?")
		fmt.Println("0 - Aleatorio")
		fmt.Println("1 - Decrescente")

		ordem = fronteira.ReadUint64()

		if ordem <= 1 {
			break
		}
	}

	return ordem
}

func executarOrdenacao(ordem uint64, algoritmo uint64) {

	//Ordenando o arranjo de acordo com a ordem selecionada
	if ordem == 0 {
		ordenacao.Aleatorio()
	} else if ordem == 1 {
		ordenacao.Decrescente()
	}

	ordenacao.Mostrar()

	var tempoGasto time.Duration

	//Executando o algoritmo de ordenacao
	if algoritmo == 1 {
		comeco := time.Now()

		ordenacao.BubbleSort()

		fim := time.Now()
		tempoGasto = fim.Sub(comeco)

	} else if algoritmo == 2 {
		comeco := time.Now()

		ordenacao.CountingSort()

		fim := time.Now()
		tempoGasto = fim.Sub(comeco)

	} else if algoritmo == 3 {
		comeco := time.Now()

		ordenacao.HeapSort()

		fim := time.Now()
		tempoGasto = fim.Sub(comeco)

	} else if algoritmo == 4 {
		comeco := time.Now()

		ordenacao.InsertionSort()

		fim := time.Now()
		tempoGasto = fim.Sub(comeco)

	} else if algoritmo == 5 {
		comeco := time.Now()

		ordenacao.QuickSort()

		fim := time.Now()
		tempoGasto = fim.Sub(comeco)

	} else if algoritmo == 6 {
		comeco := time.Now()

		ordenacao.SelectionSort()

		fim := time.Now()
		tempoGasto = fim.Sub(comeco)

	} else if algoritmo == 7 {
		comeco := time.Now()

		ordenacao.ShellSort()

		fim := time.Now()
		tempoGasto = fim.Sub(comeco)

	}

	ordenacao.Mostrar()
	fmt.Println("Está ordenado:", ordenacao.IsOrdenado())
	fmt.Println("Tempo gasto na ordenação:", tempoGasto.String())
}
