package main

import (
	"log"
	"os"
	"time"

	"github.com/alexander-scott/onboard/lib"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:     "Onboard",
		Version:  "v0.0.1",
		Compiled: time.Now(),
		Authors: []*cli.Author{
			{
				Name:  "Alexander Scott",
				Email: "alexanderscott@outlook.com",
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "init",
				Aliases: []string{"i"},
				Usage:   "Initialise an instance of a plan from a pre-defined plan.",
				Action: func(c *cli.Context) error {
					// fmt.Println("added task: ", c.Args().First())
					return lib.Run()
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
