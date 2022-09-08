FROM golang:1.19-alpine

RUN apk add --update tzdata \
    bash wget curl git;

RUN mkdir -p $$GOPATH/bin && \
    curl https://glide.sh/get | sh && \
    go get github.com/pilu/fresh

ENV TZ=Asia/Singapore

ADD . /go/src/game-service
WORKDIR /go/src/game-service

CMD glide update && fresh -c runner.conf main.go

# need to expose ports
# need tern migration