package main

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/jhonynet/hlpr/executor"
	"github.com/jhonynet/hlpr/pipeline"
	"github.com/jhonynet/hlpr/processors"
	"github.com/jhonynet/hlpr/utils/logger"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

const (
	VerboseFlag = "verbose"
)

func CreateApp(logger *zap.Logger) *cli.App {
	return &cli.App{
		Name:      "HLPR",
		Usage:     "Run extended data workflows with one tool.",
		UsageText: "hlpr workflow{.json,.yaml}",
		ArgsUsage: "[workflow file]",
		Action:    appHandler(logger),
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    VerboseFlag,
				Aliases: []string{"v"},
				Usage:   "enable verbose log level",
			},
		},
		Authors: []*cli.Author{
			{Name: "jhonynet", Email: "jonipotter@gmail.com"},
		},
	}
}

func appHandler(zapLogger *zap.Logger) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		if !c.Bool(VerboseFlag) {
			zapLogger = zapLogger.WithOptions(
				zap.IncreaseLevel(zap.NewAtomicLevelAt(zap.InfoLevel)),
			)
		}

		ctx := logger.Context(c.Context, zapLogger)

		// first argument should be the file
		file := c.Args().Get(0)

		if file == "" {
			return errors.New("workflow file is not specified")
		}

		bytes, err := ioutil.ReadFile(file)
		if err != nil {
			return fmt.Errorf("cannot read workflow file %s", err)
		}

		pipeline, err := pipeline.FromYaml(bytes)
		if err != nil {
			return err
		}

		return executor.
			NewDefaultExecutor(processors.DefaultRegistry).
			Execute(ctx, pipeline)
	}
}
