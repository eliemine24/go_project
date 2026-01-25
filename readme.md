# ELP - Go Project

Élie GAUTIER, Léa DANOBER, Valentin-Jules DUMAS

## Objectif

Paralléliser la création d'une carte de relief par [génération procédurale](https://fr.wikipedia.org/wiki/G%C3%A9n%C3%A9ration_proc%C3%A9durale) avec Go. On cherche à déterminer s'il est possible d'optimiser le processus de génération d'une carte en le parallélisant avec plusieurs goroutines simultanées.

Inspiration pour l'idée de départ : [underscore_ | Le problème de la génaration procédurale](https://www.youtube.com/watch?v=Uh-4KO33y6c)

## Structure

Le projet se structure en quatre packages principaux :

- **`matrix`** : Gère les opération sur les matrices modélisant les cartes
- **`display`** : Contient tout les processus de rendu visuel, permet de passer de matrices 2D à une image au format `png` (non paralélisé)
- **`perlin`** : permet la création de [bruit de perlin](https://fr.wikipedia.org/wiki/Bruit_de_Perlin)  sur des matrices crées à l'avance (code obtenu sur wikipédia)
- **`main`** : synchronise les goroutines des packages précédents et gère la parallélisation.

## Stratégie

### Contraintes

La contrainte principale repose dans l'assemblage des cartes générée parallèlement : comment limiter les incohérences aux frontières entre les cartes créées indépendamment et assemblée ensuite ? 

On a d'abord cherché à créer des bords de cartes connus, pour ensuite générer des bruits à partir de ces bords, mais on s'est rendu compte que l'algorithme de génération de bruit dont nous disposions n'était pas adapté à cette logique. En effet, le processus de génération de relief associé était complètement indépendant des reliefs déjà présents sur une carte, et ne pouvais pas tenir compte de bords déjà existant. On a donc décidé de générer des cartes indépendantes, de les assembler, puis d'appliquer une moyennage sur les bords des cartes, pour limiter les incohérences. 

Cette solution n'est pas optimale en fonction des reliefs généré par bruit de perlin, mais permet d'apporter une solution satisfaisante compt tenu du temps et des ressources dont on dispose. 

### Parallélisme

Ainsi, plusieurs processus sont parallélisés

- `matrix.InitMatrice(MAPSIZE)` et `perlin.GeneratePerlin(m, out)` effectuent parallèlement la génération de bruit de perlin sur des matrices indépendantes, stockées ensuite dans une liste de matrices.
- `matrix.AjouterParcelle(src, MAPSIZE, x, y, FINALMAP)` ajoute simultanément les matrices élémentaires sur la matrice finale, aux coordonnées x et y adéquates.
- `matrix.AvgOnColumn(FINALMAP, y, AVGWIDE)` et `matrix.AvgOnLine(FINALMAP, x, AVGWIDE)` effectuent les moyennages sur les lignes puis sur les colonnes. Pour éviter les conflits de processus modifiant la même ressource, on effectue d'abord simultanément les moyannages en ligne, puis en colonne. 

## Utilisation

### Installation 

- OS : Linux

Récupérer le répertoire du projet.
```bash
git clone https://github.com/eliemine24/go_project
```

Exécuter le script `setup.sh` pour initialiser le répertoire Go et instaler les bibliothèques nécessaires.
```bash
./setup.sh
```

### Paramètres

Plusieurs paramètres pour la génration d'une carte. Dans `main.go`, on peut modifier le nombre de carte générérées paralèllement et leur taille :
```go
const (
	MAPSIZE      = 100             // taille des maps élémentaires
	RATIO        = 16              // nombre de maps élémentaires sur un coté de map
	FINALMAPSIZE = MAPSIZE * RATIO // taille de la map finale
	NBMAPS       = RATIO * RATIO   // nombre de maps élémentaires sur la map finale
	AVGWIDE      = 40              // largeur du moyennage
	AVGNUMBER    = 10              //nombre de moyennages successifs
)
```

`AVGWIDE` et `AVGNUMBER` permette de régler respectivement la largeur sur laquelle le moyennage est effectué sur les bords de maps et le nombre de moyennages successif effectués. Attention, les moyennage s'exécutent en parallèle sur une seule matrice, pour éviter que les processus ne se chevauchent, ne pas dépasser la moitié de `MAPSIZE` pour `AVGWIDE`. 

---

Paramètres modifiables dans `perlin/perlin.go`

```go
const (
	SCALE         = 50  // Larger SCALE = noise varies slower = larger features (smoother, more zoomed-out)
	POIDS_MOYENNE = 0.5 // Higher weight = output follows neighbors more = smoother transitions between cells
	POIDS_PERLIN  = 0.5 // Higher weight = more randomness, less smoothing
)
```

- `SCALE` : permet de régeler la taille des reliefs. une plus grande valeur permet la création de reliefs plus larges.
- `POIDS_MOYENNE` : influence la rapidité de changement du relief. un poids plus grand donne un relief variant plus rapidement
- `POIDS_PERLIN` : influence la randomness du relief. un poids plus grand diminue la "cohérenceé du relief"

---

On obtient finalement un carte `perlin_heatmap.png` représentant le relief selon une échelle de couleur. Attention, des test sur nos machines nous ont montré qu'une trop grande définition d'image pouvait faire planter le programme. Cette limite varie selon les machine, et les paramètres par défaut ne doivent normalement pas poser problème. 
