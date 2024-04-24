package main

import (
	"os"

	"github.com/jamiegyoung/runemarkers-go/args"
	"github.com/jamiegyoung/runemarkers-go/builder"
	"github.com/jamiegyoung/runemarkers-go/server"
)

func main() {

	argsWithoutProg := os.Args[1:]
	skipThumbs := args.HasArg(argsWithoutProg, "--skip-thumbs") || args.HasArg(argsWithoutProg, "-st")

	builder.Build(skipThumbs)

	if args.HasArg(argsWithoutProg, "-s") ||
		args.HasArg(argsWithoutProg, "--server") {

		server.Start()
	}
}
