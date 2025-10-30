package dns

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jeremyinoa/dnsinsight-api/internal/modules/common"
	"github.com/jeremyinoa/dnsinsight-api/internal/services/dnscheck"
)

type Handler struct {
	service *dnscheck.DNSService
}

func NewHandler() *Handler {
	return &Handler{service: dnscheck.NewDNSService()}
}

// Resolve godoc
// @Summary Resolve DNS records
// @Description Returns A, AAAA, MX, TXT, CNAME records for a domain
// @Tags DNS
// @Produce json
// @Param domain query string true "Domain to resolve"
// @Success 200 {object} common.Response
// @Router /api/public/dns/resolve [get]
func (h *Handler) Resolve(c *fiber.Ctx) error {
	domain := c.Query("domain")
	if domain == "" {
		return c.Status(fiber.StatusBadRequest).JSON(common.Response{Status: "error", Message: "domain is required"})
	}
	res, err := h.service.Resolve(domain)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(common.Response{Status: "error", Message: err.Error()})
	}
	return c.JSON(common.Response{Status: "success", Message: "DNS records found", Data: res})
}

// Reverse godoc
// @Summary Reverse lookup of IP
// @Description Returns domains associated to an IP
// @Tags DNS
// @Produce json
// @Param ip query string true "IP to reverse lookup"
// @Success 200 {object} common.Response
// @Router /api/public/ip/reverse [get]
func (h *Handler) Reverse(c *fiber.Ctx) error {
	ip := c.Query("ip")
	if ip == "" {
		return c.Status(fiber.StatusBadRequest).JSON(common.Response{Status: "error", Message: "ip is required"})
	}
	res, err := h.service.Reverse(ip)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(common.Response{Status: "error", Message: err.Error()})
	}
	return c.JSON(common.Response{Status: "success", Message: "reverse lookup done", Data: res})
}

// TTL godoc
// @Summary Get TTL of domain
// @Description Returns TTL of A record for a domain
// @Tags DNS
// @Produce json
// @Param domain query string true "Domain"
// @Success 200 {object} common.Response
// @Router /api/public/network/ttl [get]
func (h *Handler) TTL(c *fiber.Ctx) error {
	domain := c.Query("domain")
	if domain == "" {
		return c.Status(fiber.StatusBadRequest).JSON(common.Response{Status: "error", Message: "domain is required"})
	}
	res, err := h.service.TTL(domain)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(common.Response{Status: "error", Message: err.Error()})
	}
	return c.JSON(common.Response{Status: "success", Message: "ttl fetched", Data: res})
}
