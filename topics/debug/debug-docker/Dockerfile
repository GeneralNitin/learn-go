FROM golang:1.23-alpine AS builder

WORKDIR /build
COPY . .

RUN CGO_ENABLED=0 go build -gcflags "all=-N -l" -o app .
RUN CGO_ENABLED=0 go install -ldflags "-s -w -extldflags '-static'" github.com/go-delve/delve/cmd/dlv@latest

FROM apline
WORKDIR /app
COPY --from=0 /build/app .
COPY --from=0 /go/bin/dlv .
CMD ["./dlv", "--listen:4040", "--continue", "--accept-multiclient", "--headless=true", "--api-version=2", "exec", "./app"]