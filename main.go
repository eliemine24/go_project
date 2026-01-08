// ============
// === MAIN ===
// ============

package main

import (
	"gns/matrix"
)

const (
	MAPSIZE      = 10                               // taille des maps elementaires
	RATIO        = 10                               // rapport de la taille de la map finale par la taille des maps elementaires
	FINALMAPSIZE = MAPSIZE * RATIO                  // taille de la map finale
	MAPNB        = FINALMAPSIZE / MAPSIZE           // nombre de map elementaires sur la map finale
	FINALMAP     = matrix.InitMatrice(FINALMAPSIZE) // Map finale (matrice carrée de flottant)
)

// canal pour la génération des matrices bruitées
c := make(chan [][]float64)

func GenerateNoisedMap(size int, out <- [][]float64){
	
}

func LauchGeneration(n int, c chan){
	for(N){
		c <- GenerateElemMap(size)
	}

}

func main() {
	// Générer les matrices et générer perlin dessus
	// alimenter le canal avec la go routine de génération de perlin noise

	// func Worker(canal){
	// for (j in range c)
	// do Job(j) }

	// for (number <K){go Worker}

}
