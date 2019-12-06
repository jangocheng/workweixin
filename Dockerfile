FROM golang:1.13 as gobase

ENV GO111MODULE on

WORKDIR /srv

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o main main.go

CMD ["./main"]