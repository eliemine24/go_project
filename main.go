// ============
// === MAIN ===
// ============

package main

import (
	"fmt"
	"gns/display"
	"gns/matrix"
	"gns/perlin"
)

const (
	MAPSIZE      = 10
	RATIO        = 10
	FINALMAPSIZE = MAPSIZE * RATIO
	MAPNB        = FINALMAPSIZE / MAPSIZE
)

func main() {

	// TESTS DE PERLIN ET AFFICHAGE
	// init channel(s)
	initCh := make(chan [][]float64)
	perlinCh := make(chan [][]float64)

	// init matrix (runs as goroutine that sends the matrix on initCh)
	go matrix.InitMatrice(MAPSIZE, initCh)
	TESTMAP := <-initCh

	// run perlin generator (expects to send result on perlinCh)
	go perlin.GeneratePerlin(TESTMAP, perlinCh)
	TESTMAP = <-perlinCh

	fmt.Print(TESTMAP)
	display.ShowMat(TESTMAP, 10.0)
	// FIN TEST
}
