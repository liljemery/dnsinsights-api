package network

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jeremyinoa/dnsinsight-api/internal/modules/common"
	netservice "github.com/jeremyinoa/dnsinsight-api/internal/services/network"
)

type Handler struct{}

func NewHandler() *Handler { return &Handler{} }

// Ping godoc
// @Summary Ping a host
// @Description Executes ICMP ping to a host
// @Tags Network
// @Produce json
// @Param host query string true "Host to ping"
// @Param count query int false "Packets count"
// @Success 200 {object} common.Response
// @Router /api/public/network/ping [get]
func (h *Handler) Ping(c *fiber.Ctx) error {
	host := c.Query("host")
	if host == "" {
		return c.Status(fiber.StatusBadRequest).JSON(common.Response{Status: "error", Message: "host is required"})
	}
	count := 4
	if v := c.QueryInt("count"); v > 0 {
		count = v
	}
	res, err := netservice.Ping(host, count, 8*time.Second)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(common.Response{Status: "error", Message: err.Error()})
	}
	return c.JSON(common.Response{Status: "success", Message: "ping results", Data: res})
}
