FROM golang:1.16.3-alpine3.13 AS build
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o mode_manage ./mode-manage-backend/main.go


FROM scratch AS server
WORKDIR /app
COPY ./mode-manage-backend/prod.env .env
COPY --from=build /app/mode_manage .

EXPOSE 8082
ENTRYPOINT ["./mode_manage"]
