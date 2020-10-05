package main

import (
	"fmt"
	"time"

	"./fronteira"
	"./ordenacao"
)

// main inicia o programa.
func main() {

	var repetir = true
	for repetir {
		var algoritmo = obterAlgoritmo()

		if algoritmo != 0 {
			ordenacao.SetTamanho(obterTamanhoArranjo())
			var ordem = obterOrdemArranjo()
			executarOrdenacao(ordem, algoritmo)
		} else {
			repetir = false
		}

	}

}

// obterTamanhoArranjo obtem do usuario o tamanho do arranjo que sera ordenado.
// Retorna um numero inteiro contendo o tamanho do arranjo.
func obterTamanhoArranjo() int {
	fmt.Println("Qual o tamanho do arranjo?")
	return int(fronteira.ReadInt64())

}

// obterAlgoritmo obtem do usuario o algoritmo de ordenacao que deseja executar.
// Retorna um numero inteiro unsigned de 64 bits contendo o numero do algoritmo de ordenacao.
func obterAlgoritmo() uint8 {
	var algoritmo uint8

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

		algoritmo = fronteira.ReadUint8()

		if algoritmo <= 7 {
			break
		}
	}

	return algoritmo
}

// obterOrdemArranjo obtem do usuario a ordem do arranjo que sera ordenado.
// Retorna um numero inteiro unsigned de 64 bits contendo o numero da ordem do arranjo que sera ordenado.
func obterOrdemArranjo() uint8 {
	var ordem uint8

	for {
		fmt.Println("Como deseja que o arranjo de numeros esteja?")
		fmt.Println("0 - Aleatorio")
		fmt.Println("1 - Decrescente")

		ordem = fronteira.ReadUint8()

		if ordem <= 1 {
			break
		}
	}

	return ordem
}

// executarOrdenacao executa a ordenacao do arranjo de caracteres.
// Parametro ordem: Numero indicando a ordem que o arranjo estara antes de ser ordenado.
// Parametro algoritmo: Numero indicando o algoritmo de ordenacao de sera utilizado.
func executarOrdenacao(ordem uint8, algoritmo uint8) {

	//Ordenando o arranjo de acordo com a ordem selecionada
	if ordem == 0 {
		ordenacao.Aleatorio()
	} else if ordem == 1 {
		ordenacao.Decrescente()
	}

	fmt.Println("Arranjo de inteiros antes de ordenar:")
	ordenacao.Mostrar()
	fmt.Println("")

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

	fmt.Println("Arranjo de inteiros depois de ordenar:")
	ordenacao.Mostrar()
	fmt.Println("")
	fmt.Println("Está ordenado:", ordenacao.IsOrdenado())
	fmt.Println("Tempo gasto na ordenação:", tempoGasto.String())
}
