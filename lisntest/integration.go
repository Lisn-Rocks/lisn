package lisntest

import (
	"fmt"
	"os"
	"testing"

	"github.com/sharpvik/env-go"
)

var integration bool

func init() {
	_, err := env.TryGet("INTEGRATION")
	integration = err == nil
	file, _ := os.Create("log")
	fmt.Fprintln(file, integration)
}

func SkipIfNotIntegration(t *testing.T) {
	if !integration {
		t.Skip("skipping integration test ...")
	}
}
