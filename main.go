package main

import (
	"os"

	"github.com/jamiegyoung/runemarkers-go/args"
	"github.com/jamiegyoung/runemarkers-go/builder"
	"github.com/jamiegyoung/runemarkers-go/server"
)

func main() {

	argsWithoutProg := os.Args[1:]
	skip_thumbs := args.HasArg(argsWithoutProg, "--skip-thumbs") || args.HasArg(argsWithoutProg, "-st")

	builder.Build(skip_thumbs)

	if args.HasArg(argsWithoutProg, "-s") ||
		args.HasArg(argsWithoutProg, "--server") {

		server.Start()
	}
}
