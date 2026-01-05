// ============
// === MAIN ===
// ============

package main

import (
	"image/color"
	"math"
	"math/rand"
	"time"

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
func avgOnLine(matrice [][]float64, out chan<- [][]float64) {

	out <- matrice
}

// Moyenner toute une colonne d'une matrice en fonction des valeurs alentours
func avgOnColumn(matrice [][]float64, out chan<- [][]float64) {

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
