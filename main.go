package main

import (
	"context"
	"fmt"
	"os"

	"github.com/bitrise-io/go-steputils/v2/stepconf"
	"github.com/bitrise-io/go-utils/v2/env"
	"github.com/bitrise-io/go-utils/v2/log"

	mirrorspkg "github.com/bitrise-io/bitrise-build-cache-cli/v3/pkg/gradle/mirrors"
)

type Input struct {
	Verbose bool `env:"verbose,required"`
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	logger := log.NewLogger()
	envRepo := env.NewRepository()
	parser := stepconf.NewInputParser(envRepo)

	var input Input
	if err := parser.Parse(&input); err != nil {
		return fmt.Errorf("parse inputs: %w", err)
	}

	logger.EnableDebugLog(input.Verbose)
	stepconf.Print(input)
	logger.Println()

	activator := mirrorspkg.NewActivator(mirrorspkg.ActivatorParams{
		DebugLogging: input.Verbose,
		Logger:       logger,
	})

	if err := activator.Activate(context.Background()); err != nil {
		return fmt.Errorf("activate gradle mirrors: %w", err)
	}

	return nil
}
