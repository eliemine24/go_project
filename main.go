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
	MAPSIZE      = 10
	RATIO        = 10
	FINALMAPSIZE = MAPSIZE * RATIO
	MAPNB        = FINALMAPSIZE / MAPSIZE
)

func GenerateMaps(maplist [][][]float64, c chan[][]float64){
	for (maplist range(c)){
		go matrix.InitMatrice(MAPSIZE, c)
	}
}

func main() {
	MAPLIST := [][][]float64 // liste des maps élémentaires

	// init channel(s)
	initCh := make(chan [][]float64)
	perlinCh := make(chan [][]float64)

	// init matrix (runs as goroutine that sends the matrix on initCh)
	for(MAPNB){
		go matrix.InitMatrice(MAPSIZE, initCh)
	TESTMAP := <-initCh

	}
	go matrix.InitMatrice(MAPSIZE, initCh)
	TESTMAP := <-initCh

	// run perlin generator (expects to send result on perlinCh)
	go perlin.GeneratePerlin(TESTMAP, perlinCh)
	TESTMAP = <-perlinCh

	fmt.Print(TESTMAP)
}
