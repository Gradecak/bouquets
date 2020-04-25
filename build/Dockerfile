FROM golang:1.14.0 AS builder

WORKDIR /go/src/github.com/gradecak/bouquets
COPY . .
RUN  go mod download; \
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s"  -o ./app ./cmd/bouquets;


FROM scratch

COPY --from=builder /go/src/github.com/gradecak/bouquets/app /app
ENTRYPOINT ["/app"]
