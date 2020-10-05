package ordenacao

import (
	"fmt"
	"math/rand"
)

// n tamanho do arranjo de numeros inteiros.
var n = 100

// array arranjo de numeros inteiros.
var array []int

// i cabecote primario.
var i int

// j cabecote secundario.
var j int

// Decrescente produz um arranjo ordenado de modo decrescente.
func Decrescente() {
	for i := 0; i < n; i++ {
		array = append(array, n-1-i)
	}
}

// Aleatorio produz um arranjo de numeros aleatorios.
func Aleatorio() {
	for i = 0; i < n; i++ {
		array = append(array, rand.Int()%1000)
	}
}

// Mostrar mostra os elemento de um arranjo.
func Mostrar() {
	fmt.Printf("[ ")

	for i = 0; i < n; i++ {
		fmt.Printf("%d ", array[i])
	}

	fmt.Printf("] \n")
}

// Swap troca os valores de duas posicoes do arranjo.
// Parametro i: posicao do elemento que sera trocado pelo elemento da posicao j.
// Parametro j: posiaco do elemento que sera trocado pelo elemento da posicao i.
func swap(i int, j int) {
	temp := array[i]
	array[i] = array[j]
	array[j] = temp
}

// IsOrdenado verifica se o arranjo de numeros inteiros esta ordenado.
// Retorna um valor booleano indicando se o arranjo esta ordenado.
func IsOrdenado() bool {
	resp := true
	for i := 1; i < n; i++ {
		if array[i] < array[i-1] {
			i = n
			resp = false
		}
	}
	return resp
}

// SetTamanho serve para alterar o tamanho do arranjo de numero inteiros.
// Parametro tamanho: Numero inteiro indicando o tamanho do arranjo de numeros inteiros.
func SetTamanho(tamanho int) {
	n = tamanho
}

// gerarArranjo serve para adicionar zero a todas as posicoes do arranjo de inteiros ate ele ter o tamanho igual ao tamanho especificado.
// Parametro tamanho: Numero inteiro indicando o tamanho do arranjo de numeros inteiros para que possa ser adicionados os zeros.
func gerarArranjo(tamanho int) []int {
	var arranjo []int

	for i := 0; i < tamanho; i++ {
		arranjo = append(arranjo, 0)
	}

	return arranjo
}
