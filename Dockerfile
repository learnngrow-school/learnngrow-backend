FROM golang:1.23.2-alpine

WORKDIR /app

COPY go.mod go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

EXPOSE 8080

CMD ["./main"]

