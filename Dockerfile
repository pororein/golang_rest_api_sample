FROM golang:1.16

ENV APP_MONGOURL mongodb://mongo:27017

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o app main.go
