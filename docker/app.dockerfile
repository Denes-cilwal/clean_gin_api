FROM golang:alpine

LABEL maintainer = "Dinesh"

WORKDIR /clean_gin_api
ADD . .


RUN apk update
RUN go mod download
RUN go get github.com/githubnemo/CompileDaemon


CMD echo "Building Golang-Gin-WebApp Container.."
ENTRYPOINT CompileDaemon -command="./clean_gin_api"


