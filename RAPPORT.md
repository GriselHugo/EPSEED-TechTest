# Rapport de projet

## Sommaire :

- [Présentation du projet](#présentation-du-projet)
- [Prérequis](#prérerquis)
- [Lancement du projet](#lancement-du-projet)
  - [Problèmes possibles](#problèmes-possibles)
- [Technologies utilisées](#technologies-utilisées)
    - [API REST](#api-rest)
    - [Application web](#application-web)
    - [Base de données](#base-de-données)
- [Points d'amélioration](#points-damélioration)
    - [API REST](#api-rest-1)
    - [Application web](#application-web-1)
- [Conclusion](#conclusion)

## Présentation du projet :

Ce projet est constitué de 3 parties :
- Une API REST
- Une application web
- Une base de données

## Prérerquis :

Pour lancer le projet, vous devrez avoir installé :
- Docker
- Docker-compose

## Lancement du projet :
Pour lancer le projet, il suffit de lancer la commande suivante à la racine du projet :
```bash
docker compose up
```

### Problèmes possibles :
Du fait du temps de lancement de la base de données, il est possible que l'API se lance avant que la base de données soit prête. Vous rencontrerez alors une erreur de ce type :
```bash
ett-api  | Attempting to connect to database...
ett-api  | [error] failed to initialize database, got error dial tcp 172.26.0.2:3306: connect: connection refused
ett-api  | Error: dial tcp 172.26.0.2:3306: connect: connection refused
ett-api exited with code 0
```
Dans ce cas, il suffit de relancer l'API avec la commande suivante :
```bash
docker compose down
docker compose build
docker compose up
```
Ou sinon dans un autre terminal :
```bash
docker compose up go-api
```

## Technologies utilisées :

### API REST :
L'API REST est codée en Go. Elle utilise l'ORM GORM pour communiquer avec la base de données.
Cette technologie a été choisie pour plusieurs raisons :
- La décourverte du langage Go, langage utilisé au sein d'EPSEED
- La rapidité de développement grâce à l'ORM GORM
- La facilité de déploiement grâce à la compilation du code en un seul binaire

### Application web :
L'application web est codée avec le framework ReactJS.
Cette technologie a été choisie pour plusieurs raisons :
- La facilité de déploiment d'une application NodeJS
- ReactJS est une technologie simple à prendre en main

### Base de données :
La base de données est une base de données MariaDB.
Cette technologie a été choisie pour une raison simple :
- C'est la base de données recommandée dans le sujet.

## Points d'amélioration :

### API REST :
Du fait de la découverte du langage Go, il est possible que l'architecture de l'API ne soit pas optimale. Il serait aussi possible d'ajouter des tests unitaires pour tester le bon fonctionnement de l'API. Enfin, il serait possible d'ajouter des fonctionnalités comme:
- La possibilité de partager des notes entre utilisateurs
- Ajouter du contenu multimédia dans les notes
- Pouvoir modifier/supprimer son compte.

### Application web :
L'application web est extrêmement basique. Il serait possible de modfier l'application en lui rajoutant :
- Une page avec son profil d'utilisateur
- Un meilleur design
- Une meilleure accessibilité

## Conclusion :
J'ai beaucoup apprécié réaliser ce projet. J'ai pu découvrir le Go et de me perfectionner dans les autres technologies utilisées.
