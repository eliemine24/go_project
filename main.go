// ============
// === MAIN ===
// ============

package main

import (
	"fmt"
	"image/color"
	"math"
	"math/rand"

	"gonum.org/v1/plot"
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

// Génère une matrice carrée NxN initialisée avec des flottants
func initMatrice(n int, out chan<- [][]float64) {
	matrice := make([][]float64, n)

	for i := 0; i < n; i++ {
		matrice[i] = make([]float64, n)
	}

	out <- matrice
}

// ==============================
// Fonctions Perlin
// ==============================

func smoothstep(w float64) float64 {
	if w <= 0 {
		return 0
	}
	if w >= 1 {
		return 1
	}
	return w * w * (3 - 2*w)
}

func interpolate(a0, a1, w float64) float64 {
	return a0 + (a1-a0)*smoothstep(w)
}

type Gradient struct {
	x, y float64
}

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

func dotGridGradient(ix, iy int, x, y float64, gradients [][]Gradient) float64 {
	dx := x - float64(ix)
	dy := y - float64(iy)
	g := gradients[iy][ix]
	return dx*g.x + dy*g.y
}

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
		//moyenne avec deux valeurs adjacentes
		matrice[i][X] = (matrice[i][X-1] + matrice[i][X] + matrice[i+1][X]) / 3
	}
	out <- matrice
}

// Affichage graphique d'une matrice carrée de taille n contenant des flottants entre 0 et 1
func displayMat(matrice [][]float64) {

}

// ==============================
// Heatmap interface
// ==============================

type HeatmapData [][]float64

func (h HeatmapData) Dims() (c, r int) {
	return len(h[0]), len(h)
}

func (h HeatmapData) Z(c, r int) float64 {
	return h[r][c]
}

func (h HeatmapData) X(c int) float64 {
	return float64(c)
}

func (h HeatmapData) Y(r int) float64 {
	return float64(r)
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
