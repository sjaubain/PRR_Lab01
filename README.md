# Programmation Répartie - Laboratoire I

Auteurs: Jobin Simon, Teklehaimanot Robel

## Description des fonctionnalités du logiciel

   Nous souhaitons implémenter un algorithme simple permettant de synchroniser approximativement les horloges locales des tâches d'une application répartie. Comme nous le savons, chaque site d'un environnement réparti possède sa propre horloge système, mais aussi, cette horloge a un décalage et une dérive qui lui est propre. Le but de notre algorithme est de rattraper ce décalage sans pour autant corriger l'horloge du système.
   
   Pour ce faire, nous distinguons 2 horloges. L'horloge système hsys est l'heure retournée par un appel système : cette horloge désigne la minuterie mise à jour par le système d'exploitation d'un site. Sous un système protégé, il faut avoir les privilèges administrateurs pour le modifier et, pour contourner ce problème, une tâche applicative peut interpréter le temps comme la valeur de l'horloge système sur le site où elle réside additionnée à un décalage. 
   
   Dans ce qui suit, le résultat de cette opération est appelé l'horloge locale. Ainsi pour la tâche applicative i, nous avons : hlocale(i) = hsys(site de i) + décalage(i), la synchronisation des horloges revient alors à trouver décalage(i) pour chaque tâche i de telle sorte que hlocale(i) est identique pour toutes les tâches formant l'application répartie.

## Installation 

Cloner le repository git qui se trouve dans le lien qui suit: https://github.com/sjaubain/PRR_Lab01. Pour pouvoir builder correctement, il vous faut cloner ce repository dans le GOPATH.

Les prochaines étapes vous permettront de lancer Master et le/les slaves.

Dans le dossier lab01, pour produire l'exécutable master.go

```
go build master.go
./master

```

Toujours dans le même dossier (bash différent), pour produire l'exécutable slave.go

```
go build slave.go
./slave

```


## Réalisation

Voici ce que le laboratoire doit effectué:

Multicast packet master -> slave

    [ SYNC | ID ] and [ FOLLOW_UP | MASTER_TIME | ID ]
   
Point to point packet slave -> master
    
    [ DELAY_REQUEST | SALVE_ID ]

Point to point packet master -> slave

    [ DELAY_RESPONSE | MASTER_TIME | SLAVE_ID ]
    

Pour réaliser ce laboratoire, nous nous sommes beaucoup inspiré des exemples qui nous ai données durant le cours de **programmation répartie**.  Nous avons séparer notre code en quatre parties. 

- ***master***
- ***slave***
- ***protocol***
- ***clock***

### Master

Comme voulu dans l'énoncé, le master diffuse en multicase à intervalle régulier la SYNC et le FOLLOW UP et en parralèlle, il se met en écoute pour une possibilité de synchronisation avec un slave afin de corriger le délai de transmission.

### Slave

> Fonction masterReader(): Quant à lui, il se met en écoute de toutes connexions multicast du master. Il permet de remettre à jour l'heure de l'esclave par rapport au master. 

> Fonction delayRequest(): Grâce notamment au goroutine, il envoie en parralèle au serveur master les delays request dans le but de corriger les délais de communication engendrés lors de la première étape de la synchronisation.

### Protocol

Cette partie sert a définir toutes les contantes nécessaire pour les configurations du laboratoire.

### Clock

Cette "classe" permet de simuler d'éventuel latence afin de pouvoir tester notre laboratoire.






