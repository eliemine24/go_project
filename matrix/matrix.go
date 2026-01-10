package matrix

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
func AvgOnColumn(matrice [][]float64, x int) {
	long := len(matrice)

	// Moyennage sur deux cases à partir des bords des maps, x se situe ici :
	// ----map1----|---map2----
	// ....][_][_] | [x][_][...
	for y := 0; y < long; y++ {
		matrice[x][y] = (matrice[x-3][y] + matrice[x-2][y] + matrice[x-1][y]) / 3
		matrice[x][y] = (matrice[x][y] + matrice[x+1][y] + matrice[x+2][y]) / 3
		matrice[x][y] = (matrice[x-2][y] + matrice[x-1][y] + matrice[x][y]) / 3
		matrice[x][y] = (matrice[x-1][y] + matrice[x][y] + matrice[x][y]) / 3
	}
