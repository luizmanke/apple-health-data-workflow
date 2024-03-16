FROM golang:latest AS build

WORKDIR /go/src/app

COPY go.mod .
RUN go mod download

COPY cmd/ cmd/
COPY internal/ internal/
RUN go build -o /go/bin/app ./cmd/app

FROM golang:latest AS app

WORKDIR /go/bin

COPY --from=build /go/bin/app /go/bin/app

CMD ["./app"]
