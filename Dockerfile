FROM golang:1.21-alpine AS builder

WORKDIR /app
RUN apk add --no-cache git make

COPY go.mod go.sum /
RUN go mod download

COPY . .
RUN go build -o ./bin/wex ./cmd/wex

# runner

FROM alpine as runner
WORKDIR /app

COPY --from=builder /app/bin/wex .

EXPOSE 8080

ENTRYPOINT /app/wex 
