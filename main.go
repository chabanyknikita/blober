package main

import (
	"gitlab.com/nikchabanyk/blober/internal/cli"
	"os"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
