# Génération procédurale de carte avec goroutines (Go)
## Objectif du projet
L’objectif de ce projet est de générer une carte aléatoire en la divisant en plusieurs régions calculées en parallèle, chaque région étant générée par une goroutine.
Un point clé du projet est l’évaluation des performances :
- montrer à partir de quelle taille de carte la génération parallèle (une région par goroutine) devient plus efficace qu’une génération séquentielle.

## Problématique principale
La difficulté majeure réside dans la cohérence entre les régions :
- Éviter les ruptures visibles (failles, discontinuités) aux frontières
- Assurer une continuité du relief entre régions adjacentes
- Permettre une reconstruction globale correcte de la carte finale

## Contraintes et caractéristiques
### Carte
Carte carrée
Taille variable (à tester pour l’analyse de performance)
Nombre de régions identique en ligne et en colonne
→ régions elles-mêmes carrées
Exemple initial : 9 régions (3×3), puis augmentation progressive

### Données
Carte d’élévation (relief) dans un premier temps
Représentation matricielle
Valeurs normalisées entre 0 et 1
0 : point le plus bas
1 : point le plus haut
Représentation 2D
Nuance de couleur dépend de l’élévation

## Évolutions possibles
Si la première version est fonctionnelle :
- Différencier plusieurs types de terrain :
Eau (rivières, lacs) → bleu
Prairie → vert
Forêt → vert foncé / marron
- Extension vers une représentation 4D

## Approches envisagées
- Approche 1 : Génération par quadrillage (validée par le PFR)
Générer les points d’intersection entre 4 régions
Générer ensuite les lignes du quadrillage
Générer enfin les régions en se basant sur ces points et lignes
Cette approche garantit une continuité structurelle entre régions.
- Autres idées explorées
Placer tous les points d’intersection à un même niveau mais peu réaliste visuellement
Générer les régions indépendamment puis tenter de les « recoller » mais complexe et peu robuste

## Méthode de génération (première version)
Le premier point est complètement aléatoire
Génération de la première ligne et de la première colonne à partir de ce point
Chaque point (x, y) est calculé à partir : du point au-dessus (x, y-1) et du point à gauche (x-1, y)
Valeur obtenue par une moyenne, avec un dénivelé :
- positif ou négatif
- continu sur un minimum de cases
Problème rencontré : apparition de failles visibles
Implémentation initiale en Python, peu adaptée au projet final

## Approche 2 : Bruit de Perlin
L’utilisation du bruit de Perlin permet :
- Une meilleure continuité
- Un rendu plus naturel du relief
Implémentation testée à partir de code Wikipedia (Python).

### Problématique
Gestion correcte d’une région centrale avec 4 bords imposés
Assurer la compatibilité du bruit entre régions adjacentes

## Parallélisation et performances
### Stratégie
Créer une matrice globale représentant la carte finale
Chaque goroutine écrit directement dans sa zone dédiée de la matrice
Éviter :
- la génération d’images par région
- puis leur assemblage (coût inutile)

## Sécurité concurrente
Vérifier que l’écriture concurrente dans la matrice (ou image) ne provoque pas de panic en Go
Chaque goroutine écrit dans une zone mémoire distincte → thread-safe par conception

## Mesure des performances
Le temps mesuré correspond uniquement à la génération de la matrice finale
La transcription en PNG est exclue des mesures

## Technologies envisagées
Langage : Go
Parallélisme : goroutines
Synchronisation : WaitGroup
Génération procédurale : bruit de Perlin

Analyse de performance : benchmark en fonction de la taille de la carte
