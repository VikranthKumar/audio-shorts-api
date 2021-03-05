FROM golang:latest

WORKDIR /go/src/nooble/task/audio-shorts-api
COPY . /go/src/nooble/task/audio-shorts-api
RUN go mod download