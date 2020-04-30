package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "my",
		Usage: "Get My Work Done",
		Action: func(c *cli.Context) error {
			fmt.Println("Get My Work Done")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
