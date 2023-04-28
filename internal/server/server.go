package server

import (
	"fmt"
	"github.com/gin-contrib/gzip"
	v1 "github.com/zjc17/exif-web/internal/server/api/v1"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	// LaunchParams for gin server
	LaunchParams struct {
		Port uint16
	}
)

func DefaultLaunchParam() *LaunchParams {
	return &LaunchParams{
		Port: 8080,
	}
}

func Launch(params *LaunchParams) error {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	router.GET("/api/v1/parse", v1.Parse)
	return router.Run(fmt.Sprintf(":%d", params.Port))
}
