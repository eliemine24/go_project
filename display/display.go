package display

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/palette/moreland"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

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

// Affichage final de la matrice
func ShowMat(FINALMAP [][]float64, MAPSIZE float64) {

	p := plot.New()

	p.X.Min = 0
	p.X.Max = MAPSIZE
	p.Y.Min = 0
	p.Y.Max = MAPSIZE

	p.X.Tick.Marker = plot.ConstantTicks([]plot.Tick{})
	p.Y.Tick.Marker = plot.ConstantTicks([]plot.Tick{})

	p.X.LineStyle.Width = 0
	p.Y.LineStyle.Width = 0

	cm := moreland.Kindlmann()
	cm.SetMax(1)
	cm.SetMin(0)

	palette := cm.Palette(255)
	hm := plotter.NewHeatMap(HeatmapData(FINALMAP), palette)
	hm.NaN = color.Transparent

	p.Add(hm)

	if err := p.Save(6*vg.Inch, 6*vg.Inch, "perlin_heatmap.png"); err != nil {
		panic(err)
	}
}
