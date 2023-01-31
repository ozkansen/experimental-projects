FROM golang:1.19-bullseye as builder

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -ldflags="-s -w"  -o user_login_api_service ./cmd

FROM debian:bullseye

WORKDIR /app
COPY --from=builder /app/user_login_api_service .

ENV DB_URL ""

EXPOSE 8080

CMD ["/app/user_login_api_service"]
