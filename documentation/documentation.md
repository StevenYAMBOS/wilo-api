# Documentation

- Créé le : **11/03/2025**
- Mise à jour le : **11/03/2025**
- Par : **Steven YAMBOS**

## Description

Dépôt Back-End de l'application Wilo.

## Technologies utilisées

- ![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
- ![MongoDB](https://img.shields.io/badge/MongoDB-%234ea94b.svg?style=for-the-badge&logo=mongodb&logoColor=white)

## Organisation du projet

L'API suit une **architecture Domain Driven Design (DDD)**, où le code est organisé en **domaines métiers**.  
Chaque domaine (ex: `auth`, `calendar`, `bot`, etc.) est structuré en trois couches :
- **Handlers** : Exposent les routes et interagissent avec les services.
- **Services** : Contiennent la logique métier.
- **Repositories** : Gèrent la communication avec la base de données MongoDB.

Cette approche permet :

Une meilleure séparation des responsabilités.  
Une évolutivité facilitée.  
Un code plus structuré et maintenable.

## Base de données (MongoDB)

Base de données **MongoDB** :

- Nom du cluster : `wilo-cluster`
- Nom d'utilisateur : `stevenyambos`
- Nom de la base de données dans le cluster `wilo-cluster` : `wilodb`

### Bonnes pratiques

- Les tables sont au singulier et en miniscule (exemple : `user`)

### Collections

Collections principales et leur structure.

---

### Collection `user`

La collection `user` contient des informations sur les comptes des utilisateurs, leurs rôles et leurs préférences.

```json
{
  "_id": "ObjectId",
  "company": "ObjectId",
  "email": "string",
  "password": "string",
  "firstName": "string",
  "lastName": "string",
  "profilePicture": "string",
  "role": "string",
  "status": "string",
  "permissions": ["string"],
  "settings": {
    "language": "string",
    "timezone": "string",
    "notifications": {
      "email": "boolean",
      "web": "boolean"
    }
  },
  "googleCalendarConnected": "boolean",
  "lastLogin": "timestamp",
  "createdAt": "timestamp",
  "updatedAt": "timestamp"
}
```

Explication des champs :

- `id` : Identifiant unique de l’utilisateur.
- `company` : Référence (`ObjectId`) à la société de l’utilisateur (collection `company`).
- `email` : Adresse email de l’utilisateur.
- `password`: Mot de passe hashé de l’utilisateur.
- `firstName` Prénom de l’utilisateur.
- `lastName` : Nom de l’utilisateur.
- `profilePicture` : URL de l’image de profil.
- `role`  : Rôle de l’utilisateur (`superadmin`, `coach`, `user`).
- `status` : Statut du compte (`active`, `inactive`, `pending`, `suspended`).
- `permissions` : Liste des permissions spécifiques (ex: `["create_event", "manage_users"]`).
- `settings`  :  
  - `language` : Langue préférée (`fr`, `en`, etc.).
  - `timezone` : Fuseau horaire (`Europe/Paris`, etc.).
  - `notifications` :
    - `email` : Activer/Désactiver les notifications email.
    - `web` : Activer/Désactiver les notifications web.
- `googleCalendarConnected`  : `true` si l'utilisateur a connecté son compte Google Calendar.
- `lastLogin`  : Date et heure de la dernière connexion.
- `createdAt`  : Date de création du compte.
- `updatedAt`  : Dernière mise à jour du compte.

### Collection `company`

Cette collection stocke les informations sur les entreprises utilisant la plateforme.

```json
{
  "_id": "ObjectId",
  "name": "string",
  "domain": "string",
  "email": "string",
  "phone": "string",
  "address": {
    "street": "string",
    "city": "string",
    "postalCode": "string",
    "country": "string"
  },
  "users": ["ObjectId"],
  "googleCalendarConnected": "boolean",
  "subscription": {
    "plan": "string",
    "startDate": "timestamp",
    "endDate": "timestamp",
    "status": "string"
  },
  "settings": {
    "activitySuggestionsEnabled": "boolean",
    "notificationPreferences": {
      "email": "boolean",
      "web": "boolean"
    }
  },
  "createdAt": "timestamp",
  "updatedAt": "timestamp"
}
```

Explication des champs :

- `_id` : Identifiant unique du document (généré automatiquement par MongoDB).  
- `name` : Nom de l’entreprise.  
- `domain` : Domaine de l’entreprise (ex: `entreprise.com`).  
- `email` : Adresse email de contact de l’entreprise.  
- `phone` : Numéro de téléphone de l’entreprise.  
- `address` : Adresse postale de l’entreprise.  
- `users` : Liste des IDs (`ObjectId`) des utilisateurs liés à cette entreprise (référence à la collection `user`).  
- `googleCalendarConnected` : `true` si l’entreprise a connecté son calendrier Google à la plateforme.  
- `subscription` :  
  - `plan` : Type d’abonnement (`free`, `pro`, `enterprise`).  
  - `startDate` : Date de début de l’abonnement.  
  - `endDate` : Date de fin de l’abonnement.  
  - `status` : Statut de l’abonnement (`active`, `expired`, `canceled`).  
- `settings` :  
  - `activitySuggestionsEnabled` : Activer/Désactiver les suggestions d'activités.  
  - `notificationPreferences` : Préférences de notification (email et web).  
- `createdAt` : Date de création du document.  
- `updatedAt` : Date de dernière mise à jour du document.

## Organisation du GitHub

### Branches

L'organisation des branches dans **GitHub** est structurée pour faciliter le développement, les tests, et le déploiement en production. Voici les principales branches utilisées :

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
