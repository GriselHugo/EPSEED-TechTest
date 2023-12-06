# Interview Fullstack

## Sommaire

- [Introduction](#introduction)
- [Consigne](#consigne)
- [Sujet](#sujet)
  - [Le backend](#le-backend)
  - [La base de données](#la-base-de-données)
- [Les critères d'évaluation](#les-critères-dévaluation)
- [Bonus](#bonus)

## Introduction

Bonjour,
Si vous effectuez ce test, c'est que vous avez postulé pour un poste de développeur chez [Epseed](https://epseed.com).
L'objectif grâce à cet exercice est de pouvoir estimer vos compétences techniques et votre niveau de maîtrise des différents outils.

## Consigne

Ce test technique est à réaliser d'ici la fin de la semaine.

L'objectif est de réaliser un backend contenant une API REST et une base de données.

Nous évaluerons votre capacité à réaliser une API et le schéma de votre base de données.

Pour le rendu de votre projet, vous devrez nous fournir un lien vers un repository git public (GitHub, GitLab, Bitbucket, ...) contenant votre code, un fichier RAPPORT.md expliquant vos choix (sur les langages utilisés, par exemple) et comment lancer votre projet, ainsi qu'un export de votre base de données avec les différentes tables et données stockées dedans.

Vous pouvez aussi nous fournir un fichier Docker ou docker-compose.yml afin de faciliter les tests de votre projet. Cela sera grandement apprécié.

### Sujet

**Faites un Backend pour une application de note.**

Créez une API permettant à une application de pouvoir :
- Créer un utilisateur avec un pseudo/email/mot de passe.
- Se connecter à l'aide d'un login (email/mot de passe).
- Créer, modifier, supprimer et récupérer des notes.

Les notes doivent avoir un titre, un contenu, une date de création et une date de modification.

### Le backend

Le backend pourra être réalisé en utilisant le langage et les packages de votre choix, mais ils devront être expliqués dans le fichier RAPPORT.md.

L'utilisation d'un ORM est recommandée mais pas obligatoire.

### La base de données

La base de données devra être une base de données SQL, de préférence [MariaDB](https://mariadb.org).

Un docker-compose vous est fourni pour le test.

## Les critères d'évaluation

Les principaux critères d'évaluation seront les suivants :

- La qualité du code
- Le schéma de la base de données
- La qualité du rendu

## Bonus

Vous pouvez ajouter d'autres fonctionnalités comme l'ajout de contenu multimédia dans les notes ou le partage de note entre les utilisateurs.
Vous pouvez aussi réaliser des tests unitaires ou end-to-end.
Si vous en avez envie, vous pouvez réaliser une application web afin de visualiser les notes, cela sera récompensé.
N'hésitez pas à noter dans le RAPPORT.md les bonus que vous avez effectués.