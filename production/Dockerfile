FROM golang:1.23.0 AS build
WORKDIR /go/src/app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/server

FROM debian:stable-slim AS app
RUN apt-get update && apt-get install -y \
    ca-certificates \
    gcc \
    libc6-dev \
    && rm -rf /var/lib/apt/lists/*
WORKDIR /app
EXPOSE 8000
COPY --from=build /go/src/app/app .
COPY --from=build /go/src/app/configs /app/configs
CMD ["./app"]