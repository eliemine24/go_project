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

// Moyenner toute une colonne d'une matrice en fonction des valeurs alentours
// sur une largeur avgwide
func AvgOnLine(matrice [][]float64, x int, avgwide int) {
	long := len(matrice)

	// antibug haha
	if x-avgwide < 0 || x+avgwide >= long {
		return
	}

	// Moyennage sur deux cases à partir des bords des maps, x se situe ici :
	// ----map1----|---map2----
	// ....][_][_] | [x][_][...
	for y := 0; y < long; y++ {
		sum := 0.0
		for tx := x - avgwide; tx < x+avgwide; tx++ {
			sum += matrice[tx][y]
		}
		matrice[x][y] = sum / float64(2*avgwide)
	}
}

// Moyenner toute une ligne d'une matrice en fonction des valeurs alentours
// sur une largeur avgwide
func AvgOnColumn(matrice [][]float64, y int, avgwide int) {
	long := len(matrice)

	// antibug haha
	if y-avgwide < 0 || y+avgwide >= long {
		return
	}

	for x := 0; x < long; x++ {
		sum := 0.0
		for ty := y - avgwide; ty < y+avgwide; ty++ {
			sum += matrice[x][ty]
		}
		matrice[x][y] = sum / float64(2*avgwide)
	}
}
