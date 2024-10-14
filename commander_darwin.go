//go:build darwin
// +build darwin

package main

import (
	"fmt"
	"net"
	"os"
	"log"

	ping "github.com/sparrc/go-ping"
)

// implement Commander interface for Mac
type commander struct{}

func NewCommander() Commander {
    return &commander{}
}

func (c *commander) GetSystemInfo() (sysInfo SystemInfo, err error) {
    hostname, err := os.Hostname()
    if err != nil {
        return SystemInfo{}, err
    }
    
    // Get IP address
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		log.Printf("Unable to get interface addresses: %s", err)
		return
	}

	var address string
	for _, addr := range addresses {
		ipnet, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		// filter out loop back
		if ipnet.IP.IsLoopback() {
			continue
		}
		// requirement is ambigous so just get the first address.
		if ipnet.IP.To4() != nil {
			address = ipnet.IP.String()
			break
		}
	}

	if address == "" {
		log.Printf("Unable to get valid IP address")
		err = fmt.Errorf("Unable to get valid IP address")
		return
	}

    return SystemInfo{
        Hostname:  hostname,
        IPAddress: address,
    }, nil
}

func (c *commander) Ping(host string) (results PingResult, err error) {
	pinger, err := ping.NewPinger(host)
	if err != nil {
		return
	}
	pinger.Count = 5
	pinger.Run()
	stats := pinger.Statistics()
	if stats == nil {
		err = fmt.Errorf("Unable to get pinger statistic")
		return
	}
	// requirement is ambiguous so just use the average round-trip time
	results = PingResult{
		Successful: true,
		Time: stats.AvgRtt,
	}
	return
}