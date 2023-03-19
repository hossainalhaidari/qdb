FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

ENV CGO_ENABLED=1
COPY . .
RUN go build -v -o /app

EXPOSE 1323
CMD ["./qdb"]