FROM golang:1.19-alpine

WORKDIR /api-gateway-service-app

COPY . .

RUN go mod download

RUN go install -mod=mod github.com/githubnemo/CompileDaemon

EXPOSE 8010

ENTRYPOINT CompileDaemon -log-prefix=false -polling -build="go build ./cmd/api/" -command="./api"