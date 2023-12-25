package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var Commands = []*cli.Command{
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
	{
		Name:      "service",
		Aliases:   []string{"svc"},
		Usage:     "Creates a Service manifest",
		ArgsUsage: " ",
		Action: func(cCtx *cli.Context) error {
			fmt.Println(ServiceManifest)
			return nil
		},
	},
}
