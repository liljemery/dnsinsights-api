package dnscheck

import (
	"net"

	mdns "github.com/miekg/dns"
)

type DNSService struct{}

func NewDNSService() *DNSService { return &DNSService{} }

type ResolveResult struct {
	A      []string `json:"A"`
	AAAA   []string `json:"AAAA"`
	MX     []string `json:"MX"`
	TXT    []string `json:"TXT"`
	CNAME  string   `json:"CNAME,omitempty"`
}

func (s *DNSService) Resolve(domain string) (*ResolveResult, error) {
	var res ResolveResult

	ips, _ := net.LookupIP(domain)
	for _, ip := range ips {
		if ip.To4() != nil {
			res.A = append(res.A, ip.String())
		} else {
			res.AAAA = append(res.AAAA, ip.String())
		}
	}

	mx, _ := net.LookupMX(domain)
	for _, m := range mx {
		res.MX = append(res.MX, m.Host)
	}

	txt, _ := net.LookupTXT(domain)
	for _, t := range txt {
		res.TXT = append(res.TXT, t)
	}

	cname, _ := net.LookupCNAME(domain)
	res.CNAME = cname

	return &res, nil
}

type ReverseResult struct {
	Domains []string `json:"domains"`
}

func (s *DNSService) Reverse(ip string) (*ReverseResult, error) {
	hosts, err := net.LookupAddr(ip)
	if err != nil {
		return nil, err
	}
	return &ReverseResult{Domains: hosts}, nil
}

type TTLResult struct {
	TTL uint32 `json:"ttl"`
}

func (s *DNSService) TTL(domain string) (*TTLResult, error) {
	m := new(mdns.Msg)
	m.SetQuestion(mdns.Fqdn(domain), mdns.TypeA)
	c := new(mdns.Client)
	in, _, err := c.Exchange(m, net.JoinHostPort("8.8.8.8", "53"))
	if err != nil {
		return nil, err
	}
	for _, ans := range in.Answer {
		if rr, ok := ans.(*mdns.A); ok {
			return &TTLResult{TTL: rr.Hdr.Ttl}, nil
		}
	}
	return &TTLResult{TTL: 0}, nil
}
