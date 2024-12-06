FROM golang:1.20-alpine AS build
RUN apk add --no-cache gcc libc-dev

WORKDIR /app
COPY main.go .

RUN go build -o gcc_server main.go

FROM alpine:latest
RUN apk add --no-cache gcc libc-dev
WORKDIR /app
COPY --from=build /app/gcc_server .

EXPOSE 9090
CMD ["./gcc_server"]
