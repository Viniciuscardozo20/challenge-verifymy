FROM golang:alpine

WORKDIR /app

COPY . .

COPY .env ..

RUN go build -o main cmd/main.go

EXPOSE 8082

CMD ["./main"]