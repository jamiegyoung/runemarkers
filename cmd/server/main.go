package main

import (
	"github.com/jamiegyoung/runemarkers/internal/builder"
	"github.com/jamiegyoung/runemarkers/internal/server"
	"github.com/jamiegyoung/runemarkers/internal/logger"
)

var log = logger.New("server")

func main() {
	err := builder.Build(false)
	if err != nil {
		log("An error occured: " + err.Error())
		return
	}

	server.Start()
}
