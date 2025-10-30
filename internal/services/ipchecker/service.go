package ipchecker

import (
	"net"
)

type Info struct {
	IP              string `json:"ip"`
	Version         string `json:"version"`
	IsPrivate       bool   `json:"isPrivate"`
	IsLoopback      bool   `json:"isLoopback"`
	IsMulticast     bool   `json:"isMulticast"`
	IsGlobalUnicast bool   `json:"isGlobalUnicast"`
}

func Lookup(ipStr string) (*Info, error) {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return nil, ErrInvalidIP
	}
	version := "IPv6"
	if ip.To4() != nil {
		version = "IPv4"
	}
	info := &Info{
		IP:              ip.String(),
		Version:         version,
		IsPrivate:       isPrivateIP(ip),
		IsLoopback:      ip.IsLoopback(),
		IsMulticast:     ip.IsMulticast(),
		IsGlobalUnicast: ip.IsGlobalUnicast(),
	}
	return info, nil
}

var ErrInvalidIP = &net.ParseError{Type: "IP address", Text: "invalid"}

func isPrivateIP(ip net.IP) bool {
	privateBlocks := []string{
		"10.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/16",
		"127.0.0.0/8",
		"169.254.0.0/16",
	}
	for _, cidr := range privateBlocks {
		_, network, _ := net.ParseCIDR(cidr)
		if network.Contains(ip) {
			return true
		}
	}
	return false
}
