package noise

import (
	// affichages compil
	//affichage images
	"math"      // takapte
	"math/rand" //aleatoire
	//des trucs obscurs à détailler
)

const (
	SCALE         = 8.0
	POIDS_MOYENNE = 0.8
	POIDS_PERLIN  = 0.2
)

// ========================
// === Fonctions Perlin ===
// ========================

// smoothstep applique l'interpolation smoothstep sur w (borné entre 0 et 1).
func smoothstep(w float64) float64 {
	if w <= 0 {
		return 0
	}
	if w >= 1 {
		return 1
	}
	return w * w * (3 - 2*w)
}

// Fonction d'interpolation lisse entre a0 et a1
// Le poids w doit être dans l'intervalle [0.0, 1.0]
func interpolate(a0, a1, w float64) float64 {
	return a0 + (a1-a0)*smoothstep(w)
}

// Définit un type gradient pour la génération de bruit
type Gradient struct {
	x, y float64
}

// generateGradients crée une grille de vecteurs gradients aléatoires pour le bruit de Perlin.
func generateGradients(width, height int) [][]Gradient {
	gradients := make([][]Gradient, height+1)
	for y := 0; y <= height; y++ {
		gradients[y] = make([]Gradient, width+1)
		for x := 0; x <= width; x++ {
			angle := rand.Float64() * 2 * math.Pi //génère un vecteur unitaire dans une direction aléatoire.
			gradients[y][x] = Gradient{           //conversion angle en vecteur unitaire
				x: math.Cos(angle),
				y: math.Sin(angle),
			}
		}
	}
	return gradients
}

// dotGridGradient calcule le produit scalaire entre le vecteur distance (dx,dy)
// et le gradient stocké à ce coin de la grille (ix,iy).
func dotGridGradient(ix, iy int, x, y float64, gradients [][]Gradient) float64 {
	dx := x - float64(ix)
	dy := y - float64(iy)
	g := gradients[iy][ix] //
	return dx*g.x + dy*g.y
}

// perlin calcule la valeur du bruit de Perlin en (x,y) à partir des gradients fournis.
func perlin(x, y float64, gradients [][]Gradient) float64 {
	x0 := int(math.Floor(x)) //On détermine les quatre coins de la cellule dans laquelle se trouve le point (x,y).
	x1 := x0 + 1
	y0 := int(math.Floor(y))
	y1 := y0 + 1

	sx := x - float64(x0)
	sy := y - float64(y0)

	n0 := dotGridGradient(x0, y0, x, y, gradients) //produits scalaires aux 4 coins
	n1 := dotGridGradient(x1, y0, x, y, gradients)
	ix0 := interpolate(n0, n1, sx) //ligne du haut.

	n0 = dotGridGradient(x0, y1, x, y, gradients) //ligne du bas.
	n1 = dotGridGradient(x1, y1, x, y, gradients)
	ix1 := interpolate(n0, n1, sx)

	return interpolate(ix0, ix1, sy)
}

// Générer un bruit de perlin sur une matrice vide carree de taille n
func GeneratePerlin(matrice [][]float64, out chan<- [][]float64) {
	n := len(matrice) // use rows
	gradients := generateGradients(n, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {

			// Bruit de Perlin normalisé
			p := perlin(float64(i)/SCALE, float64(j)/SCALE, gradients)
			p = (p + 1) / 2

			var moyenne float64

			switch {
			case i == 0 && j == 0:
				// Coin supérieur gauche : pas de voisins
				moyenne = p

			case i == 0:
				// Première ligne : on ne peut prendre que la valeur à gauche
				moyenne = matrice[i][j-1]

			case j == 0:
				// Première colonne : on ne peut prendre que la valeur au-dessus
				moyenne = matrice[i-1][j]

			default:
				// Cas général : moyenne des deux voisins
				moyenne = (matrice[i][j-1] + matrice[i-1][j]) / 2
			}

			// Mélange pondéré
			valeur := POIDS_MOYENNE*moyenne + POIDS_PERLIN*p

			// Clamp entre 0 et 1
			matrice[i][j] = math.Max(0, math.Min(1, valeur))
		}
	}

	out <- matrice
}
