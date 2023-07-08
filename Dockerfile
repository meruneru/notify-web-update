FROM golang:1.17

WORKDIR /app

COPY . .

RUN go mod tidy
RUN GOOS=linux go build -o main .

CMD ["/app/main"]
