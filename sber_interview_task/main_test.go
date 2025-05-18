package main

import (
	"testing"
)

func TestNewObject(t *testing.T) {
	testTable := []struct {
		kind         string
		expectedFail bool
	}{
		{kind: "base", expectedFail: false},
		{kind: "child", expectedFail: false},
		{kind: "robot", expectedFail: true},
	}

	for _, tc := range testTable {
		t.Logf("Creating %s", tc.kind)
		obj := NewObject(tc.kind)

		switch tc.kind {
		case "child":
			_, ok := obj.(Child)
			if !ok {
				t.Fatalf("Test failed, expected Child object")
			}
			break
		case "base":
			_, ok := obj.(Base)
			if !ok {
				t.Fatalf("Test failed, expected Base object")
			}
			break
		default:
			if !tc.expectedFail {
				t.Fatal("Unexpected args")
			}
		}
	}
}
