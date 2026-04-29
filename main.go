package main

import (
	"context"
	"fmt"
	"os"

	"github.com/bitrise-io/go-steputils/v2/stepconf"
	"github.com/bitrise-io/go-utils/v2/env"
	"github.com/bitrise-io/go-utils/v2/log"

	mirrorspkg "github.com/bitrise-io/bitrise-build-cache-cli/v2/pkg/gradle/mirrors"
)

type Input struct {
	Mavencentral       bool `env:"mavencentral,required"`
	MavencentralApache bool `env:"mavencentral_apache,required"`
	Google             bool `env:"google,required"`
	Verbose            bool `env:"verbose,required"`
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

	selected := selectedFlags(input)

	activator := mirrorspkg.NewActivator(mirrorspkg.ActivatorParams{
		SelectedFlags: selected,
		DebugLogging:  input.Verbose,
		Logger:        logger,
	})

	if err := activator.Activate(context.Background()); err != nil {
		return fmt.Errorf("activate gradle mirrors: %w", err)
	}

	return nil
}

func selectedFlags(input Input) []string {
	var selected []string

	if input.Mavencentral {
		selected = append(selected, "mavencentral")
	}

	if input.MavencentralApache {
		selected = append(selected, "mavencentral-apache")
	}

	if input.Google {
		selected = append(selected, "google")
	}

	return selected
}
