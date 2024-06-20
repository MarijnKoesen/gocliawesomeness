package main

import (
	"github.com/marijnkoesen/gocliawesomeness"
	"github.com/urfave/cli/v2"
	"log"
	"log/slog"
	"os"
)

type CmdConfig struct {
	logFormat string
	logLevel  string
}

var config = CmdConfig{}

func main() {
	app := &cli.App{
		Name:  "Example cli apii",
		Usage: "Showing off log format options",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "log-format",
				EnvVars:     []string{"LOG_FORMAT"},
				Value:       "text",
				Usage:       "Either: text or json",
				Destination: &config.logFormat,
			},
			&cli.StringFlag{
				Name:        "log-level",
				EnvVars:     []string{"LOG_LEVEL"},
				Value:       "info",
				Usage:       "Options: debug, info, warn, error",
				Destination: &config.logLevel,
			},
		},
		Before: beforeAction,
		Action: func(context *cli.Context) error {
			return run()
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	logger := slog.With("app", "foobar")
	logger.Info("Starting uptime-checker")

	logger2 := slog.With("app", "tarzar")
	logger2.Error("OK this is an error")
	logger2.Debug("And a debug")

	logger.Info("and an info")

	logger3 := slog.With("service", "api")
	logger3.Error("OK this is an error")
	logger3.Warn("Warning warning!")
	logger3.Debug("And a debug")
	return nil
}

func beforeAction(_ *cli.Context) error {
	options := &gocliawesomeness.Options{
		Format:     gocliawesomeness.Format(config.logFormat),
		Level:      gocliawesomeness.Level(config.logLevel),
		PrefixKeys: []string{"app", "service"},
	}
	slog.SetDefault(gocliawesomeness.NewAwesomeLogger(options))

	return nil
}
