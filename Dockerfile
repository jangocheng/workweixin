FROM golang:1.13

ENV TZ Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

ENV GO111MODULE on

WORKDIR /srv

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o ./bin/appsrv services/appsrv/main.go \
    && go build -ldflags="-s -w" -o ./bin/contactsrv services/contactsrv/main.go \
    && go build -ldflags="-s -w" -o ./bin/todosrv services/todosrv/main.go