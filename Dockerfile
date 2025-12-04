# Build stage
FROM golang:1.23 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN GOOS=linux GOARCH=amd64  go build -o app .

# Final stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
RUN chmod +x ./app
EXPOSE 8080
CMD ["./app"]