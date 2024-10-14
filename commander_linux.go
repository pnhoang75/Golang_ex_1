//go:build linux
// +build linux

package main

import (
	"os"
)

// implement Commander interface for Linux
type commander struct{}

func NewCommander() Commander {
	return &commander{}
}

func (c *commander) GetSystemInfo() (SystemInfo, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return SystemInfo{}, err
	}

	// Get IP address (implement this)

	return SystemInfo{
		Hostname:  hostname,
		IPAddress: "implement me",
	}, nil
}

func (c *commander) Ping(host string) (results PingResult, err error) {
	// TODO: implement this
	return
}
