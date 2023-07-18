# syntax=docker/dockerfile:1
FROM golang

RUN go version

WORKDIR /app
COPY ./ ./

RUN go mod download
RUN go build -o crm-admin ./cmd/main.go
RUN chmod 777 crm-admin

CMD ["./crm-admin"]