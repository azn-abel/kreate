package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

type CustomError struct {
	Message string
}

func (e CustomError) Error() string {
	return e.Message
}

func main() {

	app := &cli.App{
		Name:                   "kreate",
		Usage:                  "Easily create Kubernetes manifest files",
		ArgsUsage:              " ",
		EnableBashCompletion:   true,
		UseShortOptionHandling: true,
		Commands:               Commands,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}

}
