FROM golang:1.13

ENV TZ Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

ENV GO111MODULE on

WORKDIR /srv

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o main main.go

CMD ["./main"]