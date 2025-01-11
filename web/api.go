package web

import (
	"context"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	appShared "github.com/placeholder/boiler/app/shared"
	"github.com/placeholder/boiler/pkg/logger"
	"github.com/placeholder/boiler/web/routes"
)

func New(lib appShared.AppInterface) *APIServer {
	api := &APIServer{
		app: lib,
	}

	server := echo.New()
	server.IPExtractor = api.ipMiddleware

	server.Pre(middleware.RemoveTrailingSlash())
	server.Use(middleware.Static("/static"))
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())
	server.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	server.Use(middleware.CORS())
	server.Use(middleware.BodyLimit("5M"))
	server.Use(middleware.Secure())
	server.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      middleware.DefaultSkipper,
		ErrorMessage: "request timed out, please try again in a moment.",
		OnTimeoutRouteErrorHandler: func(err error, c echo.Context) {
			logger.Logger().Warnf("request timed out: %s\n", c.Path())
		},
		Timeout: 15 * time.Second,
	}))

	api.echo = server

	routes.Mount(api.echo)

	return api
}

type APIServer struct {
	echo *echo.Echo
	app  appShared.AppInterface
}

func (s *APIServer) Shutdown(ctx context.Context) {
	s.echo.Shutdown(ctx)
}

func (s *APIServer) ipMiddleware(req *http.Request) string {
	// check if behind CF proxy
	ip := req.Header.Get("CF-Connecting-IP")

	if ip != "" {
		return ip
	}

	// otherwise fallback on echo's solution
	directIP := echo.ExtractIPDirect()(req)
	realIP := req.Header.Get(echo.HeaderXRealIP)
	if realIP == "" {
		return directIP
	}

	realIP = strings.TrimPrefix(realIP, "[")
	realIP = strings.TrimSuffix(realIP, "]")
	if rIP := net.ParseIP(realIP); rIP != nil {
		return realIP
	}

	return directIP
}

func (s *APIServer) Listen() {
	addr := os.Getenv("API_LISTEN_ADDR")
	if addr == "" {
		addr = ":24816"
	}

	s.echo.Start(addr)
}
