// ============
// === MAIN ===
// ============

package main

import (
	"fmt"
	"gns/matrix"
	"gns/perlin"
)

const (
	MAPSIZE      = 100
	RATIO        = 100
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

	// Recep des maps élémentaires depuis les canaux
	for i := 0; i < NBMAPS; i++ {
		matrice := <-out
		MAPLIST = append(MAPLIST, matrice)
	}

	fmt.Print(MAPLIST)
}
