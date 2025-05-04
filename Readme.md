# GIN V2 - rework en go
/project
│
├── /cmd
│   └── main.go             # Point d'entrée de l'application (initialise Gin, setup router, etc.)
│
├── /api
│   ├── /service
│   │   ├── /user
│   │   │   ├── handler.go  # Contient la logique des handlers
│   │   │   ├── repository.go # Interaction avec la DB
│   │   │   └── user.go     # Les structures de l'utilisateur
│   │   └── /auth
│   │       └── auth.go     # Authentification et gestion des mots de passe
│   ├── router.go           # Fichier des routes
│   └── /types
│       └── user.go         # Les types (structs) utilisés dans l'application
│
├── /utils
│   └── utils.go            # Fonction utilitaire (par exemple, pour valider les données)
│
└── /config
    └── config.go   