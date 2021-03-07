FROM golang:latest

WORKDIR /go/src/nooble/task/audio-shorts-api
COPY . .
RUN go mod download