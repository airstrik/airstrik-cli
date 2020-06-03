FROM golang:latest
COPY ./ /app
COPY ../gobase /gobase
WORKDIR /app

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go bunetworksild main.go" --command=./main