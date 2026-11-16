package main

import (
	"fmt"
	"os"

	"github.com/bitrise-io/go-steputils/v2/stepconf"
	"github.com/bitrise-io/go-utils/v2/env"
	"github.com/bitrise-io/go-utils/v2/log"
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

	logger.Warnf("Gradle mirrors are temporarily disabled; this step is a no-op and no mirror was activated.")

	return nil
}
