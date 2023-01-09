FROM golang:1.19-alpine

WORKDIR /app

COPY . .

RUN go mod download

#RUN apk update
RUN apk add make
RUN make migrate

RUN go build -o /practice ./cmd/main.go

CMD [ "/practice", "run" ]