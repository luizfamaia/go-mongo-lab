# build
FROM golang:1.22 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o user-api ./cmd/user-api

# run
FROM alpine:3.20
WORKDIR /app
COPY --from=build /app/user-api /usr/local/bin/user-api
EXPOSE 8088
ENV PORT=8088
ENTRYPOINT ["/usr/local/bin/user-api"]
