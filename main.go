package main

import (
	"github.com/bitrise-io/go-utils/v2/log"
)

func main() {
	log.NewLogger().Warnf("Gradle mirrors are temporarily disabled; this step is a no-op and no mirror was activated.")
}
