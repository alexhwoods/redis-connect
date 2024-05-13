# syntax=docker/dockerfile:1

FROM golang:1.22

WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /hello-redis

# Run
CMD ["/hello-redis"]