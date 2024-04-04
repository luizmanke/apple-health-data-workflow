FROM golang:latest

WORKDIR /go/src

COPY go.sum .
COPY go.mod .
RUN go mod download

COPY cmd/ cmd/
COPY internal/ internal/
COPY pkg/ pkg/

RUN go build -o /go/bin/backend ./cmd/backend
RUN go build -o /go/bin/frontend ./cmd/frontend
