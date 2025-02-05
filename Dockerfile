FROM golang:1.23.6 AS builder
WORKDIR /work
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build cmd/app/main.go

# hadolint ignore=DL3006
FROM gcr.io/distroless/static-debian11
COPY --from=builder /work/main /usr/local/bin/
ENTRYPOINT ["main"]
