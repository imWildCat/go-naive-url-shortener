FROM golang:1.22 as builder

WORKDIR /app

COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o url-shortener .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/url-shortener .

EXPOSE 8080

CMD ["./url-shortener"]