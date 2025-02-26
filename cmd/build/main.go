package main

import (
	"fmt"
	"os"

	"github.com/jamiegyoung/runemarkers/internal/args"
	"github.com/jamiegyoung/runemarkers/internal/builder"
	"github.com/jamiegyoung/runemarkers/internal/logger"
)

var log = logger.New("build")

func main() {
	argsWithoutProg := os.Args[1:]
	skipThumbs := args.HasArg(argsWithoutProg, "--skip-thumbs") || args.HasArg(argsWithoutProg, "-st")

	err := builder.Build(skipThumbs)
	if err != nil {
		// FIXME: doesn't show stack trace
		log(fmt.Sprintf("An error occured: %+v", err))
	}

}
