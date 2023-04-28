package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zjc17/exif-web/internal/fetcher"
	"github.com/zjc17/exif-web/internal/parser"
	"net/http"
)

func Parse(c *gin.Context) {
	url := c.Query("url")
	fmt.Println("url", url)
	image, _ := fetcher.GetImagePartial(url)
	p := parser.NewParser()
	result, _ := p.Parse(image)
	fmt.Printf("fulfilled: %+v\n", result)
	c.JSON(http.StatusOK, result)
	//data, _ := json.Marshal(result)
	//c.Data(http.StatusOK, "appliction/json; charset=UTF-8", data)
	return
}
