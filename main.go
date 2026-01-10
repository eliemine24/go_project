// ============
// === MAIN ===
// ============

package main

import (
	"fmt"
	"gns/matrix"
	"gns/perlin"
	"sync"
)

const (
	MAPSIZE      = 10
	RATIO        = 10
	FINALMAPSIZE = MAPSIZE * RATIO
	NBMAPS       = RATIO * RATIO
)

func main() {
	// Canal pour récup les maps elementaires générées par perlin
	out := make(chan [][]float64)

	// Liste des NBMAPS matrices générées et stockées
	var MAPLIST [][][]float64

	// Lancement parallèle de initMat puis perlin
	fmt.Println("--- launching perlin generation ---")
	for i := 0; i < NBMAPS; i++ {
		m := matrix.InitMatrice(MAPSIZE)
		go perlin.GeneratePerlin(m, out)
	}

	// Recep des maps élémentaires depuis les canaux
	fmt.Println("--- recep maps ---")
	for i := 0; i < NBMAPS; i++ {
		matrice := <-out
		MAPLIST = append(MAPLIST, matrice)
	}

	fmt.Println(MAPLIST)

	// Init la matrice finale
	fmt.Println("--- initialize finalmap ---")
	FINALMAP := matrix.InitMatrice(FINALMAPSIZE)

	// Creer un waitgroup pour l'ajout des matrices
	var wg sync.WaitGroup

	// Lancement en parallèle de l'ajout des MAPS sur FINALMAP
	index := 0
	fmt.Println("--- boucle c'est tipar ---")
	for ty := 0; ty < RATIO; ty++ {
		for tx := 0; tx < RATIO; tx++ {

			source := MAPLIST[index]
			x := tx * MAPSIZE
			y := ty * MAPSIZE
			fmt.Println(index)

			wg.Add(1)
			go func(src [][]float64, x, y int) {
				defer wg.Done()
				matrix.AjouterParcelle(src, MAPSIZE, x, y, FINALMAP)
			}(source, x, y)

			index++
		}
	}

	wg.Wait() // attendre que la concaténation soit terminée.

	fmt.Print(FINALMAP)
}
