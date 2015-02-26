package client

import (
	"testing"
)

const (
	OUTPUT_1 = `UNIT					LOAD		ACTIVE	SUB	MACHINE
test1.service			-		-	-	-
test2.service			loaded		active	running	cc3a877c.../127.0.0.1
test3.service			loaded		active	exited	f3274266.../127.0.0.1`
)

func AssertStatusParsedAs(t *testing.T, status UnitStatus,
	expectedUnitName, expectedLoad, expectedActive,
	expectedSub, expectedMachine, expectedMachineIP string) {
	if status.Unit != expectedUnitName {
		t.Fatalf("Unexpected unit name: %s", status.Unit)
	}
	if status.Load != expectedLoad {
		t.Fatalf("Unexpected unit name: %s", status.Load)
	}
	if status.Active != expectedActive {
		t.Fatalf("Unexpected unit name: %s", status.Active)
	}
	if status.Sub != expectedSub {
		t.Fatalf("Unexpected unit name: %s", status.Sub)
	}
	if status.Machine != expectedMachine {
		t.Fatalf("Unexpected machine: %s", status.Machine)
	}
	if status.MachineIP() != expectedMachineIP {
		t.Fatalf("Unexpected IP: %s", status.MachineIP())
	}
}

func TestParser__OUTPUT_1(t *testing.T) {
	status, err := parseFleetStatusOutput(OUTPUT_1)
	if err != nil {
		t.Fatalf("%v", err)
	}
	if len(status) != 3 {
		t.Fatalf("Invalid number of status objects returned, expected 7, got: %d", len(status))
	}
	AssertStatusParsedAs(t,
		status[0],
		"test1.service",
		"-",
		"-",
		"-",
		"-",
		"",
	)
	AssertStatusParsedAs(t,
		status[1],
		"test2.service",
		"loaded",
		"active",
		"running",
		"cc3a877c.../127.0.0.1",
		"127.0.0.1",
	)
	AssertStatusParsedAs(t,
		status[2],
		"test3.service",
		"loaded",
		"active",
		"exited",
		"f3274266.../127.0.0.1",
		"127.0.0.1",
	)
}
