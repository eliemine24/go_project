// ============
// === MAIN ===
// ============

package main

import (
	"fmt"
	"gns/display"
	"gns/matrix"
	"gns/perlin"
	"sync"
)

const (
	MAPSIZE      = 100             // taille des maps élémentaires
	RATIO        = 10              // nombre de maps élémentaires du un coté de map finale
	FINALMAPSIZE = MAPSIZE * RATIO // taille de la map finale
	NBMAPS       = RATIO * RATIO   // nombre de maps élémentaires sur la map finale
	AVGWIDE      = 40              // largeur du moyennage
	AVGNUMBER    = 100             //nombre de moyennages successifs
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

			wg.Add(1)
			go func(src [][]float64, x, y int) {
				defer wg.Done()
				matrix.AjouterParcelle(src, MAPSIZE, x, y, FINALMAP)
			}(source, x, y)

			index++
		}
	}
	wg.Wait() // attendre que la concaténation soit terminée.
	fmt.Println("--- fin ajout matrices ---")

	// Moyennage pour adoucir les bords
	fmt.Println("-- début moyennage lignes ---")

	// On commence par les moyennages sur les colomnes pour éviter les chevauchements
	var wgx sync.WaitGroup

	for tx := 1; tx < RATIO; tx++ {

		x := tx * MAPSIZE

		wgx.Add(1)

		go func(FINALMAP [][]float64, x int) {
			defer wgx.Done()
			for i := 0; i < AVGNUMBER; i++ { //plusieurs moyannages successif parce qu'on est des bourrins
				matrix.AvgOnLine(FINALMAP, x, AVGWIDE)
			}
		}(FINALMAP, x)

	}
	wgx.Wait()

	// Moyennage pour adoucir les bords
	fmt.Println("-- début moyennage colonnes ---")

	// On commence par les moyennages sur les colomnes pour éviter les chevauchements
	var wgy sync.WaitGroup

	for ty := 1; ty < RATIO; ty++ {

		y := ty * MAPSIZE

		wgy.Add(1)

		go func(FINALMAP [][]float64, y int) {
			defer wgy.Done()
			for i := 0; i < AVGNUMBER; i++ { //plusieurs moyannages successif parce qu'on est des bourrins
				matrix.AvgOnColumn(FINALMAP, y, AVGWIDE)
			}
		}(FINALMAP, y)

	}
	wgy.Wait()

	// afficher la matrice finie avec display.showmat
	display.ShowMat(FINALMAP, MAPSIZE)
}
