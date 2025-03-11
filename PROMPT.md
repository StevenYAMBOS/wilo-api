# Go app

```shell
wilo-api/
│── cmd/                     # Point d'entrée de l'application
│   ├── main.go              # Démarrage du serveur Gin
│
│── config/                  # Configuration (fichiers env, etc.)
│   ├── config.go            # Chargement des variables d'environnement
│
│── internal/                # Code métier (logique principale)
│   ├── auth/                # Gestion de l'authentification
│   │   ├── handler.go       # Handlers (routes)
│   │   ├── service.go       # Logique métier
│   │   ├── repository.go    # Interaction avec la base de données
│   │   ├── middleware.go    # Middleware JWT
│   │
│   ├── calendar/            # Intégration avec Google Calendar
│   │   ├── handler.go       
│   │   ├── service.go       
│   │   ├── repository.go    
│   │
│   ├── dashboard/           # Gestion du dashboard utilisateur
│   │   ├── handler.go       
│   │   ├── service.go       
│   │   ├── repository.go    
│   │
│   ├── bot/                 # Notifications et rappels automatiques
│   │   ├── service.go       
│   │   ├── scheduler.go     # Planificateur de tâches
│   │
│   ├── activities/          # Gestion des activités
│   │   ├── handler.go       
│   │   ├── service.go       
│   │   ├── repository.go    
│
│── pkg/                     # Code réutilisable (helpers, utilitaires)
│   ├── jwt/                 # Gestion des JWT
│   ├── hash/                # Hashing des mots de passe
│
│── web/                     # Frontend (SvelteJS)
│   ├── public/              # Assets statiques
│   ├── src/                 # Code source Svelte
│
│── database/                # Initialisation de la base de données
│   ├── migrations/          # Scripts SQL/Migrations
│   ├── database.go                # Connexion à la BDD
│
│── .env                     # Variables d'environnement (ex: clés API)
│── go.mod                   # Fichier Go Modules
│── go.sum                   # Dépendances Go
│── README.md                # Documentation du projet
```

Création du projet

## Ajouts et améliorations :

- Ajout des routes d'authentification :
  - `POST /register` : Inscription d'un utilisateur avec stockage sécurisé du mot de passe (bcrypt).
  - `POST /login` : Connexion et génération d'un token JWT.
  - `GET /dashboard` : Page sécurisée, accessible uniquement avec un token valide.
  - `POST /logout` : Suppression du token côté client (recommandé) + option de réponse API.

- Mise en place de la gestion des tokens JWT :
  - Génération d’un JWT sécurisé après connexion.
  - Validation automatique du token sur les routes protégées via un middleware JWT.
  - Extraction et vérification des permissions utilisateur.

- Ajout de la gestion des utilisateurs avec MongoDB :
  - Création d’une collection `user` avec les champs essentiels (`email`, `password`, `role`, `createdAt`, etc.).
  - Connexion sécurisée à MongoDB via `database.go`.
  - Vérification de l’unicité de l’email à l'inscription.

- 🔧 Middleware d’authentification (`AuthMiddleware`) :
  - Vérifie la présence et la validité du token JWT dans le header `Authorization`.
  - Bloque l'accès aux routes sécurisées si le token est invalide ou expiré.