FROM golang:1.20.5 AS builder
WORKDIR /work
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build cmd/app/main.go

FROM gcr.io/distroless/static-debian11:nonroot
COPY --from=builder /work/main /usr/local/bin/
ENTRYPOINT ["main"]
