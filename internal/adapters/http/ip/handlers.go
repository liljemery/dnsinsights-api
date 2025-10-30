package ip

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/jeremyinoa/dnsinsight-api/internal/modules/common"
	"github.com/jeremyinoa/dnsinsight-api/internal/services/ipchecker"
)

type Handler struct{}

func NewHandler() *Handler { return &Handler{} }

// Lookup godoc
// @Summary Basic IP information
// @Description Returns basic information about an IP address
// @Tags IP
// @Produce json
// @Param ip query string true "IP"
// @Success 200 {object} common.Response
// @Router /api/public/ip/lookup [get]
func (h *Handler) Lookup(c *fiber.Ctx) error {
	ip := c.Query("ip")
	if ip == "" {
		return c.Status(fiber.StatusBadRequest).JSON(common.Response{Status: "error", Message: "ip is required"})
	}
	info, err := ipchecker.Lookup(ip)
	if err != nil {
		if errors.Is(err, ipchecker.ErrInvalidIP) {
			return c.Status(fiber.StatusBadRequest).JSON(common.Response{Status: "error", Message: "invalid ip"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(common.Response{Status: "error", Message: err.Error()})
	}
	return c.JSON(common.Response{Status: "success", Message: "ip info", Data: info})
}
