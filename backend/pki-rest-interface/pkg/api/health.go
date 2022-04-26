package api

import (
	"net/http"
	"sync/atomic"

	"github.com/labstack/echo/v4"
)

type healthResponse struct {
	Status string `json:"status"`
}

// healthzHandler godoc
func (s *Server) healthzHandler(c echo.Context) (err error) {
	if atomic.LoadInt32(&healthy) == 1 {
		return c.JSON(http.StatusOK, healthResponse{Status: "OK"})
	}
	return c.JSON(http.StatusServiceUnavailable, healthResponse{Status: "Service Unavailable"})
}

// readyzHandler godoc
func (s *Server) readyzHandler(c echo.Context) (err error) {
	if atomic.LoadInt32(&ready) == 1 {
		return c.JSON(http.StatusOK, healthResponse{Status: "OK"})
	}
	return c.JSON(http.StatusServiceUnavailable, healthResponse{Status: "Service Unavailable"})
}
