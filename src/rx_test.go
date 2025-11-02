package main

import "testing"

func TestIsIPFull(t *testing.T) {
	assertCidrNotateSplit("192.168.33.1", []string{"192.168.33.1", "32"}, t)
	assertCidrNotateSplit("192.168.33.1/30", []string{"192.168.33.1", "30"}, t)
	assertCidrNotateSplit("255.255.0.0", []string{"255.255.0.0", "32"}, t)
	assertCidrNotateSplit("255.255.0.0/16", []string{"255.255.0.0", "16"}, t)
	assertCidrNotateSplit("255.0", []string{"255.0", "32"}, t)
	assertCidrNotateSplit("255.0/8", []string{"255.0", "8"}, t)
	assertCidrNotateSplit("1", []string{"1", "32"}, t)
	assertCidrNotateSplit("1/30", []string{"1", "30"}, t)

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

	assertIsIPv4Fragment("192.168.33.1", true, t)
	assertIsIPv4Fragment("255.255.0.0", true, t)
	assertIsIPv4Fragment("255.255.0", true, t)
	assertIsIPv4Fragment("255.0", true, t)
	assertIsIPv4Fragment("1", false, t)

	assertIsInt("192.168.33.1", false, t)
	assertIsInt("255.255.0.0", false, t)
	assertIsInt("255.255.0", false, t)
	assertIsInt("255.0", false, t)
	assertIsInt("1", true, t)
}

func assertCidrNotateSplit(inp string, exp []string, t *testing.T) {
	res := cidrNotateSplit(inp)
	if len(res) != len(exp) {
		t.Errorf(
			"result and expecation differ in length: %s -> %s != %s",
			inp, exp, res,
		)
	}
	for i := 1; i <= len(res)-1; i++ {
		if exp[i] != res[i] {
			t.Errorf(
				"result and expectation differ: %s -> %s != %s",
				inp, exp, res,
			)
		}
	}
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

func assertIsIPv4Fragment(inp string, exp bool, t *testing.T) {
	res := isIPv4Fragment(inp)
	if res != exp {
		t.Errorf("isIPv4Fragment failed: %s -> %v != %v", inp, exp, res)
	}
}

func assertIsInt(inp string, exp bool, t *testing.T) {
	res := isInt(inp)
	if res != exp {
		t.Errorf("isInt failed: %s -> %v != %v", inp, exp, res)
	}
}
