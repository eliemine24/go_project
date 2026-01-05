// ============
// === MAIN ===
// ============

package main

import "fmt"

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
// N taille de la matrice, Y coordonnee y de la ligne à moyenner
func avgOnLine(matrice [][]float64, N int, Y int, out chan<- [][]float64) {
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
func avgOnColumn(matrice [][]float64, N int, X int, out chan<- [][]float64) {
	// verif bug
	if X <= 0 || X >= N-1 {
		out <- matrice // passe à la suite
		fmt.Print("erreur : taille de matrice incompatible")
		return
	}
	for i := 0; i < N; i++ {
		//moyenne avec deux valeurs adjacentes
		matrice[i][X] = (matrice[i][X-1] + matrice[i][X] + matrice[i+1][X]) / 3
	}
	out <- matrice
}

// Affichage graphique d'une matrice carrée de taille n contenant des flottants entre 0 et 1
func displayMat(matrice [][]float64) {

}

// Fonction principale génère la carte finale et l'affiche
func main() {

	// Initialisation des valeurs des grandeurs utilisées
	matsize := 10                           // taille des petites matrices
	finalMatSize := 10 * matsize            // taille de la matrice finale (carrée de 100 par 100 pour l'instant)
	matriceFinale := make(chan [][]float64) // matrice finale de taille finalmatsize
	go initMatrice(finalMatSize, matriceFinale)
	// attendre que la goroutine soi terminée avec un waitgroup, j'y reviens plus tard

	// Moyennage simultanné des lignes de la matrice finale, puis des colomnes
	// test à réalise plus tard

}
