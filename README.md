# Wilo API

Dépôt Back-End de l'application Wilo.

## Technologies utilisées

- ![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
- ![MongoDB](https://img.shields.io/badge/MongoDB-%234ea94b.svg?style=for-the-badge&logo=mongodb&logoColor=white)

## Organisation du dépôt ⚠️

L'organisation des branches du dépôt est structurée pour faciliter le développement, les tests, et le déploiement en production. Voici les principales branches utilisées :

- **`main`** (vous vous trouvez ici) :
  Point d'entrée du projet, contient tous les documents nécessaires à la compréhension du projet. 
- **`dev`** :  
  La branche principale de développement continu. Elle sert d'environnement bac à sable pour les développeurs où toutes les nouvelles fonctionnalités et corrections de bugs sont intégrées après validation initiale.

- **`pre-prod`** :  
  Cette branche est destinée à présenter les fonctionnalités au client. Une fois que les développements de la branche `dev` sont stabilisés et validés, ils sont fusionnés dans cette branche pour des démonstrations.

- **`prod`** :  
  La branche finale de production qui contient la version stable et prête à être déployée de l'application. Elle est mise à jour uniquement lorsque les changements dans `pre-prod` sont entièrement validés.

- **`<developerName>-dev`** :  
  Chaque développeur dispose de sa propre branche. Ces branches sont fusionnées dans la branche `dev` une fois les développements terminés et validés.

**Bonnes pratiques :**

- Tester les fonctionnalités dans la branche `dev` avant de les intégrer dans `pre-prod`.
- Ne jamais effectuer de développement direct sur les branches `pre-prod` et `prod`.
- Maintenir la branche `prod` uniquement avec du code stable et prêt pour les utilisateurs finaux.

## Lancer le projet

### Installation

```shell
git clone https://github.com/StevenYAMBOS/wilo-api
cd wilo-api
# Une fois dans le projet, basculez sur la branche `dev` pour toute modification :
git checkout dev
```

### Variables d'environnements

Créer un fichier `.env` à la racine du projet (au même niveau que le fichier `go.mod`) et ajouter les informations suivantes :

```shell
PORT=  # Port sur lequel l'API écoute
MONGO_URI=  # Connexion MongoDB
DATABASE_NAME=  # Nom de la base de données MongoDB
JWT_SECRET=  # Clé secrète pour la génération des JWT
```

### Lancement

```bash
go run cmd/main.go
```

## Liens utiles

- [Dépôt Front-End](https://github.com/StevenYAMBOS/wilo-frontend) *(en cours de développement)*
- [Documentation du projet](./documentation.md)
- [Matrices de décision](https://github.com/StevenYAMBOS/wilo-api/blob/main/Matrices%20de%20d%C3%A9cision.md)
- [Site internet](#) *(à venir)*

## Auteur

### Steven YAMBOS

<a href="https://github.com/StevenYAMBOS"><img src="https://cdn-icons-png.flaticon.com/512/25/25231.png" width="30px" alt="" /><a/>
<a href="https://x.com/StevenYambos">
<img src="https://upload.wikimedia.org/wikipedia/commons/5/53/X_logo_2023_original.svg" width="30px" alt="X Steven YAMBOS"/>
</a>
