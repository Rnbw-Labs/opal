package window

import (
	"testing"
)

func TestWindowPackage(t *testing.T) {
	window := New(600, 400, "Test Window")
	window.SetVsync(true)
	for window.ShouldRun() {
		window.Draw()
	}
}
