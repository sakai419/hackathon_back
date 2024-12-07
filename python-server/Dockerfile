FROM golang:1.20-alpine AS build
WORKDIR /app
COPY main.go .
RUN go build -o python_server main.go

FROM alpine:latest
RUN apk add --no-cache python3
WORKDIR /app
COPY --from=build /app/python_server .

EXPOSE 9010
CMD ["./python_server"]
