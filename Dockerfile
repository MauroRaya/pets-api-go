FROM golang:1.23.4

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY * ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /pets-api

EXPOSE 80

CMD ["/pets-api"]