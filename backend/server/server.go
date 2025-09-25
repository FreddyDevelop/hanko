package server

import (
	"github.com/labstack/echo/v4"
	"github.com/FreddyDevelop/hanko/backend/v2/config"
	"github.com/FreddyDevelop/hanko/backend/v2/handler"
	"github.com/FreddyDevelop/hanko/backend/v2/mapper"
	"github.com/FreddyDevelop/hanko/backend/v2/persistence"
	"sync"
)

func StartPublic(cfg *config.Config, wg *sync.WaitGroup, persister persistence.Persister, prometheus echo.MiddlewareFunc, authenticatorMetadata mapper.AuthenticatorMetadata) {
	defer wg.Done()
	router := handler.NewPublicRouter(cfg, persister, prometheus, authenticatorMetadata)
	router.Logger.Fatal(router.Start(cfg.Server.Public.Address))
}

func StartAdmin(cfg *config.Config, wg *sync.WaitGroup, persister persistence.Persister, prometheus echo.MiddlewareFunc) {
	defer wg.Done()
	router := handler.NewAdminRouter(cfg, persister, prometheus)
	router.Logger.Fatal(router.Start(cfg.Server.Admin.Address))
}
