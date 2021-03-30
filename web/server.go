package web

import (
	"github.com\LERSONG\scaffold-example/config"
	"github.com\LERSONG\scaffold-example/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type server struct {
	config  *config.Config
	router  *gin.Engine
	service model.Service
}

// NewServer returns new http.Server.
func NewServer(cfg *config.Config, service model.Service) *http.Server {
	gin.SetMode(gin.ReleaseMode)
	srv := &server{
		config:  cfg,
		router:  gin.Default(),
		service: service,
	}

	return &http.Server{
		Addr:    cfg.HTTP.ListenAddr,
		Handler: srv.routes(),
	}
}
