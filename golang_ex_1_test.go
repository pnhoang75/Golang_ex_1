package main

import (
	"fmt"
	"testing"
)

func TestGetSystemInfo(t *testing.T) {
	cmdr := NewCommander()
	info, err := cmdr.GetSystemInfo()

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if info.Hostname == "" {
		t.Error("Expected hostname to be non-empty")
	}

	if info.IPAddress == "" {
		t.Error("Expected IP address to be non-empty")
	}
}

func TestPing(t *testing.T) {
	cmdr := NewCommander()
	tests := []struct {
		testname    string
		hostname    string
		expectedErr error
	}{
		{
			testname:    "Invalid host",
			hostname:    "randomhost",
			expectedErr: fmt.Errorf("lookup randomhost: no such host"),
		},
		{
			testname: "localhost",
			hostname: "localhost",
		},
		{
			testname: "remote host",
			hostname: "www.cnn.com",
		},
	}
	for _, tc := range tests {
		fmt.Printf("testcase: %s\n", tc.testname)
		pingRes, err := cmdr.Ping(tc.hostname)
		if tc.expectedErr != nil {
			if err.Error() != tc.expectedErr.Error() {
				t.Fatalf("Expected %q but received %q", tc.expectedErr, err)
			}
			continue
		}
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if !pingRes.Successful {
			t.Error("Expected successful but received failure")
		}
	}
}
