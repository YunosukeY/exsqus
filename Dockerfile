FROM golang:1.20.5 AS builder
WORKDIR /work
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build cmd/main.go

FROM gcr.io/distroless/static-debian11
COPY --from=builder /work/main /usr/local/bin/
USER nonroot
