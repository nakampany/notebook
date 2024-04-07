FROM golang:1.22.2-alpine3.19 AS go

WORKDIR /src/app

COPY app/go.mod app/go.sum ./
RUN go mod download
COPY app ./

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
RUN go build \
    -o /go/bin/main \
    -ldflags '-s -w'
CMD ["/go/bin/main"]
