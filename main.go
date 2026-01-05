// ============
// === MAIN ===
// ============

package main

import (
	"fmt"
)

// Génère une matrice carrée NxN initialisée avec des flottants
func initMatrice(n int, out chan<- [][]float64) {
	matrice := make([][]float64, n)

	for i := 0; i < n; i++ {
		matrice[i] = make([]float64, n)
	}

	out <- matrice
}

// Générer un bruit de perlin sur une matrice vide carree de taille n
func perlin(matrice [][]float64, out chan<- [][]float64) {

	out <- matrice
}

// Concatener des matrices en une matrice plus grande (vérifier qu'on obtient une matrice carrée)
// ajouter les valeurs d'une matrice plus petite sur une matrice plus grande
func ajouterParcelle(matrice [][]float64, out chan<- [][]float64) {

	out <- matrice
}

// Moyenner toute une ligne d'une matrice en fonction des valeurs alentours
func avgOnLine(matrice [][]float64, out chan<- [][]float64) {

	out <- matrice
}

// Moyenner toute une colonne d'une matrice en fonction des valeurs alentours
func avgOnColumn(matrice [][]float64, out chan<- [][]float64) {

	out <- matrice
}

// Affichage graphique d'une matrice carrée de taille n contenant des flottants entre 0 et 1
func displayMat(matrice [][]float64) {

}

// Fonction principale génère la carte finale et l'affiche
func main() {
	n := 80
	ch := make(chan [][]float64)

	go initMatrice(n, ch)

	matrice := <-ch // attend la goroutine

	fmt.Println("Matrice générée :")
	for _, ligne := range matrice {
		fmt.Println(ligne)
	}
}
