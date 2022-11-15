FROM golang:1.19-alpine

WORKDIR /game-service-app

COPY . .

RUN go mod download

RUN go install -mod=mod github.com/githubnemo/CompileDaemon

EXPOSE 8030

ENTRYPOINT CompileDaemon -log-prefix=false -polling -build="go build ./cmd/api/" -command="./api"