// ============
// === MAIN ===
// ============

package main

import (
	"fmt"         // affichages compil
	"image/color" //affichage images
	"math"        // takapte
	"math/rand"   //aleatoire

	"gonum.org/v1/plot" //des trucs obscurs à détailler
	"gonum.org/v1/plot/palette/moreland"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

const (
	TAILLE        = 50
	SCALE         = 8.0
	POIDS_MOYENNE = 0.8
	POIDS_PERLIN  = 0.2
)

// matrice utilisée par main() pour l'affichage (evite variable non définie dans main)
var matrice [][]float64

func init() {
	matrice = make([][]float64, TAILLE)
	for i := 0; i < TAILLE; i++ {
		matrice[i] = make([]float64, TAILLE)
	}
}

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
func effectuerperlin(matrice [][]float64, out chan<- [][]float64) {
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

// ===========================
// === Fusions de matrices ===
// ===========================

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
		// moyenne avec deux valeurs adjacentes (colonne gauche et droite)
		matrice[i][X] = (matrice[i][X-1] + matrice[i][X] + matrice[i][X+1]) / 3
	}
	out <- matrice
}

// Affichage graphique d'une matrice carrée de taille n contenant des flottants entre 0 et 1
func displayMat(matrice [][]float64) {

}

// =========================
// === Heatmap interface ===
// =========================

type HeatmapData [][]float64

// Dims retourne le nombre de colonnes (c) et de lignes (r) pour l'interface HeatmapData.
func (h HeatmapData) Dims() (c, r int) {
	return len(h[0]), len(h)
}

// Z retourne la valeur z à la colonne c et ligne r.
func (h HeatmapData) Z(c, r int) float64 {
	return h[r][c]
}

// X retourne la coordonnée x correspondant à la colonne c (identité ici).
func (h HeatmapData) X(c int) float64 {
	return float64(c)
}

// Y retourne la coordonnée y correspondant à la ligne r (identité ici).
func (h HeatmapData) Y(r int) float64 {
	return float64(r)
}

// ============
// === MAIN ===
// ============
// Fonction principale génère la carte finale et l'affiche
func main() {

	// Je t'en supplie valentin comprend et explique nous ensuite
	p := plot.New()

	p.X.Min = 0
	p.X.Max = TAILLE
	p.Y.Min = 0
	p.Y.Max = TAILLE

	p.X.Tick.Marker = plot.ConstantTicks([]plot.Tick{})
	p.Y.Tick.Marker = plot.ConstantTicks([]plot.Tick{})

	p.X.LineStyle.Width = 0
	p.Y.LineStyle.Width = 0

	cm := moreland.Kindlmann()
	cm.SetMax(1)
	cm.SetMin(0)

	palette := cm.Palette(255)
	hm := plotter.NewHeatMap(HeatmapData(matrice), palette)
	hm.NaN = color.Transparent

	p.Add(hm)

	if err := p.Save(6*vg.Inch, 6*vg.Inch, "perlin_heatmap.png"); err != nil {
		panic(err)
	}
}
