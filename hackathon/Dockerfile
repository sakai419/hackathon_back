FROM golang:1.23.0 AS build
WORKDIR /go/src/app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/server

FROM debian:stable-slim as app
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*
WORKDIR /app
EXPOSE 8000
COPY --from=build /go/src/app/app .
CMD ["./app"]