package network

import (
	"time"

	"github.com/go-ping/ping"
)

type PingResult struct {
	PacketsSent     int           `json:"packetsSent"`
	PacketsReceived int           `json:"packetsReceived"`
	PacketLoss      float64       `json:"packetLoss"`
	MinRtt          time.Duration `json:"minRtt"`
	AvgRtt          time.Duration `json:"avgRtt"`
	MaxRtt          time.Duration `json:"maxRtt"`
}

func Ping(host string, count int, timeout time.Duration) (*PingResult, error) {
	pinger, err := ping.NewPinger(host)
	if err != nil {
		return nil, err
	}
	pinger.Count = count
	pinger.Timeout = timeout
	pinger.SetPrivileged(true)
	if err := pinger.Run(); err != nil {
		return nil, err
	}
	stats := pinger.Statistics()
	return &PingResult{
		PacketsSent:     stats.PacketsSent,
		PacketsReceived: stats.PacketsRecv,
		PacketLoss:      stats.PacketLoss,
		MinRtt:          stats.MinRtt,
		AvgRtt:          stats.AvgRtt,
		MaxRtt:          stats.MaxRtt,
	}, nil
}
