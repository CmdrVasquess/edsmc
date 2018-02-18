package edsm

import (
	"testing"
)

func TestSystem(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	srv := NewService()
	sys := srv.System("Beagle Point")
	if sys == nil {
		t.Fatal("cannot find Beagle Point in EDSM")
	} else {
		t.Log(sys)
	}
}
