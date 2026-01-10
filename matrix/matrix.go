package matrix

import "fmt"

// Init Matrice
func InitMatrice(size int) [][]float64 {
	m := make([][]float64, size)
	for i := 0; i < size; i++ {
		m[i] = make([]float64, size)
	}
	return m
}

// ===========================
// === Fusions de matrices ===
// ===========================

// Concatener des matrices en une matrice plus grande (vérifier qu'on obtient une matrice carrée)
// ajouter les valeurs d'une matrice plus petite sur une matrice plus grande
// x, y coordonnées de la première valeur de la petite matrice dans la plus grande, N taille de la matrice d'entrée
// chat m'a donné la fonction copy(dest, source) qui est acrréement plus efficace
func AjouterParcelle(source [][]float64, N, x, y int, dest [][]float64) {
	for i := 0; i < N; i++ {
		copy(dest[x+i][y:y+N], source[i])
	}
}

// Moyenner toute une ligne d'une matrice en fonction des valeurs alentours
// N taille de la matrice, Y coordonnee y de la ligne à moyenner
func AvgOnLine(matrice [][]float64, N int, Y int, out chan<- [][]float64) {
	// verif bug
	if Y <= 0 || Y >= N-1 {
		out <- matrice // passe à la suite
		fmt.Print("erreur : taille de matrice incompatible")
		return
	}
	for i := 0; i < N; i++ {
		// use vertical neighbors
		matrice[Y][i] = (matrice[Y-1][i] + matrice[Y][i] + matrice[Y+1][i]) / 3
	}
	out <- matrice
}

// Moyenner toute une colonne d'une matrice en fonction des valeurs alentours
func AvgOnColumn(matrice [][]float64, N int, X int, out chan<- [][]float64) {
	// verif bug
	if X <= 0 || X >= N-1 {
		out <- matrice // passe à la suite
		fmt.Print("erreur : taille de matrice incompatible")
		return
	}
	for i := 0; i < N; i++ {
		// moyenne avec deux valeurs adjacentes (colonne gauche et droite)
		matrice[i][X] = (matrice[i][X-1] + matrice[i][X] + matrice[i][X+1]) / 3
	}
	out <- matrice
}
