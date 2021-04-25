FROM golang:1.16.3-alpine3.13 AS build
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o model_manage ./model-manage-backend/main.go


FROM scratch AS server
WORKDIR /app
COPY ./model-manage-backend/prod.env .env
COPY --from=build /app/model_manage .

EXPOSE 8080 8090
ENTRYPOINT ["./model_manage"]
