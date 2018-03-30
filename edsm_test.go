package edsm

import (
	"testing"
)

func TestSystem(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	srv := NewService(Test)
	sys, err := srv.System("Beagle Point", SYSTEM_ALL)
	if err != nil {
		t.Error(err)
	}
	if sys == nil {
		t.Fatal("cannot find Beagle Point in EDSM")
	} else {
		t.Log(sys)
	}
}

func TestDiscard(t *testing.T) {
	if testing.Short() {
		t.Skip("skipped test in short mode")
	}
	srv := NewService(Test)
	discs := []string{"-MUST GO AWAY-"}
	err := srv.Discard(&discs)
	if err != nil {
		t.Error(err)
	}
	if len(discs) == 0 {
		t.Fatal("suspicious: no event to discard")
	}
	for _, e := range discs {
		if e == "-MUST GO AWAY-" {
			t.Error("discard list did not overwrite")
		}
		t.Logf("discard: '%s'\n", e)
	}
}

func TestJournal(t *testing.T) {
	if testing.Short() {
		t.Skip("skipped test in short mode")
	}
	srv := NewService(Test)
	eventStr := `{ "timestamp":"2018-04-29T14:30:52Z", "event":"Cargo", "Inventory":[ { "Name":"drones", "Name_Localised":"Limpet", "Count":4, "Stolen":0 }, { "Name":"grain", "Count":1, "Stolen":0 } ] }`
	err := srv.Journal("goEDSMc-Tester", eventStr)
	if err != nil {
		t.Fatal(err)
	}
}
