# Use golang:1.17-alpine3.14 as build environment
FROM golang:1.17-alpine3.14 AS build

LABEL maintainer="mendes <pedromendes.godev@gmail.com>"

# Set working directory
WORKDIR /app

# Copy source code
COPY . .
COPY .env .

# Download dependencies
RUN go mod download

# Build the binary with some flags
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o go-docker ./cmd

# Use alpine:3.14 as production environment
FROM alpine:3.14

# Set working directory
WORKDIR /app

# Copy binary from the build environment
COPY --from=build /app/go-docker .

# Expose port 3030
EXPOSE 3030

# Run the binary
CMD ["./go-docker"]
