package main

import (
	"os"

	"github.com/jamiegyoung/runemarkers-go/internal/args"
	"github.com/jamiegyoung/runemarkers-go/internal/builder"
)

func main() {
	argsWithoutProg := os.Args[1:]
	skipThumbs := args.HasArg(argsWithoutProg, "--skip-thumbs") || args.HasArg(argsWithoutProg, "-st")

	builder.Build(skipThumbs)
}
