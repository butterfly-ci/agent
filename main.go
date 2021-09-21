package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/butterfly-ci/agent/internal"
	"github.com/butterfly-ci/agent/pkg/run"
	"github.com/urfave/cli/v2"
)

func init() {
	logLevel := internal.GetEnvstring("LOG_LEVEL", log.InfoLevel.String())
	logType := internal.GetEnvstring("LOG_TYPE", "text")
	if logType == "json" {
		log.SetFormatter(&log.JSONFormatter{PrettyPrint: true})
	}
	parsedLogLevel, err := log.ParseLevel(logLevel)

	if err != nil {
		log.Errorf("Bad LOG_LEVEL defined: %v. Defaulting to INFO level", err)
		parsedLogLevel = log.InfoLevel
	}
	log.SetLevel(parsedLogLevel)
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
				log.Debugf("key: %s", c.String("key")) // TODO: don't log this...
				// Run our fun run.go code.
				a := run.NewRun()
				a.Runner()
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
