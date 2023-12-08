# EPSEED TECHNICAL TEST


### Problèmes rencontrés
```bash
ett-api  | Attempting to connect to database...
ett-api  |
ett-api  | 2023/12/08 08:49:49 /app/internal/db/db.go:21
ett-api  | [error] failed to initialize database, got error dial tcp 172.26.0.2:3306: connect: connection refused
ett-api  | Error: dial tcp 172.26.0.2:3306: connect: connection refused
ett-api  | Erreur: dial tcp 172.26.0.2:3306: connect: connection refused
ett-api exited with code 0
```

Alors

```bash
docker compose down
docker compose build
docker compose up
```

ou sinon dans un autre terminal

```bash
docker compose up go-api
```