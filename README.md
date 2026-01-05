# Génération procédurale de carte avec goroutines (Go)
## Objectif du projet
L’objectif de ce projet est de générer une carte aléatoire en la divisant en plusieurs régions calculées en parallèle, chaque région étant générée par une goroutine.   
Un point clé du projet est l’évaluation des performances :   
- montrer à partir de quelle taille de carte la génération parallèle (une région par goroutine) devient plus efficace qu’une génération séquentielle.

## Problématique principale
La difficulté majeure réside dans la cohérence entre les régions :   
- Éviter les ruptures visibles (failles, discontinuités) aux frontières en assurant une continuité du relief entre régions adjacentes   
- Permettre une reconstruction globale correcte de la carte finale

## Contraintes et caractéristiques
### Carte
Carte carrée   
Taille variable (à tester pour l’analyse de performance)   
Nombre de régions identique en ligne et en colonne → régions elles-mêmes carrées  
Premiers tests avec 9 régions (3×3), puis augmentation progressive

### Données
Carte d’élévation (relief) dans un premier temps   
Représentation matricielle   
Valeurs normalisées entre 0 et 1 (0 est le point le plus bas et 1 le point le plus haut)   
Représentation 2D → nuance de couleur dépend de l’élévation

## Évolutions possibles
Si la première version est fonctionnelle :   
- Différencier plusieurs types de terrain 
- Tentative de représentation en 3D

## Approches envisagées
- Génération par quadrillage
  - Générer les points d’intersection entre 4 régions
  - Générer ensuite les lignes du quadrillage
  - Générer enfin les régions en se basant sur ces points et lignes    
Cette approche garantit une continuité structurelle entre régions.   
- Placer tous les points d’intersection à un même niveau (mais très visible)
- Générer les régions indépendamment puis moyenner aux intersections pour les lisser (reste visible)

### Méthode de génération (première version)
Génération d'un premier point aléatoirement   
Génération de la première ligne et de la première colonne à partir de ce point   
Chaque point (x, y) est calculé à partir : du point au-dessus (x, y-1) et du point à gauche (x-1, y)   
Valeur obtenue par une moyenne, avec un dénivelé :   
- positif ou négatif   
- continu sur un minimum de cases   
Problème rencontré : apparition de failles visibles    

### Approche 2 : Bruit de Perlin
L’utilisation du bruit de Perlin permet :   
- Une meilleure continuité
- Un rendu plus naturel du relief   

## Parallélisation et performances
### Stratégie
Créer une matrice globale représentant la carte finale   
Chaque goroutine écrit directement dans sa zone dédiée de la matrice   

## Mesure des performances
Le temps mesuré correspond uniquement à la génération de la matrice finale    
La transcription en PNG est exclue des mesures

## Technologies envisagées
Langage : Go    
Parallélisme : goroutines    
Génération procédurale : bruit de Perlin   


# Organisation
## Programmation séquentielle
Créer un premier algorithme de génération de bruit de perlin sur une matrice carrée, et le tester pour différentes tailles de cartes. A partir de cet algorithme, identifier le temps que prend la génération de carte en programmation séquentielle (programme témoin)    
- créer une matrice carrée   
- implémenter algo bruit de perlin   
- afficher l’image créée 

## Programmation parallèle
Diviser la génération de la carte en plusieurs programmes parallèles.    
- génération de cartes plus petites en parallèle ⇒ faire des tests pour comparer la vitesse de calcul en fonction
  - de la taille des cartes
  - du nombre de cartes
- arranger les bords de la carte pour les rendre cohérents
  - première solution : moyennages le long des bords

