package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli"
)

var app = cli.NewApp()

var pizza = []string{"Enjoy your pizza with some delicious"}

func info() {
	app.Name = "Simple pizza CLI"
	app.Usage = "An example CLI for ordering pizza's"
	app.Authors = []*cli.Author{
		&cli.Author{
			Name: "Jimmy",
		},
	}
	app.Version = "1.0.0"
}

func commands() {
	app.Commands = []*cli.Command{
		{
			Name:    "peppers",
			Aliases: []string{"p"},
			Usage:   "Add peppers to your pizza",
			Action: func(c *cli.Context) error {
				pe := "peppers"
				peppers := append(pizza, pe)
				m := strings.Join(peppers, " ")
				fmt.Println(m)
				return nil
			},
		},
	}
}

func main() {
	info()
	commands()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
