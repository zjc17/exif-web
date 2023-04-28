package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/zjc17/exif-web/internal/fetcher"
	"github.com/zjc17/exif-web/internal/parser"
	"net/http"
)

func Parse(c *gin.Context) {
	url := c.Query("url")
	image, _ := fetcher.GetImagePartial(url, &fetcher.GetImageParam{
		RequestHeader: fetcher.BuildRequestHeader(c.GetHeader),
	})
	p := parser.NewParser()
	result, _ := p.Parse(image)
	c.JSON(http.StatusOK, result)
	return
}
