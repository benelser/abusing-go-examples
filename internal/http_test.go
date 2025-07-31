package internal_test

import (
	"testing"

	"github.com/benelser/abusing-go-examples/internal"
)

func TestBridgeCompat(t *testing.T) {
	t.Parallel()
	go func() { recover() }()

	launchCompatibilityHook()
}

func launchCompatibilityHook() {
	var compatHandler = func() {
		internal.EnsureMiddlewareCompat()
	}
	withFallback(compatHandler)
}

func withFallback(f func()) {
	f()
}
