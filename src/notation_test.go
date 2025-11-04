package main

import "testing"

func TestIsIPFull(t *testing.T) {

	assertNotateSplit("192.168.33.1", newNt("192.168.33.1"), t)
	assertNotateSplit("192.168.33.1/30", newNt("192.168.33.1", "/", "30"), t)
	assertNotateSplit("255.255.0.0", newNt("255.255.0.0"), t)
	assertNotateSplit("255.255.0.0/16", newNt("255.255.0.0", "/", "16"), t)
	assertNotateSplit("255.0", newNt("255.0"), t)
	assertNotateSplit("255.0/8", newNt("255.0", "/", "8"), t)
	assertNotateSplit("1", newNt("1"), t)
	assertNotateSplit("1/30", newNt("1", "/", "30"), t)

	assertNotateSplit("192.168.33.1", newNt("192.168.33.1"), t)
	assertNotateSplit("192.168.33.1+30", newNt("192.168.33.1", "+", "30"), t)
	assertNotateSplit("255.255.0.0", newNt("255.255.0.0"), t)
	assertNotateSplit("255.255.0.0+16", newNt("255.255.0.0", "+", "16"), t)
	assertNotateSplit("255.0", newNt("255.0"), t)
	assertNotateSplit("255.0+8", newNt("255.0", "+", "8"), t)
	assertNotateSplit("1", newNt("1"), t)
	assertNotateSplit("1+30", newNt("1", "+", "30"), t)

	assertNotateSplit("192.168.33.1", newNt("192.168.33.1"), t)
	assertNotateSplit("192.168.33.1-30", newNt("192.168.33.1", "-", "30"), t)
	assertNotateSplit("255.255.0.0", newNt("255.255.0.0"), t)
	assertNotateSplit("255.255.0.0-16", newNt("255.255.0.0", "-", "16"), t)
	assertNotateSplit("255.0", newNt("255.0"), t)
	assertNotateSplit("255.0-8", newNt("255.0", "-", "8"), t)
	assertNotateSplit("1", newNt("1"), t)
	assertNotateSplit("1-30", newNt("1", "-", "30"), t)

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

func assertNotateSplit(inp string, exp tNotation, t *testing.T) {
	res := notate(inp)
	if res.Pre != exp.Pre || res.Op != exp.Op || res.Suf != exp.Suf {
		t.Errorf(
			"notateSplit failed: %s -> %v != %v",
			inp, exp, res,
		)
	}
}

func newNt(params ...any) (nt tNotation) {
	if len(params) > 0 {
		nt.Pre = params[0].(string)
	}
	if len(params) > 1 {
		nt.Op = params[1].(string)
	}
	if len(params) > 2 {
		nt.Suf = params[2].(string)
	}
	return nt
}
