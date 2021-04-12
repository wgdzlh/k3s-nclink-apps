FROM golang:1.16.3-alpine3.13 AS build
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o adapter

FROM scratch
WORKDIR /app
COPY prod.env .env
COPY --from=build /app/adapter .

ENTRYPOINT ["./adapter"]
