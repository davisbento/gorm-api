FROM golang:latest as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build web/main.go

FROM alpine:latest  

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 4000 

CMD ["./main"]