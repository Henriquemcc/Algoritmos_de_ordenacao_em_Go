package ordenacao

import (
	"fmt"
	"math/rand"
)

var n = 100
var array []int
var i int
var j int

//Produz um arranjo ordenado de modo decrescente.
func Decrescente() {
	for i := 0; i < n; i++ {
		array = append(array, n-1-i)
	}
}

//Produz um arranjo de numeros aleatorios.
func Aleatorio() {
	for i = 0; i < n; i++ {
		array = append(array, rand.Int()%1000)
	}
}

//Mostrar os elemento de um arranjo.
func Mostrar() {
	fmt.Printf("[ ")

	for i = 0; i < n; i++ {
		fmt.Printf("%d ", array[i])
	}

	fmt.Printf("] \n")
}

func swap(i int, j int) {
	temp := array[i]
	array[i] = array[j]
	array[j] = temp
}

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

func SetTamanho(tamanho int) {
	n = tamanho
}

func gerarArranjo(tamanho int) []int {
	var arranjo []int

	for i := 0; i < tamanho; i++ {
		arranjo = append(arranjo, 0)
	}

	return arranjo
}
