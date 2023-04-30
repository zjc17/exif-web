package server

import (
	"fmt"
	"github.com/gin-contrib/cors"
	_ "github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	v1 "github.com/zjc17/exif-web/internal/server/api/v1"
	"github.com/zjc17/exif-web/internal/storage"
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

func Launch(params *LaunchParams) (err error) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Use(cors.New(params.CorsConfig))

	router.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html", []byte(PageDocument))
	})
	router.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	router.GET("/api/v1/parse", v1.GetParse)
	router.POST("/api/v1/parse", v1.PostParse)

	if err = storage.InitDatabase(nil); err != nil {
		return
	}

	fmt.Printf("Server started on :%d\n", params.Port)
	if err = router.Run(fmt.Sprintf(":%d", params.Port)); err != nil {
		fmt.Println("Failed to start server.")
		return
	}
	return
}
