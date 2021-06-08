package example_test

import (
	"testing"

	"github.com/driftprogramming/go-cache/example"
)

func TestExample(t *testing.T) {
	t.Parallel()
	example.StartApp()
}
