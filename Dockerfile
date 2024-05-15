FROM docker.io/golang:1.22

WORKDIR /usr/src/app

COPY . .

RUN go mod download &&\
go mod verify &&\
go build -v -o /usr/local/bin/server ./cmd/server

CMD ["server"]

