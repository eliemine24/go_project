package matrix

import "fmt"

// Init Matrice
func init() {
	matrice = make([][]float64, TAILLE)
	for i := 0; i < TAILLE; i++ {
		matrice[i] = make([]float64, TAILLE)
	}
}

// ===========================
// === Fusions de matrices ===
// ===========================

// Concatener des matrices en une matrice plus grande (vérifier qu'on obtient une matrice carrée)
// ajouter les valeurs d'une matrice plus petite sur une matrice plus grande
// x, y coordonnées de la première valeur de la petite matrice dans la plus grande, N taille de la matrice d'entrée
func AjouterParcelle(matrice [][]float64, N int, x int, y int, out chan<- [][]float64) {
	if x+N-1 >= len(out[0]) || y+N-1 >= len(out) {
		out <- matrice // passe à la suite
		fmt.Print("erreur : matrices incompatibles")
		return
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			out[i+x][j+y] = matrice[i][j]
		}
	}
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
		// moyenne avec deux valeurs adjacentes (colonne gauche et droite)
		matrice[i][X] = (matrice[i][X-1] + matrice[i][X] + matrice[i][X+1]) / 3
	}
	out <- matrice
}

func main(){
	m1 := [][]int{
	{1, 2, 3},
	{4, 5, 6},
	{7, 8, 9},
}

	m2 = make([][]float64, 9)
		for i := 0; i < 9; i++ {
			m2[i] = make([]float64, 9)
		}

	print(m1)
	print(m2)
}