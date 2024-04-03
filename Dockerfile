FROM golang:latest

WORKDIR /go/src/app

COPY go.sum .
COPY go.mod .
RUN go mod download

COPY cmd/ cmd/
COPY internal/ internal/
COPY pkg/ pkg/
RUN go build -o /go/bin/app ./cmd/app

CMD ["/go/bin/app"]
