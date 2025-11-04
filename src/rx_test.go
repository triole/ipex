package main

import "testing"

func TestNotations(t *testing.T) {
	assertIsCIDRNotation("192.168.33.1", false, t)
	assertIsCIDRNotation("192.168.33.1/30", true, t)
	assertIsCIDRNotation("255.255.0.0", false, t)
	assertIsCIDRNotation("255.255.0.0/16", true, t)
	assertIsCIDRNotation("255.0", false, t)
	assertIsCIDRNotation("255.0/8", true, t)
	assertIsCIDRNotation("1", false, t)
	assertIsCIDRNotation("1/30", true, t)

	assertIsIPv4Full("192.168.33.1", true, t)
	assertIsIPv4Full("255.255.0.0", true, t)
	assertIsIPv4Full("255.255.0", false, t)
	assertIsIPv4Full("255.0", false, t)
	assertIsIPv4Full("1", false, t)

	assertIsInt("192.168.33.1", false, t)
	assertIsInt("255.255.0.0", false, t)
	assertIsInt("255.255.0", false, t)
	assertIsInt("255.0", false, t)
	assertIsInt("1", true, t)
}

func assertIsCIDRNotation(inp string, exp bool, t *testing.T) {
	res := isCIDRNotation(inp)
	if res != exp {
		t.Errorf("isCIDRNotation failed: %s -> %v != %v", inp, exp, res)
	}
}

func assertIsIPv4Full(inp string, exp bool, t *testing.T) {
	res := isIPv4Full(inp)
	if res != exp {
		t.Errorf("isIPv4Full failed: %s -> %v != %v", inp, exp, res)
	}
}

func assertIsInt(inp string, exp bool, t *testing.T) {
	res := isInt(inp)
	if res != exp {
		t.Errorf("isInt failed: %s -> %v != %v", inp, exp, res)
	}
}
