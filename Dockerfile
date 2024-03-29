FROM golang:latest

RUN mkdir /app

WORKDIR /app

COPY . .

RUN go mod download

RUN go build ./cmd/web

EXPOSE 8081


CMD ["./start.sh"]