FROM golang:1.16.3-alpine3.13 AS build
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o config_dist ./config-distribute/main.go
RUN CGO_ENABLED=0 go build -o adapter ./adapter-simulator/main.go


FROM scratch AS server
WORKDIR /app
COPY ./config-distribute/prod.env .env
COPY --from=build /app/config_dist .

EXPOSE 8082
ENTRYPOINT ["./config_dist"]


FROM scratch AS client
WORKDIR /app
COPY ./adapter-simulator/prod.env .env
COPY --from=build /app/adapter .

ENTRYPOINT ["./adapter"]
