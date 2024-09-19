package cmd

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"go-redis-url-shortener/config"
	"go-redis-url-shortener/connections"
	"go-redis-url-shortener/controllers"
	repoImpl "go-redis-url-shortener/repositories/impl"
	"go-redis-url-shortener/routes"
	"go-redis-url-shortener/server"
	serviceImpl "go-redis-url-shortener/services/impl"
)

var serveCmd = &cobra.Command{
	Use: "serve",
	Run: serve,
}

func serve(cmd *cobra.Command, args []string) {
	// config
	config.Load()
	cfg := config.AllConfig()

	connections.NewRedisClient(cfg.Redis)

	// redis
	cacheClient := connections.RedisClient()

	// register all repositories
	cacheRepository := repoImpl.NewRedisRepository(cacheClient)
	urlService := serviceImpl.NewUrlService(cacheRepository)
	urlController := controllers.NewUrlController(urlService)

	// Server
	var framework = echo.New()
	var Routes = routes.New(framework, urlController)
	var Server = server.New(cfg, framework)

	Routes.Init()
	Server.Start()
}
