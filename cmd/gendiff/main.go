package main

import (
	"context"
	"github.com/urfave/cli/v3"
	"log"
	"os"
)

func main() {
	cmd := &cli.Command{
		Name:      "gendiff",
		Usage:     "Compares two configuration files and shows a difference.",
		UsageText: "gendiff [global options]",
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
