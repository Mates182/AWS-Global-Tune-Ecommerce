FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod main.go ./

RUN go mod tidy

COPY . .

RUN go build -o main main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 80

CMD ["./main"]
