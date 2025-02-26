package main

import (
	"github.com/jamiegyoung/runemarkers/internal/builder"
	"github.com/jamiegyoung/runemarkers/internal/server"
)

func main() {
	builder.Build(false)
	server.Start()
}
