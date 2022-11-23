FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY . .

RUN go mod download

RUN go build -o /practice ./cmd/main.go

CMD [ "/practice", "run" ]