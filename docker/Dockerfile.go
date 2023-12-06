# Utilisez l'image officielle de Go comme image de base
FROM golang:latest

# Définissez le répertoire de travail à l'intérieur du conteneur
WORKDIR /app

# Copiez le fichier go.mod et go.sum dans le répertoire de travail
COPY server/go.mod server/go.sum ./

# Téléchargez les dépendances du projet
RUN go mod download

# Copiez tout le contenu du répertoire actuel dans le répertoire de travail du conteneur
COPY ./server .

CMD ["sh", "-c", "go build -o server cmd/server/main.go && ./server"]