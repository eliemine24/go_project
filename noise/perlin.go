package noise

import (
	// affichages compil
	//affichage images
	"math"      // takapte
	"math/rand" //aleatoire
	//des trucs obscurs à détailler
)

const (
	TAILLE        = 50
	SCALE         = 8.0
	POIDS_MOYENNE = 0.8
	POIDS_PERLIN  = 0.2
)

//lsugflzdefuv
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

// interpolate interpole entre a0 et a1 en utilisant smoothstep(w) comme pondération.
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
			angle := rand.Float64() * 2 * math.Pi
			gradients[y][x] = Gradient{
				x: math.Cos(angle),
				y: math.Sin(angle),
			}
		}
	}
	return gradients
}

// dotGridGradient calcule le produit scalaire entre le vecteur distance (dx,dy)
// et le gradient stocké à la grille (ix,iy).
func dotGridGradient(ix, iy int, x, y float64, gradients [][]Gradient) float64 {
	dx := x - float64(ix)
	dy := y - float64(iy)
	g := gradients[iy][ix]
	return dx*g.x + dy*g.y
}

// perlin calcule la valeur du bruit de Perlin en (x,y) à partir des gradients fournis.
func perlin(x, y float64, gradients [][]Gradient) float64 {
	x0 := int(math.Floor(x))
	x1 := x0 + 1
	y0 := int(math.Floor(y))
	y1 := y0 + 1

	sx := x - float64(x0)
	sy := y - float64(y0)

	n0 := dotGridGradient(x0, y0, x, y, gradients)
	n1 := dotGridGradient(x1, y0, x, y, gradients)
	ix0 := interpolate(n0, n1, sx)

	n0 = dotGridGradient(x0, y1, x, y, gradients)
	n1 = dotGridGradient(x1, y1, x, y, gradients)
	ix1 := interpolate(n0, n1, sx)

	return interpolate(ix0, ix1, sy)
}

// Générer un bruit de perlin sur une matrice vide carree de taille n
func GeneratePerlin(matrice [][]float64, out chan<- [][]float64) {
	gradients := generateGradients(len(matrice[0]), len(matrice[0]))

	for i := 1; i < len(matrice[0]); i++ {
		for j := 1; j < len(matrice[0]); j++ {
			moyenne := (matrice[i][j-1] + matrice[i-1][j]) / 2

			p := perlin(float64(i)/SCALE, float64(j)/SCALE, gradients)
			p = (p + 1) / 2

			valeur := POIDS_MOYENNE*moyenne + POIDS_PERLIN*p
			matrice[i][j] = math.Max(0, math.Min(1, valeur))
		}
	}
	out <- matrice
}
