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
		UseShortOptionHandling: true,
		Commands: []*cli.Command{
			{
				Name:      "pod",
				Usage:     "Creates a Pod manifest",
				ArgsUsage: " ",
				Action: func(cCtx *cli.Context) error {
					fmt.Println(PodManifest)
					return nil
				},
			},
			{
				Name:      "deployment",
				Aliases:   []string{"deploy"},
				Usage:     "Creates a Deployment manifest",
				ArgsUsage: " ",
				Action: func(cCtx *cli.Context) error {
					fmt.Println(DeploymentManifest)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}

}
