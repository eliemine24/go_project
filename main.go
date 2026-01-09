// ============
// === MAIN ===
// ============

package main

import (
	"gns/matrix"
	"gns/perlin"
)

const (
	MAPSIZE      = 10
	RATIO        = 10
	FINALMAPSIZE = MAPSIZE * RATIO
	NBMAPS       = FINALMAPSIZE / MAPSIZE
)

func main() {
	// Canal pour récup les maps elementaires générées par perlin
	out := make(chan [][]float64)

	// Liste des NBMAPS matrices générées et stockées
	var MAPLIST [][][]float64

	// Lancement parallèle de initMat puis perlin
	for i := 0; i < NBMAPS; i++ {
		m := matrix.InitMatrice(MAPSIZE)
		go perlin.GeneratePerlin(m, out)
	}

	// Recep des maps élémentaires depuis les canaux sefdvsdfv
	for i := 0; i < NBMAPS; i++ {
		matrice := <-out
		MAPLIST = append(MAPLIST, matrice)
	}
}
