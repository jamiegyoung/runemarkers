package main

import (
	"github.com/jamiegyoung/runemarkers-go/internal/builder"
	"github.com/jamiegyoung/runemarkers-go/internal/server"
)

func main() {
	builder.Build(false)
	server.Start()
}
