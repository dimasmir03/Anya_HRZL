FROM golang:alpine as builder
RUN apk update && apk add --no-cache gcc musl-dev
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o bitcoinmonitor ./cmd/bitcoinmonitor

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/bitcoinmonitor .
EXPOSE 8080
ENTRYPOINT ["/app/bitcoinmonitor"]