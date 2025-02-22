FROM golang:1.22.3 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o adminauth ./cmd

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/adminauth .
COPY --from=builder /app/.env .env
EXPOSE 8081

CMD ["./adminauth"]