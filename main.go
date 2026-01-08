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
	MAPSIZE      = 50
	RATIO        = 10
	FINALMAPSIZE = MAPSIZE * RATIO
	MAPNB        = FINALMAPSIZE / MAPSIZE
)

// canal pour la génération des matrices bruitées

func main() {
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
}
