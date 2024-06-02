FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod download && go mod verify

RUN go build -v -o main .

EXPOSE 8080

CMD ["/app/main"]