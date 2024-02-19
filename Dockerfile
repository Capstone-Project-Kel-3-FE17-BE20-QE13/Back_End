FROM golang:1.22-alpine

RUN mkdir /app

WORKDIR /app

COPY ./ /app

RUN go mod tidy

RUN go build -o beaapi

CMD ["./beaapi"]