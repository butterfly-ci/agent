package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/butterfly-ci/agent/run"
	"github.com/urfave/cli/v2"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {
	app := &cli.App{}
	app.UseShortOptionHandling = true
	app.EnableBashCompletion = true
	app.Commands = []*cli.Command{
		{
			Name:  "run",
			Usage: "run a pipeline on this machine",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "key", Aliases: []string{"k"}},
			},
			Action: func(c *cli.Context) error {
				log.Debugf("key:", c.String("key")) // TODO: don't log this...
				// Run our fun run.go code.
				run.Run()
				return nil
			},
		},
		{
			Name:  "daemon",
			Usage: "daemon mode which listens for work from the master",
			Flags: []cli.Flag{
				&cli.BoolFlag{Name: "serve", Aliases: []string{"s"}},
				&cli.BoolFlag{Name: "option", Aliases: []string{"o"}},
				&cli.StringFlag{Name: "message", Aliases: []string{"m"}},
			},
			Action: func(c *cli.Context) error {
				fmt.Println("serve:", c.Bool("serve"))
				fmt.Println("option:", c.Bool("option"))
				fmt.Println("message:", c.String("message"))
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
