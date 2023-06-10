FROM --platform=$BUILDPLATFORM golang:1.20.5 AS builder
WORKDIR /work
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOARCH=${TARGETARCH} go build cmd/app/main.go

# hadolint ignore=DL3006
FROM gcr.io/distroless/static-debian11
COPY --from=builder /work/main /usr/local/bin/
ENTRYPOINT ["main"]
