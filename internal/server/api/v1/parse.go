package v1

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zjc17/exif-web/internal/define"
	"github.com/zjc17/exif-web/internal/storage"
	"github.com/zjc17/exif-web/internal/util"
	"github.com/zjc17/exif-web/pkg/fetcher"
	"github.com/zjc17/exif-web/pkg/parser"
	"io"
	"log"
	"net/http"
)

func GetParse(c *gin.Context) {
	url := c.Query("url")

	var jsonStr string

	// Check Storage
	jsonStr, err := storage.Get(url)

	if err != nil {
		fmt.Println("cache miss:", url)
		if !errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		image, _ := fetcher.GetImagePartial(url, &fetcher.GetImageParam{
			RequestHeader: fetcher.BuildRequestHeader(c.GetHeader),
		})

		p := parser.NewParser()
		jsonStr, err = p.ParseJsonString(image)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		err = storage.Store(url, jsonStr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "failed to store exif")
			return
		}
	} else {
		fmt.Println("cache hit:", url)
	}

	result := new(parser.ParseResult)
	err = result.From(jsonStr)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, result)
	return
}

func PostParse(c *gin.Context) {
	image, err := getUploadedFilePartial(c)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	p := parser.NewParser()
	result, _ := p.Parse(image)
	c.JSON(http.StatusOK, result)
	return
}

func getUploadedFilePartial(c *gin.Context) (image []byte, err error) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		log.Printf("get file header error: %v", err)
		c.Status(http.StatusBadRequest)
		return
	}
	// read the file
	file, err := fileHeader.Open()
	if err != nil {
		log.Printf("read file error: %v", err)
		c.Status(http.StatusBadRequest)
		return
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		log.Printf("read file error: %v", err)
		c.Status(http.StatusBadRequest)
		return
	}
	image = util.TrimByteArray(data, define.ParseDataLength)
	return
}
