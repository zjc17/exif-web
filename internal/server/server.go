package server

import (
	"fmt"
	"github.com/gin-contrib/cors"
	_ "github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	v1 "github.com/zjc17/exif-web/internal/server/api/v1"
	"net/http"
)

type (
	// LaunchParams for gin server
	LaunchParams struct {
		Port       uint16
		CorsConfig cors.Config
	}
)

func DefaultLaunchParam() *LaunchParams {
	//TODO config cors by env
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	return &LaunchParams{
		Port:       8080,
		CorsConfig: config,
	}
}

func Launch(params *LaunchParams) error {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Use(cors.New(params.CorsConfig))

	router.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	router.GET("/api/v1/parse", v1.Parse)
	return router.Run(fmt.Sprintf(":%d", params.Port))
}
