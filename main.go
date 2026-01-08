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

// canal pour la génération des matrices bruitées

func jobfeeder_maps(c chan<- [][]float64) {
	for i:=0; i < MAPNB; i+=MAPSIZE {
		for j:=0; j < MAPNB; j+=MAPSIZE {
			c <- (i, j)
		}
	}
	close c
}

func jobfeeder_avg(k chan<- [][]float64) {
	for i:=MAPSIZE; i < FINALMAPSIZE; i+=MAPSIZE {
		k <- (i, 0)
	}
	for j:=MAPSIZE; j < FINALMAPSIZE; j+=MAPSIZE {
		k <- (0, j)
	}
	close k
}

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
