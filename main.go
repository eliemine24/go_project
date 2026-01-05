// ============
// === MAIN ===
// ============

package main

import (
	"fmt"
)

// Génère une matrice carrée NxN initialisée avec des flottants, tout est à 0
func initMatriceZero(n int, out chan<- [][]float64) {
	matrice := make([][]float64, n)

	for i := 0; i < n; i++ {
		matrice[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			matrice[i][j] = 0.0
		}
	}

	out <- matrice
}

func main() {
	n := 80
	ch := make(chan [][]float64)

	go initMatriceZero(n, ch)

	matrice := <-ch // attend la goroutine

	fmt.Println("Matrice générée :")
	for _, ligne := range matrice {
		fmt.Println(ligne)
	}
}
