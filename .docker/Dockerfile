# Step 1
FROM golang:1.23.3-alpine AS builder
RUN apk update && apk add git gcc g++
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo ./cmd/api

# Step 2
FROM alpine
RUN apk update && apk add --no-cache ca-certificates make
WORKDIR /opt

COPY ./Makefile .
COPY --from=builder /app/api .

CMD ["/bin/sh"]
