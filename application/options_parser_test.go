package application

import (
	"testing"
)

func TestParsesLongHelpParamProperly(t *testing.T) {
	opts := parseOptions([]string{"program", "--help"})

	if !opts.Help {
		t.Fatal("parses --help param properly")
	}
}

func TestParsesShortHelpParamProperly(t *testing.T) {
	opts := parseOptions([]string{"program", "-h"})

	if !opts.Help {
		t.Fatal("parses -h param properly")
	}
}

func TestEnablesHelpIfNoArgumentWasGiven(t *testing.T) {
	opts := parseOptions([]string{"program"})

	if !opts.Help {
		t.Fatal("prints help if no one argument was given")
	}
}

func TestEnablesHelpFlagIfUnrecognizedParameterWasGiven(t *testing.T) {
	opts := parseOptions([]string{"program", "bbb"})

	if !opts.Help {
		t.Fatal("prints help unrecognized arg was given")
	}

	if !opts.hasUnrecognizedArgs() {
		t.Fatal("should have unrecognized params")
	}
}



