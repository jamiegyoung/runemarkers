FROM docker.io/golang:1.24

WORKDIR /usr/src/app

# allows git to read history
RUN git config --system --add safe.directory /usr/src/app

COPY . .

RUN go mod download &&\
go mod verify &&\
go build -v -o /usr/local/bin/server ./cmd/server

CMD ["server"]

