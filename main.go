// ============
// === MAIN ===
// ============

package main

import (
	"fmt"
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/palette/moreland"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// matrice utilisée par main() pour l'affichage (evite variable non définie dans main)
var matrice [][]float64

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
