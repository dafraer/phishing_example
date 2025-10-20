FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o service .

CMD ["sh", "-c", "./service $PORT"]