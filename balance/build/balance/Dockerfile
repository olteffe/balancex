# Step 1: Modules caching
FROM golang:1.17.5-alpine as modules
WORKDIR /modules
COPY go.mod go.sum /modules/
RUN go mod download

# Step 2: Builder
FROM golang:1.17.5-alpine as builder
COPY --from=modules /go/pkg /go/pkg
WORKDIR /build
COPY . /build
RUN go build -ldflags="-s -w" -o balance ./cmd/balance/main.go

# Step 3: Final
FROM scratch
COPY --from=builder ["/build/balance", "/build/config/config-docker.yml", "/"]
EXPOSE 9000/tcp
USER 1001
ENTRYPOINT ["/balance"]
