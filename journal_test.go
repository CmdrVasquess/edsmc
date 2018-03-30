package edsm

import (
	"testing"
)

var jumpTemplates = []string{
	`{ "timestamp":"2018-05-13T10:26:33Z", "event":"FSDJump",
	   "StarSystem":"Col 285 Sector RT-E b13-3", "SystemAddress":7269098268041, "StarPos":[76.59375,81.15625,-69.34375],
	   "SystemAllegiance":"",
	   "SystemEconomy":"$economy_None;",
	   "SystemEconomy_Localised":"None",
	   "SystemSecondEconomy":"$economy_None;",
	   "SystemSecondEconomy_Localised":"None",
	   "SystemGovernment":"$government_None;",
	   "SystemGovernment_Localised":"None",
	   "SystemSecurity":"$GAlAXY_MAP_INFO_state_anarchy;",
	   "SystemSecurity_Localised":"Anarchy",
	   "Population":0,
	   "JumpDist":33.784,
	   "FuelUsed":2.969149,
	   "FuelLevel":13.030851 }`,
	`{ "timestamp":"2018-05-13T10:27:21Z", "event":"FSDJump",
	   "StarSystem":"Col 285 Sector OY-E b13-4", "SystemAddress":9467853153673, "StarPos":[59.46875,106.56250,-68.56250],
	   "SystemAllegiance":"",
	   "SystemEconomy":"$economy_None;",
	   "SystemEconomy_Localised":"None",
	   "SystemSecondEconomy":"$economy_None;",
	   "SystemSecondEconomy_Localised":"None",
	   "SystemGovernment":"$government_None;",
	   "SystemGovernment_Localised":"None",
	   "SystemSecurity":"$GAlAXY_MAP_INFO_state_anarchy;",
	   "SystemSecurity_Localised":"Anarchy",
	   "Population":0,
	   "JumpDist":30.649,
	   "FuelUsed":2.373660,
	   "FuelLevel":13.626340 }`,
}

func TestFsdJump(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	//srv := NewService(Test)
	// TODO …
}
